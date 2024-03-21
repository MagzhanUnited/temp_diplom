# Use a Node.js image as base
FROM node:16.15.0

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 2003

CMD ["npm", "run", "serve"]