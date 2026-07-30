#!/usr/bin/env bash
set -e

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
MASTER="$ROOT/ai/MASTER_RULES.md"

cp "$MASTER" "$ROOT/AGENTS.md"
cp "$MASTER" "$ROOT/CLAUDE.md"
cp "$MASTER" "$ROOT/GEMINI.md"

mkdir -p "$ROOT/.github"
cp "$MASTER" "$ROOT/.github/copilot-instructions.md"

cp "$MASTER" "$ROOT/.cursorrules"

echo
echo "AI rule files updated."
