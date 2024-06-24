FROM node

WORKDIR /app

COPY package.json .
COPY package-lock.json ./

COPY ./ ./
RUN npm i
COPY . .
## EXPOSE [Port you mentioned in the vite.config file]
EXPOSE 3000

CMD ["npm", "run", "dev"]