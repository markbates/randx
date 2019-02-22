#! /bin/bash

set -xe

echo $(modulePath)
mkdir -p $GOBIN
mkdir -p $GOPATH/pkg
mkdir -p $(modulePath)
pwd
ls -la
echo $GOPATH
