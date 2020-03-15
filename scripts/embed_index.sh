#!/bin/bash

html="$(cat index.html)"

cat <<EOF >index.html.go
package events

var version string

var indexHTML = []byte(\`$html\`)
EOF

sed 's/<\/title>/ v\`\+version\+\`<\/title>/' -i index.html.go

gofmt -w index.html.go