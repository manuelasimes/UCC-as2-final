FROM node:16

ENV NODE_ENV development

ADD . /frontend
WORKDIR /frontend

COPY package.json ./
COPY package-lock.json ./
COPY ./ ./
RUN npm i
CMD npm start