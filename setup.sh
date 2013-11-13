#!/bin/sh

# Link the git pre-commit hook
ln -s ../../scripts/pre-commit.sh .git/hooks/pre-commit

# this will install vet in $GOROOT, so make sure you can write there
# or have root.
go get code.google.com/p/go.tools/cmd/vet 
