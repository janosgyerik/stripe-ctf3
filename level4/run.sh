#!/bin/sh -e

cd $(dirname "$0")

ts=$(date +%F-%H-%M)
./test/harness 2>err | tee out

logdir=tmp/logs/$ts
mkdir -p $logdir
cp err out $logdir
