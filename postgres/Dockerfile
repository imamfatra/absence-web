FROM postgres:17.4

ENV POSTGRES_USER=admin
ENV POSTGRES_PASSWORD=admin123
ENV POSTGRES_DB=absensi

COPY ./init.sql /docker-entrypoint-initdb.d/

EXPOSE 5432
