## Go Service Docker

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
If you wish to use internal repository from private repos, you will need to include this as part of your your docker file. You will need to set your `PRIVATE_PRIVATE_ACCESS_KEY` in github to allow access to private repos using the key.

```
RUN mkdir ~/.ssh
RUN echo "$PRIVATE_PRIVATE_ACCESS_KEY" > ~/.ssh/id_rsa
RUN chmod 600 ~/.ssh/id_rsa
RUN echo "Host *" > ~/.ssh/config
RUN echo "   StrictHostKeyChecking no" >> ~/.ssh/config
RUN echo "User git" >> ~/.ssh/config
RUN echo "IdentitiesOnly yes" >> ~/.ssh/config
RUN git config --global --add url."git@github.customdomain.com:".insteadOf https://github.customdomain.com/
```
