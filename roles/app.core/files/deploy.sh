#!/bin/bash

if [ -z $1 ]; then
	REF="origin/master"
else
	REF=$1
fi

set -ex

cd app
rm -rf *
git fetch
git reset --hard $REF
make

sudo systemctl restart app
