#!/bin/bash

PROJECT_DIR="$(dirname $(dirname "$(readlink -f "$0")"))"
PREVIEW_FILE=$PROJECT_DIR/docs/PREVIEW.md

echo '# Preview' > $PREVIEW_FILE
echo '' >> $PREVIEW_FILE
echo 'This is sample of error log and parsed result' >> $PREVIEW_FILE
while IFS= read -r line; do
  echo '' >> $PREVIEW_FILE
  echo '```txt' >> $PREVIEW_FILE
  echo $line >> $PREVIEW_FILE
  echo '```' >> $PREVIEW_FILE
  echo '' >> $PREVIEW_FILE
  echo '```json' >> $PREVIEW_FILE
  $PROJECT_DIR/nginx-error-log-parser test --log "$line" | jq '.' >> $PREVIEW_FILE
  echo '```' >> $PREVIEW_FILE
  echo '' >> $PREVIEW_FILE
  echo '---' >> $PREVIEW_FILE
done < $PROJECT_DIR/samples/error.log

