FROM --platform=linux/amd64 node:16-buster-slim  as build-stage
WORKDIR /app
RUN apt-get update || : && apt-get install python -y
RUN apt-get install build-essential -y
RUN apt-get install bash -y
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build


# etapa de producción
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
