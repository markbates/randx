#! /bin/bash

set -xe

echo $(modulePath)
mkdir -p $(GOBIN)
mkdir -p $(GOPATH)/pkg
mkdir -p $(modulePath)
pwd
ls -la
echo '##vso[task.prependpath]$(GOBIN)'
echo '##vso[task.prependpath]$(GOROOT)/bin'
