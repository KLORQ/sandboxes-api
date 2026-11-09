#!/usr/bin/env bash
set -euo pipefail

root=$(cd "$(dirname "$0")/.." && pwd)
tmp=$(mktemp -d "${TMPDIR:-/tmp}/rheic-pretag-check.XXXXXX")
trap 'rm -rf "$tmp"' EXIT

git -c init.defaultBranch=main init --bare "$tmp/origin.git" >/dev/null
git init --initial-branch=main "$tmp/checkout" >/dev/null
git -C "$tmp/checkout" config user.name 'Rheic Release Test'
git -C "$tmp/checkout" config user.email 'release-test@example.invalid'
mkdir -p "$tmp/checkout/scripts"
cp "$root/scripts/validate-release-candidate.sh" "$tmp/checkout/scripts/"
cp "$root/scripts/extract-release-notes.sh" "$tmp/checkout/scripts/"
chmod +x "$tmp/checkout/scripts/"*.sh
cat > "$tmp/checkout/CHANGELOG.md" <<'CHANGELOG'
# Changelog

## v0.0.1 — 2026-11-09

Synthetic release notes.

## v0.0.2 — 2026-11-09

Synthetic existing-tag notes.
CHANGELOG
git -C "$tmp/checkout" add CHANGELOG.md scripts
git -C "$tmp/checkout" commit -m 'Add synthetic candidate fixture' >/dev/null
git -C "$tmp/checkout" remote add origin "$tmp/origin.git"
git -C "$tmp/checkout" push -u origin main >/dev/null 2>&1
sha=$(git -C "$tmp/checkout" rev-parse HEAD)
helper="$tmp/checkout/scripts/validate-release-candidate.sh"

if ! (cd "$tmp/checkout" && GITHUB_REF=refs/heads/main GITHUB_SHA="$sha" "$helper" 0.0.1); then
  printf 'valid main SHA without tag was rejected\n' >&2
  exit 1
fi
expect_failure() {
  local name=$1
  shift
  if (cd "$tmp/checkout" && env "$@" "$helper" 0.0.1) >/dev/null 2>&1; then
    printf 'pre-tag validation accepted negative control: %s\n' "$name" >&2
    exit 1
  fi
}
expect_failure wrong-ref GITHUB_REF=refs/heads/feature/release GITHUB_SHA="$sha"
expect_failure wrong-sha GITHUB_REF=refs/heads/main GITHUB_SHA=0000000000000000000000000000000000000000
if (cd "$tmp/checkout" && GITHUB_REF=refs/heads/main GITHUB_SHA="$sha" "$helper" 0.0.x) >/dev/null 2>&1; then
  printf 'pre-tag validation accepted invalid version syntax\n' >&2
  exit 1
fi

git -C "$tmp/checkout" remote set-url origin "$tmp/missing-origin.git"
if (cd "$tmp/checkout" && GITHUB_REF=refs/heads/main GITHUB_SHA="$sha" "$helper" 0.0.1) >/dev/null 2>&1; then
  printf 'pre-tag validation accepted an unavailable origin\n' >&2
  exit 1
fi
git -C "$tmp/checkout" remote set-url origin "$tmp/origin.git"

git -C "$tmp/checkout" tag v0.0.2
git -C "$tmp/checkout" push origin refs/tags/v0.0.2 >/dev/null 2>&1
if (cd "$tmp/checkout" && GITHUB_REF=refs/heads/main GITHUB_SHA="$sha" "$helper" 0.0.2) >/dev/null 2>&1; then
  printf 'pre-tag validation accepted an existing remote tag\n' >&2
  exit 1
fi

