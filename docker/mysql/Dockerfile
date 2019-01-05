FROM mysql/mysql-server:5.7.23

WORKDIR /root

ENV MYSQL_ROOT_PASSWORD my-secret-pw
ENV MYSQL_DATABASE tagdemo

# SQL to Initialize DB
ADD grant.sql /docker-entrypoint-initdb.d/
ADD make-tables.sql /docker-entrypoint-initdb.d/
ADD sample-data.sql /docker-entrypoint-initdb.d/
