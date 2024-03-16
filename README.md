## grafana-beyla-kubernetes-test

Testing out Grafana Beyla, with the following:

- 2x Go HTTP services
- NodeJS gRPC service
- NodeJS HTTP service

## Prerequisites
1. You have a local [Kind](https://kind.sigs.k8s.io/) cluster running
2. Docker is installed
3. Grafana Cloud account, with an [OpenTelemetry endpoint configured](https://grafana.com/docs/grafana-cloud/send-data/otlp/send-data-otlp/)

## Usage
These instructions are specific to running a local [Kind](https://kind.sigs.k8s.io/) cluster, and assume it is already running. 

1. Configure a `.env` file, based on the example `.env.example`. 
2. Run `npm i && npm run build` for all 3 services in `/services`, and in `/proto`
3. At the root directory, run the bash script `./run.sh` to build the images and push them to the local cluster.