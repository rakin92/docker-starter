## Container for React App

#### Command

Build docker container:
```
docker build -t APP_CONTAINER_NAME .
```

Run container:
```
docker run -it -p 3000:3000 --rm APP_CONTAINER_NAME
```

Using Docker Compose:
```
docker-compose -f docker-compose.yml up -d --build
```

---

If you’re using React Router, then you’ll need to use the custom Nginx config at production build step.

```
COPY nginx/nginx.conf /etc/nginx/nginx.conf
COPY nginx/default.conf /etc/nginx/conf.d/default.conf
```

To get rid of use of custom nginx configuration, remove the two lines above from the Dockerfile.

---

If you wish to enable brotli use the following docker image:

[fholzer/nginx-brotli](https://hub.docker.com/r/fholzer/nginx-brotli/builds)

```
FROM fholzer/nginx-brotli:latest
```

---

#### Separating build step
If you find yoursef separating the build step in your CI. You can simplify the Dockerfile even more.

```
FROM nginx:1.16.0-alpine

COPY nginx/nginx.conf /etc/nginx/nginx.conf
COPY nginx/nginx-site.conf /etc/nginx/conf.d/default.conf

WORKDIR /var/www/app

COPY build/ .
```
