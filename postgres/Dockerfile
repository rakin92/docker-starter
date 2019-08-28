FROM clkao/postgres-plv8:latest
COPY *.sql /docker-entrypoint-initdb.d/
HEALTHCHECK --interval=5s CMD ["docker-healthcheck"]
