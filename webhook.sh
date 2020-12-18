#!/bin/bash

CA_ROOT_PATH=$(mkcert -CAROOT)
CAROOT=$(cat ${CA_ROOT_PATH}/rootCA.pem | base64 -w0)
sed "s/{{{CA_BUNDLE}}}/${CAROOT}/g" webhook.yaml | kubectl apply -f -
