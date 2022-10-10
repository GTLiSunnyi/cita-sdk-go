#!/usr/bin/env bash

set -eo pipefail

proto_dirs=$(find ./protos -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  protoc \
  -I "protos" \
  --go_out=plugins=grpc:. \
  $(find "${dir}" -maxdepth 1 -name '*.proto')

done

# move proto files to the right places
cp -r github.com/GTLiSunnyi/cita-sdk-go/* ./
rm -fr github.com
