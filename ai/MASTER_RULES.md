# HashCenter AI Master Rules

Version: R0.1

===============================================================================
PURPOSE
===============================================================================

This document is the single source of truth for every AI assistant
working on HashCenter.

All AI-specific instruction files must be generated from this document.

Do not maintain multiple independent rule sets.

===============================================================================
PROJECT
===============================================================================

Project Name:

HashCenter

Purpose:

A management platform for Hashcat.

Hashcat is the computation engine.

HashCenter controls everything else.

===============================================================================
MISSION
===============================================================================

Extend the project.

Do not redesign it.

Do not rewrite working code.

Preserve architecture.

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
CORE PRINCIPLES
===============================================================================

Reuse existing code.

Minimize changes.

Avoid duplication.

Preserve backward compatibility.

One responsibility per component.

===============================================================================
FORBIDDEN
===============================================================================

Handler → SQL

Handler → Hashcat

Repository → HTTP

Repository → HTML

Repository → Gin

Scanner → Database

Hashcat → HTTP

Hashcat → HTML

Circular dependencies.

===============================================================================
WORKFLOW
===============================================================================

Read documentation.

Understand task.

Find existing implementation.

Design minimal change.

Implement.

Run tests.

Update documentation.

Commit.

===============================================================================
CHECKS
===============================================================================

go fmt ./...

go vet ./...

go test ./...

go build ./...

===============================================================================
WHEN UNSURE
===============================================================================

Never guess.

Ask the developer.

===============================================================================
MAIN RULE
===============================================================================

The best solution is the smallest correct solution.
