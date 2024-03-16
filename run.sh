#!/bin/bash
source .env
docker build -t go-http-forwarder-service -f services/go-http-forwarder-service/Dockerfile .
docker build -t go-http-receiver-service -f services/go-http-receiver-service/Dockerfile .
docker build -t node-grpc-service -f services/node-grpc-service/Dockerfile .
docker build -t node-http-service -f services/node-http-service/Dockerfile .
kind load docker-image go-http-forwarder-service --name=$CLUSTER_NAME
kind load docker-image go-http-receiver-service --name=$CLUSTER_NAME
kind load docker-image go-grpc-service --name=$CLUSTER_NAME
kind load docker-image go-grpc-service-2 --name=$CLUSTER_NAME
kind load docker-image node-grpc-service --name=$CLUSTER_NAME
kind load docker-image node-http-service --name=$CLUSTER_NAME
kubectl apply -f ./kubernetes
kubectl apply -f ./kubernetes/beyla