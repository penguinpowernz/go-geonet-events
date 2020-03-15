#!/bin/bash

html="$(cat index.html)"

cat <<EOF >index.html.go
package events

var version string

var indexHTML = []byte(\`$html\`)
EOF

sed 's/<\/h1>/ v\`\+version\+\`<\/h1>/' -i index.html.go

gofmt -w index.html.go