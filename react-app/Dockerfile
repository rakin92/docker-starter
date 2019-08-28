# build environment
FROM node:12.2.0-alpine as build
WORKDIR /var/www/app
ENV PATH /var/www/app/node_modules/.bin:$PATH
COPY package.json /var/www/app/package.json
RUN npm install --silent
RUN npm install react-scripts@3.0.1 -g --silent
COPY . /var/www/mysite
RUN npm run build

# production environment
FROM nginx:1.16.0-alpine
COPY --from=build /var/www/app/build /var/www/app
COPY nginx/nginx.conf /etc/nginx/nginx.conf
COPY nginx/default.conf /etc/nginx/conf.d/default.conf
EXPOSE 3000
CMD ["nginx", "-g", "daemon off;"]