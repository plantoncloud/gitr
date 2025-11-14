# Gitr Changelog

This directory contains detailed changelogs documenting significant features, improvements, and changes to gitr.

## Structure

Changelogs are organized by year and month:

```
_changelog/
├── YYYY-MM/
│   ├── YYYY-MM-DD-HHMMSS-feature-slug.md
│   ├── YYYY-MM-DD-HHMMSS-another-feature.md
│   └── ...
└── README.md (this file)
```

## Creating a Changelog

To create a new changelog, use the cursor rule:

```
@create-gitr-changelog
```

This will generate a properly structured changelog document with:
- Automatic timestamp-based filename
- Standard sections (Summary, Problem Statement, Solution, etc.)
- Appropriate sizing based on the change's scope

## When to Create Changelogs

Create changelogs for:
- ✅ New commands or significant command enhancements
- ✅ Major refactoring or architectural changes
- ✅ Features that improve user workflows
- ✅ Changes affecting multiple packages or components
- ✅ Work that took significant effort (1+ hours)

Skip changelogs for:
- ✋ Trivial bug fixes
- ✋ Minor configuration changes
- ✋ Work already documented in PR descriptions

## Changelog Purpose

Changelogs serve to:
- Document the "why" behind changes
- Capture design decisions and trade-offs
- Provide context for future developers
- Track feature evolution over time
- Help onboarding and knowledge transfer

---

For guidelines on writing changelogs, see `.cursor/rules/git/create-gitr-changelog.mdc`

