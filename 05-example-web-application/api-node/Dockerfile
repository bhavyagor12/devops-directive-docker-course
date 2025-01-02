FROM node:19.6.0-alpine

COPY package.json ./

RUN npm install

COPY . .

CMD [ "npm", "run","dev" ]

