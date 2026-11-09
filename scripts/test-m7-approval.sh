#!/usr/bin/env bash
set -euo pipefail
root=$(cd "$(dirname "$0")/.." && pwd)
helper="$root/scripts/require-m7-approval.sh"
sha=0123456789abcdef0123456789abcdef01234567
"$helper" "$sha" "$sha"
expect_failure() {
  local label=$1
  shift
  if "$helper" "$@" >/dev/null 2>&1; then
    printf 'M7 approval accepted negative control: %s\n' "$label" >&2
    exit 1
  fi
}
expect_failure missing '' "$sha"
expect_failure malformed short "$sha"
expect_failure uppercase 0123456789ABCDEF0123456789ABCDEF01234567 "$sha"
expect_failure mismatch "$sha" 1123456789abcdef0123456789abcdef01234567
printf 'PASS: 1 positive and 4 negative M7 approval controls\n'
condition='if [[ ! "$M7_EXTERNAL_SECURITY_REVIEW_APPROVED_COMMIT" =~ ^[0-9a-f]{40}$ || "$M7_EXTERNAL_SECURITY_REVIEW_APPROVED_COMMIT" != "$GITHUB_SHA" ]]; then'
matched=0
for workflow in "$root"/.github/workflows/*.yml; do
  if grep -Fq 'M7_EXTERNAL_SECURITY_REVIEW_APPROVED_COMMIT' "$workflow"; then
    grep -Fq "$condition" "$workflow"
    grep -Fq 'M7_EXTERNAL_SECURITY_REVIEW_APPROVED_COMMIT: ${{ vars.M7_EXTERNAL_SECURITY_REVIEW_APPROVED_COMMIT }}' "$workflow"
    grep -Fq 'environment: m7-external-security-review' "$workflow"
    matched=$((matched + 1))
  fi
done
[[ "$matched" -eq 1 ]]
printf 'PASS: commit-bound inline condition present in %s gated workflow\n' "$matched"
