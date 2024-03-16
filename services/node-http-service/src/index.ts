import 'dotenv/config';
import Fastify, { FastifyInstance, FastifyRequest } from 'fastify';
import { nodeGrpcClient } from './clients/node';

export const app: FastifyInstance = Fastify({
  logger: {
    enabled: true,
    level: 'debug',
    transport: {
      target: 'pino-pretty',
      options: {
        colorize: true,
      },
    },
  },
});

app.post(
  '/send/:serviceName',
  async (
    request: FastifyRequest<{
      Params: {
        serviceName: string;
      };
      Body: {
        message: string;
      };
    }>,
  ) => {
    try {
      const { serviceName } = request.params;
      switch (serviceName) {
        case 'node-grpc-service':
          app.log.info('Received request for node-http-service');
          await nodeGrpcClient.receive({ message: request.body.message });
          return { status: 'success' };
        default:
          return { status: 'failure' };
      }
    } catch (e) {
      app.log.error(e);
      return;
    }
  },
);

app.post(
  '/forward/:fromServiceName/:toServiceName',
  async (
    request: FastifyRequest<{
      Params: {
        fromServiceName: string;
        toServiceName: string;
      };
    }>,
  ) => {
    try {
      const { fromServiceName, toServiceName } = request.params;
      console.log(fromServiceName, toServiceName);
    } catch (e) {
      app.log.error(e);
    }
  },
);

const start = async () => {
  try {
    await app.listen({ port: process.env['PORT'] ? Number(process.env['PORT']) : 8000, host: '0.0.0.0' });
    await app.ready();
  } catch (err) {
    app.log.error(err);
    process.exit(1);
  }
};

start();

const shutdown = async (signal: string) => {
  app.log.info(`${signal} received: shutting down server`);
  await app.close();
  process.exit(0);
};

process.on('SIGINT', shutdown);
process.on('SIGTERM', shutdown);
