#!/bin/sh
set -eu

# Default param
IMAGE_NAME="tmp-friends/victo-api"
uid=1000
gid=1000
user=web

base=$(cd $(dirname $0); pwd)
docker image build -t $IMAGE_NAME:latest -f $base/Dockerfile . \
  --build-arg uid=$uid \
  --build-arg gid=$gid \
  --build-arg user=$user
