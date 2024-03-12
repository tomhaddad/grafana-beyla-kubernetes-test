import 'dotenv/config';
import { Server, createServer } from 'nice-grpc';
import { pino } from 'pino';
import { NodeGrpcServiceDefinition } from 'demo-proto/dist/node/v1/node';
import { forward, receive } from './handlers';

export const logger = pino({
  transport: {
    target: 'pino-pretty',
    options: {
      colorize: true,
    },
  },
});

let server: Server;

const startup = async (): Promise<void> => {
  server = createServer();

  server.add(NodeGrpcServiceDefinition, {
    receive,
    forward,
  });

  server.listen(`0.0.0.0:${process.env['PORT'] ?? 8080}`).then(port => {
    logger.info(`Server listening on port ${port}`);
  });
};

const shutdown = async (signal: string): Promise<void> => {
  logger.info(`${signal} received: shutting down server`);
  await server.shutdown();
  logger.info('server closed');
  process.exit(0);
};

startup();

process.on('SIGINT', shutdown);
process.on('SIGTERM', shutdown);
