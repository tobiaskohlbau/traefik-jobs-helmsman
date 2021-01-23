#!/bin/bash

CA_ROOT_PATH=$(mkcert -CAROOT)
base64_opts=""
if [ "$(uname)" != "Darwin" ]; then
  base64_opts="-w0"
fi
CAROOT=$(cat "${CA_ROOT_PATH}"/rootCA.pem | base64 ${base64_opts})
sed "s/{{{CA_BUNDLE}}}/${CAROOT}/g" webhook.yaml | kubectl apply -f -
