FROM node:lts-alpine

WORKDIR /usr/src/legacy-api
COPY api-legacy/package.json .
COPY api-legacy/ .
RUN npm install --only=prod