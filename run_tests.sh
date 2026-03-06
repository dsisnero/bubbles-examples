#!/bin/bash
# Run tests in batches to avoid compilation issues

set -e

echo "Running tests in batches..."

# Find all spec files
SPEC_FILES=$(find spec -name "*_spec.cr" | sort)

# Count total specs
TOTAL=$(echo "$SPEC_FILES" | wc -l | tr -d ' ')
echo "Found $TOTAL spec files"

# Run in batches of 5
BATCH_SIZE=5
BATCH_NUM=1
FAILED=0

for i in $(seq 1 $BATCH_SIZE $TOTAL); do
  BATCH_FILES=$(echo "$SPEC_FILES" | tail -n +$i | head -n $BATCH_SIZE)

  if [ -z "$BATCH_FILES" ]; then
    break
  fi

  echo ""
  echo "=== Batch $BATCH_NUM ==="

  # Run the batch
  if crystal spec $BATCH_FILES 2>&1; then
    echo "✓ Batch $BATCH_NUM passed"
  else
    echo "✗ Batch $BATCH_NUM failed"
    FAILED=1
    # Continue with other batches even if one fails
  fi

  BATCH_NUM=$((BATCH_NUM + 1))
done

echo ""
if [ $FAILED -eq 0 ]; then
  echo "✅ All tests passed!"
  exit 0
else
  echo "❌ Some tests failed"
  exit 1
fi