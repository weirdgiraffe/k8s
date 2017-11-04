#!/bin/sh
DIR=$(dirname "`readlink -f "$0"`")
VERSION=$(cat $DIR/Dockerfile | grep -Po "(?<=version=).*")
eval $(minikube docker-env)
docker image build -t httpecho:$VERSION $DIR
