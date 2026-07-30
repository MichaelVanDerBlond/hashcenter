# HashCenter AI Agent Instructions

This file contains mandatory rules for any AI agent working on HashCenter.

The complete documentation is located in the `docs/` directory.
This file is only a quick operational guide.

===============================================================================
READ FIRST
===============================================================================

Before making any changes, read in this order:

1. docs/AI_CONTEXT.md
2. docs/PROJECT_SPEC.md
3. docs/PROJECT_STRUCTURE.md
4. docs/CODEMAP.md
5. docs/DEPENDENCIES.md
6. docs/ARCHITECTURE_DECISIONS.md
7. docs/DEVELOPMENT_WORKFLOW.md

===============================================================================
MISSION
===============================================================================

Your mission is to extend HashCenter without changing its architecture.

Prefer extending existing code over creating new components.

Do not rewrite working code unless explicitly requested.

===============================================================================
ARCHITECTURE
===============================================================================

Browser
    ↓
Handler
    ↓
Service
    ↓
Repository
    ↓
Database
    ↓
Hashcat

Never bypass layers.

===============================================================================
RESPONSIBILITIES
===============================================================================

Handler:
- HTTP
- Validation
- Response formatting

Service:
- Business logic
- Workflow
- Orchestration

Repository:
- Database only

Hashcat:
- External process execution

Scanner:
- Capture analysis

===============================================================================
FORBIDDEN
===============================================================================

Handler -> SQL

Handler -> Hashcat

Repository -> HTTP

Repository -> HTML

Repository -> Gin

Scanner -> Database

Hashcat -> HTTP

Hashcat -> HTML

Circular dependencies

Utility packages with vague names:
- helper
- helpers
- utils
- common
- generic
- manager

===============================================================================
WHEN WRITING CODE
===============================================================================

Always:

- Search for an existing implementation.
- Reuse existing services.
- Reuse repositories.
- Minimize the number of modified files.
- Avoid duplicated logic.

===============================================================================
CHECKLIST
===============================================================================

Before finishing:

go fmt ./...

go vet ./...

go test ./...

go build ./...

Update documentation if architecture changed.

===============================================================================
WHEN UNSURE
===============================================================================

Never guess.

Ask the developer.

===============================================================================
MAIN RULE
===============================================================================

The best change is the smallest correct change.
