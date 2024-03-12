import { createChannel, createClientFactory } from 'nice-grpc';
import { NodeGrpcServiceClient, NodeGrpcServiceDefinition } from 'demo-proto/dist/node/v1/node';

export const nodeGrpcClient: NodeGrpcServiceClient = createClientFactory().create(
  NodeGrpcServiceDefinition,
  createChannel(`${process.env['NODE_SERVICE_HOST'] ?? 'node-grpc-service'}`),
);
