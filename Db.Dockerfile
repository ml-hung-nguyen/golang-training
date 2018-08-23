FROM postgres:10.4-alpine

ADD ./migration/db.sql /docker-entrypoint-initdb.d/database.sql

ENV POSTGRES_USER=postgres
ENV POSTGRES_DBNAME=simple
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_PORT=5432

EXPOSE 5432
