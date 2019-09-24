#!/usr/bin/env bash

set -ex

export TAG="$USER"
export HUB="istio"

make docker

kubectl create namespace istio-system || :

helm template \
  --name istio-init \
  --namespace istio-system \
  --set global.hub="$HUB" \
  --set global.tag="$TAG" install/kubernetes/helm/istio-init | kubectl apply --namespace istio-system -f-

sleep 10

helm template \
  --name istio \
  --namespace istio-system \
  --set global.hub="$HUB" \
  --set global.tag="$TAG" install/kubernetes/helm/istio | kubectl apply --namespace istio-system -f-
