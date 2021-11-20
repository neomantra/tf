#!/bin/bash -x

set -e

go test neomantra/gotf/internal/gotf

go build  -o gotf cmd/gotf/main.go
