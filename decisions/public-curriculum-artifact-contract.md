---
id: public-curriculum-artifact-contract
type: decision
governs:
    - public curriculum artifact structure
    - plans/*/metadata.yaml
    - plans/*/architecture.html
    - plans/*/entity-relationship.html
    - plans/*/coverage.md
    - plans/*/lesson-plan.md
    - plans/*/lesson-plan.json
    - plans/*/guides/survival.md
    - plans/*/references/domain-concepts.md
    - plans/*/references/language-concepts.md
    - plans/*/references/entities.md
status: current
confidence: high
source: implementation
decision: |
    Every public curriculum is a portable, generated snapshot under plans/<stable-slug>/ with a
    paired human and machine learning path. metadata.yaml declares format 1, the public codebase
    when available, a concise description, bounded tags, and the contributor's public-source and
    publication-rights attestations. architecture.html is a standalone, script-free overview that
    explains the system hierarchy before bounded subsystem deep dives. entity-relationship.html
    is a standalone, script-free visual reference that groups meaningful entities into readable
    relationship clusters, supplies a textual register and glossary, and complements rather than
    replaces references/entities.md. coverage.md makes indexing, evidence, citation, scope,
    relationship-resolution, and omission boundaries visible.

    lesson-plan.md is the readable syllabus: stable numbered lesson IDs, interleaved domain and
    implementation topics, assigned references, examples, and explainable prerequisite DAGs.
    lesson-plan.json is its valid machine-readable counterpart for the tutor and tools; it must
    represent the same curriculum without becoming a separately edited syllabus. The survival
    guide gives an early contribution-oriented orientation. Domain and implementation-language
    references are guided explanatory reading paths, not dictionary glossaries: each grounded
    topic builds from prerequisites, connects to evidence and small code examples where useful,
    and names its relationship to the curriculum. The entity reference documents meaningful
    persistent, externally exchanged, or cross-cutting structures with their role, lifecycle,
    fields, relationships, and evidence.

    Rope-ladder regeneration is the normal way to update these correlated artifacts. Human
    corrections remain reviewable pull-request changes and must preserve the catalog's static,
    standalone, source-evidenced posture. Structural requirements are enforced by the catalog
    validation contract; source rights, fact checking, and educational quality remain explicit
    maintainer review responsibilities.
rationale: |
    Onboarding requires more than a diagram or an inventory: readers need an ordered path from
    conceptual vocabulary to concrete implementation examples, while tutoring tools need a
    stable machine representation. Publishing both forms, with coverage and evidence boundaries,
    lets people assess what the curriculum can support rather than mistaking an exhaustive-looking
    artifact for complete source coverage. Standalone static HTML remains portable and safe to
    host, and separate domain, language, and entity references make unfamiliar systems legible to
    humans and agents without conflating their distinct questions.
alternatives_rejected:
    - option: Publish only a machine-readable lesson graph
      reason: Learners and reviewers need concise explanations, visual orientation, and source-linked examples without a dedicated client.
    - option: Publish only Markdown and derive tutoring state at runtime
      reason: A stable JSON representation is needed for deterministic lesson identity, prerequisite checks, and tool integration.
    - option: Use interactive or externally hosted diagrams
      reason: Public curricula must remain portable, reviewable, and usable without third-party runtime dependencies.
    - option: Treat every data type as an entity or every language term as a curriculum topic
      reason: References must prioritize structures and concepts that materially support understanding and contribution, not reproduce an uncurated index.
triggers_review_if: |
    A curriculum adds a new public artifact, changes lesson identity or graph semantics, permits
    non-static presentation content, or separates the JSON syllabus from the human-readable
    lesson plan.
supersedes: null
last_validated: 2026-07
---

## Consistency rules

The architecture page provides the broad top-to-bottom map; deeper implementation details belong
in bounded subsystem sections and references. The relationship page must be readable without UML
fluency: arrows and cardinality are paired with descriptive register rows. The three reference
guides deliberately answer different questions — what the domain means, how the implementation
expresses it, and which data structures carry it — but lesson units connect them through the same
grounded topic identities.

## Evidence and maintenance

Each artifact is a point-in-time reading aid, not a claim that every source file or relationship
was analyzed. coverage.md surfaces exclusions, unverified citations, severed relationship targets,
and omitted material. Regeneration must respect managed artifact regions and produce reviewable
conflict proposals rather than overwrite an independently edited region.
