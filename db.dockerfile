FROM postgres:10.4-alpine

ADD ./migrations/db.sql /docker-entrypoint-initdb.d/database.sql

ENV POSTGRES_USER=postgres 
ENV POSTGRES_DBNAME=test_db 
ENV POSTGRES_PASSWORD=1234 
ENV POSTGRES_PORT=5432

EXPOSE 5432
