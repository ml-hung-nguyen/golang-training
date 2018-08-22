FROM postgres:10.4-alpine

# ADD ./migrations/db.sql /docker-entrypoint-initdb.d/database.sql

# ENV POSTGRES_USER=postgres 
# ENV POSTGRES_DBNAME=go_training 
# ENV POSTGRES_PASSWORD=mypassword 
# ENV POSTGRES_PORT=5432

EXPOSE 5432
