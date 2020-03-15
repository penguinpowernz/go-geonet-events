VERSION=$(shell git describe --tags|tr -d 'v')

build:
	./scripts/embed_index.sh
	go build -o bin/quakepub ./cmd/quakepub

pkg:
	mkdir -p dpkg/bionic/usr/bin
	cp bin/quakepub dpkg/bionic/usr/bin
	IAN_DIR=dpkg/bionic ian set -v ${VERSION}
	IAN_DIR=dpkg/bionic ian pkg