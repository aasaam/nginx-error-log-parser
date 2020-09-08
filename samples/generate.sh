#!/bin/bash

echo '# Preview nginx error log parser' > PREVIEW.md
while IFS= read -r line; do
  echo '' >> PREVIEW.md
  echo '```txt' >> PREVIEW.md
  echo $line >> PREVIEW.md
  echo '```' >> PREVIEW.md
  echo '' >> PREVIEW.md
  echo '```json' >> PREVIEW.md
  ../nginx-error-log-parser test --log "$line" | jq '.' >> PREVIEW.md
  echo '```' >> PREVIEW.md
  echo '' >> PREVIEW.md
  echo '---' >> PREVIEW.md
done < error.log

