# rope-ladder lesson plans

This public, MIT-licensed catalog contains reviewable curricula generated with
[rope-ladder](https://github.com/golbi-ai/rope-ladder). Each merged entry is a
standalone curriculum for a public codebase; the catalog never stores learner journals,
provider credentials, workspace state, source archives, or private repositories.

## Browse and submit

Browse merged curricula at [rpldr.golbi.ai](https://rpldr.golbi.ai). To propose a
curriculum, generate it in a separate workspace, prepare `metadata.yaml`, then run:

```sh
rope-ladder publish --metadata ./metadata.yaml --slug your-codebase --confirm-public <curriculum-directory>
```

The command uses your authenticated GitHub CLI session to fork this repository, create
a branch, and open a pull request. It does not persist GitHub credentials or publish
anything unless you explicitly run it with `--confirm-public`.

## Catalog layout

Each entry lives at `plans/<stable-resource-slug>/`:

```text
plans/<slug>/
  metadata.yaml
  architecture.html
  entity-relationship.html
  coverage.md
  lesson-plan.md
  lesson-plan.json
  guides/survival.md
  references/domain-concepts.md
  references/language-concepts.md
  references/entities.md
```

`metadata.yaml` is format-versioned and includes a name, optional public codebase URL,
description, tags, and the contributor's publication attestation. See
[catalog.schema.json](catalog.schema.json) and [metadata.example.yaml](metadata.example.yaml)
for the contract and a starting point.

## Review and safety

Pull requests are validated automatically for format, required artifacts, and unsafe
hosted HTML assets. Maintainers review every contribution for source rights, safety,
and educational quality before merging. Corrections and takedowns are made through
ordinary reviewable pull requests.

## Development

Validate the catalog locally with Go 1.25 or later:

```sh
go run ./cmd/catalog-validate
```
