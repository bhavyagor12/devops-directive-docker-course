FROM node:20.11.1-alpine3.19 AS build

WORKDIR /usr/src/app

COPY package*.json ./

RUN --mount=type=cache,target=/usr/src/app/.npm \
    npm set cache /usr/src/app/.npm && \
    npm ci

ENV PATH /app/node_modules/.bin:$PATH

COPY . .

RUN npm run build 

FROM nginx:1.25.4-alpine3.18

COPY ./nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /usr/src/app/dist /var/www/html/

EXPOSE 3000

ENTRYPOINT ["nginx","-g","daemon off;"]
