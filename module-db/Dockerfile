FROM postgres:latest

COPY init.sql /docker-entrypoint-initdb.d/

ENV POSTGRES_PASSWORD=forecast
ENV POSTGRES_USER=forecast
ENV POSTGRES_DB=forecast
