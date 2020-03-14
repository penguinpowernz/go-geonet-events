VERSION=$(shell git describe --tags|tr -d 'v')

build:
	go build -o bin/geonet-events ./cmd/geonet-events

pkg:
	mkdir -p dpkg/bionic/usr/bin
	cp bin/geonet-events dpkg/bionic/usr/bin
	IAN_DIR=dpkg/bionic ian set -v ${VERSION}
	IAN_DIR=dpkg/bionic ian pkg