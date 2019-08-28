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

If you wish to enable brotli use the following docker image:

[docker image: nginx-brotli](https://hub.docker.com/r/fholzer/nginx-brotli/builds)

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
