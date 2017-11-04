#!/bin/sh

ISTIO_DIR=$(find `pwd` -name "istio*" -type d)
if [ "$ISTIO_DIR" = "" ]
then
	curl -L https://git.io/getLatestIstio | sh -
	ISTIO_DIR=$(find `pwd` -name "istio*" -type d)
fi

export ISTIO_DIR=$ISTIO_DIR
export PATH=$PATH:$ISTIO_DIR/bin


