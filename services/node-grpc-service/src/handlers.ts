import { ForwardRequest, ForwardResponse, ReceiveRequest, ReceiveResponse } from 'proto/dist/node/v1/node';

export const receive = async (request: ReceiveRequest): Promise<ReceiveResponse> => {
  return {
    message: `Received: '${request.message}'`,
  };
};

export const forward = async (request: ForwardRequest): Promise<ForwardResponse> => {
  return {
    message: `Forwareded '${request.message}' to ${request.to}`,
  };
};
