#!/bin/sh

for file in $(git diff --name-only HEAD pkg/domain); do
  go generate "${file}"
done;
