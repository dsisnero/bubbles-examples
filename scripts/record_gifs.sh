#!/usr/bin/env bash
# Record the component example GIFs used by the Bubbles README.
#
# Requires: crystal, vhs, ttyd, ffmpeg (VHS dependencies).
# Outputs:  gifs/*.gif
set -euo pipefail

cd "$(dirname "$0")/.."
mkdir -p build gifs

examples=(
  spinner
  textinput
  textarea
  tree
  table
  progress_animated
  viewport
  list_default
  file_picker
  help
  paginator
  timer
  stopwatch
)

for example in "${examples[@]}"; do
  if [[ ! -x "build/${example}" ]]; then
    echo "Building ${example}..."
    crystal build -o "build/${example}" "src/${example}.cr"
  fi
done

for example in "${examples[@]}"; do
  echo "Recording ${example}..."
  vhs "tapes/${example}.tape"
done

echo "GIFs written to gifs/."
