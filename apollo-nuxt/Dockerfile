FROM node:19-alpine

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 2001

CMD ["npm", "run", "dev"]