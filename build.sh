#!/usr/bin/env bash
echo "Building Macos version... "
GOOS=darwin GOARCH=amd64 go build -o as2make.darwin .
echo "Building Linux version... "
GOOS=linux GOARCH=amd64 go build -o as2make.linux .
echo "Building FreeBSD version... "
GOOS=freebsd GOARCH=amd64 go build -o as2make.freebsd .
