---
id: public-catalog-validation-contract
type: decision
governs:
    - public catalog entry validation
    - catalog metadata schema
    - static curriculum artifact safety
    - .github/workflows/validate.yml
    - cmd/catalog-validate/main.go
    - catalog.schema.json
status: current
confidence: high
source: implementation
decision: |
    The public curriculum catalog is a fail-closed structural and safety gate, not an
    assessment of a curriculum's pedagogical or factual quality. Every directory directly
    under plans/ uses a stable lowercase slug and must contain the versioned metadata file,
    the standalone architecture and entity-relationship pages, coverage.md, lesson-plan.md,
    lesson-plan.json, guides/survival.md, and the domain, language, and entity references.
    Metadata must satisfy format 1: a bounded name and description, one to eight unique
    lowercase tags, an optional absolute HTTP(S) public codebase URL, and true public-source
    and publication-rights attestations. The machine-readable lesson plan must be valid JSON;
    catalog HTML must be standalone and contain neither scripts nor hosted assets. CI runs the
    same validator for every pull request and main update. A new validation rule or a change to
    this contract requires an accompanying decision and matching schema, validator, and docs
    update.
rationale: |
    The catalog is public and serves directly from generated files, so an incomplete entry,
    invalid manifest, remotely loaded asset, or embedded script would create a broken or unsafe
    user experience. A single local Go validator makes the published contract reproducible for
    contributors and CI. Structural validation deliberately stops short of judging source rights,
    facts, or instructional quality: attestations and automated checks reduce routine mistakes,
    while maintainers retain responsibility for review of public-disclosure rights and quality.
alternatives_rejected:
    - option: Rely on pull-request review without an executable catalog contract
      reason: Required artifacts and static-HTML safety checks are deterministic and should fail before human review, not depend on reviewer memory.
    - option: Permit scripts or externally hosted HTML assets in catalog pages
      reason: Catalog pages must remain portable, reviewable, and safe to host without runtime trust in a third party.
    - option: Treat the validator as a pedagogical or factual-quality gate
      reason: Those judgments need source context and human review; structural checks cannot establish them reliably.
triggers_review_if: |
    The catalog adds a new artifact type, permits interactive or hosted content, changes the
    metadata version, or introduces a machine-readable quality or provenance standard that can
    be checked deterministically.
supersedes: null
last_validated: 2026-07
---

## Enforcement boundary

The validator walks only top-level plan directories, so hidden files and nested support files do
not become accidental catalog entries. It reports every invalid entry in deterministic order and
returns nonzero before publication or merge. The workflow invokes the checked-in command rather
than duplicating its rules in YAML; the root-only build-binary ignore rule must never exclude the
command source directory.

## Required artifact roles

- `metadata.yaml` identifies the public resource and contributor attestations.
- `architecture.html` and `entity-relationship.html` are standalone, human-readable visual aids.
- `coverage.md` reports analysis boundaries and omissions.
- `lesson-plan.md` and `lesson-plan.json` provide the human and machine representations of the
  curriculum.
- `guides/survival.md` and `references/*.md` provide the introductory, domain, implementation,
  and entity reading materials.
