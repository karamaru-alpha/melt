#!/bin/sh

for file in $(git diff --name-only HEAD pkg); do
  go generate "${file}"
done;
