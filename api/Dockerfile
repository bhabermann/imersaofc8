FROM node:lts-alpine

RUN apk add --no-cache bash

RUN npm install --location=global @nestjs/cli #could chose compatible version

USER node

WORKDIR /home/node/app