#!/bin/bash

html="$(cat index.html)"

cat <<EOF >index.html.go
package events

var indexHTML = []byte(\`$html\`)
EOF

gofmt -w index.html.go