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
COPY --from=build-stage /app/config/default.conf /etc/nginx/conf.d/default.conf
COPY --from=build-stage /app/config/nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
RUN chmod +x /usr/share/nginx/html/write.sh
CMD /usr/share/nginx/html/write.sh
