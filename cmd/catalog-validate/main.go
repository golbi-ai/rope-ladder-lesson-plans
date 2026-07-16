package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

const catalogFormat = 1

var slugPattern = regexp.MustCompile(`^[a-z0-9][a-z0-9-]{0,79}$`)
var tagPattern = regexp.MustCompile(`^[a-z0-9][a-z0-9-]{0,31}$`)

type metadata struct {
	Format      int      `yaml:"format"`
	Name        string   `yaml:"name"`
	Codebase    string   `yaml:"codebase,omitempty"`
	Description string   `yaml:"description"`
	Tags        []string `yaml:"tags"`
	Attestation struct {
		PublicSource    bool `yaml:"public_source"`
		RightsToPublish bool `yaml:"rights_to_publish"`
	} `yaml:"attestation"`
}

var requiredFiles = []string{
	"metadata.yaml",
	"architecture.html",
	"entity-relationship.html",
	"coverage.md",
	"lesson-plan.md",
	"lesson-plan.json",
	"guides/survival.md",
	"references/domain-concepts.md",
	"references/language-concepts.md",
	"references/entities.md",
}

func main() {
	entries, err := os.ReadDir("plans")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return
		}
		fail(err)
	}
	var failures []string
	for _, entry := range entries {
		if !entry.IsDir() || strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		if err := validateEntry(filepath.Join("plans", entry.Name()), entry.Name()); err != nil {
			failures = append(failures, err.Error())
		}
	}
	if len(failures) > 0 {
		sort.Strings(failures)
		fail(errors.New(strings.Join(failures, "\n")))
	}
	fmt.Println("catalog validation passed")
}

func validateEntry(root, slug string) error {
	if !slugPattern.MatchString(slug) {
		return fmt.Errorf("%s: invalid stable resource slug", root)
	}
	for _, name := range requiredFiles {
		if info, err := os.Stat(filepath.Join(root, name)); err != nil || info.IsDir() {
			return fmt.Errorf("%s: missing required file %s", root, name)
		}
	}
	data, err := os.ReadFile(filepath.Join(root, "metadata.yaml"))
	if err != nil {
		return err
	}
	var meta metadata
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	decoder.KnownFields(true)
	if err := decoder.Decode(&meta); err != nil {
		return fmt.Errorf("%s: parse metadata: %w", root, err)
	}
	if err := validateMetadata(meta); err != nil {
		return fmt.Errorf("%s: %w", root, err)
	}
	if data, err := os.ReadFile(filepath.Join(root, "lesson-plan.json")); err != nil || !json.Valid(data) {
		return fmt.Errorf("%s: lesson-plan.json must be valid JSON", root)
	}
	return filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil || entry.IsDir() || filepath.Ext(path) != ".html" {
			return walkErr
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		lower := strings.ToLower(string(data))
		if strings.Contains(lower, "<script") || strings.Contains(lower, "src=\"http") || strings.Contains(lower, "href=\"http") {
			return fmt.Errorf("%s: HTML must be standalone without scripts or hosted assets", path)
		}
		return nil
	})
}

func validateMetadata(meta metadata) error {
	if meta.Format != catalogFormat {
		return fmt.Errorf("metadata format must be %d", catalogFormat)
	}
	if len(strings.TrimSpace(meta.Name)) == 0 || len(meta.Name) > 120 {
		return errors.New("name must contain 1-120 characters")
	}
	if len(strings.TrimSpace(meta.Description)) == 0 || len(meta.Description) > 600 {
		return errors.New("description must contain 1-600 characters")
	}
	if meta.Codebase != "" {
		parsed, err := url.Parse(meta.Codebase)
		if err != nil || (parsed.Scheme != "https" && parsed.Scheme != "http") || parsed.Host == "" {
			return errors.New("codebase must be an absolute http(s) URL when provided")
		}
	}
	if len(meta.Tags) == 0 || len(meta.Tags) > 8 {
		return errors.New("tags must contain 1-8 values")
	}
	seen := map[string]bool{}
	for _, tag := range meta.Tags {
		if !tagPattern.MatchString(tag) || seen[tag] {
			return fmt.Errorf("invalid or duplicate tag %q", tag)
		}
		seen[tag] = true
	}
	if !meta.Attestation.PublicSource || !meta.Attestation.RightsToPublish {
		return errors.New("public_source and rights_to_publish attestations must both be true")
	}
	return nil
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, "catalog validation failed:", err)
	os.Exit(1)
}
