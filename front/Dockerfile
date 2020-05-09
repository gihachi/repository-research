FROM node:10.17.0-alpine3.9
WORKDIR /workspace
COPY . /workspace
RUN apk update &&\
    npm install
CMD ["npm", "run", "dev"]