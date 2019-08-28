## DataBase Dockers

Build docker container:
```
docker build -t DB_CONTAINER_NAME .
```

Run container:
```
docker run --rm DB_CONTAINER_NAME
```

Using Docker Compose:
```
docker-compose -f docker-compose.yml up -d --build
```

---

This starts a db containers with Postgres.

* Add your sql files
* Can create data bases using sql
* Can setup schemas using sql
