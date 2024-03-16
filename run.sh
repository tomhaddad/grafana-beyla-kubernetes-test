#!/bin/bash
source .env
docker build -t go-grpc-service -f services/go-grpc-service/Dockerfile .
docker build -t node-grpc-service -f services/node-grpc-service/Dockerfile .
docker build -t node-http-service -f services/node-http-service/Dockerfile .
kind load docker-image go-grpc-service --name=$CLUSTER_NAME
kind load docker-image node-grpc-service --name=$CLUSTER_NAME
kind load docker-image node-http-service --name=$CLUSTER_NAME
kubectl apply -f ./kubernetes
envsubst < ./kubernetes/beyla/config.yaml | kubectl apply -f ./kubernetes/beyla