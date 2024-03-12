FROM node:20-alpine AS installer
RUN apk add --no-cache git libc6-compat

WORKDIR /app
COPY ./proto/ ./proto
COPY ./services/node-grpc-service ./services/node-grpc-service
WORKDIR /app/proto
RUN npm install
RUN npm run build

WORKDIR /app/services/node-grpc-service
RUN npm install
RUN npm run build

FROM node:20-alpine AS server
WORKDIR /app

COPY proto/package*.json ./proto/
COPY --from=installer app/proto ./proto/

WORKDIR /app/services/node-grpc-service
COPY services/node-grpc-service/package*.json ./
COPY --from=installer app/services/node-grpc-service/dist ./dist
COPY --from=installer app/services/node-grpc-service/node_modules ./node_modules

EXPOSE 8000
ENTRYPOINT ["npm"]
CMD ["start"]
