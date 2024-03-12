import { createChannel, createClientFactory } from 'nice-grpc';
import { GoGrpcServiceClient, GoGrpcServiceDefinition } from 'demo-proto/dist/go/v1/go';

export const goGrpcClient: GoGrpcServiceClient = createClientFactory().create(
  GoGrpcServiceDefinition,
  createChannel(`${process.env['GO_SERVICE_HOST'] ?? 'go-grpc-service'}`),
);
