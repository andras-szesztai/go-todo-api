FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD=password

COPY ./init_db.sql /docker-entrypoint-initdb.d/

EXPOSE 3306