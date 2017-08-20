FROM mysql/mysql-server

WORKDIR /root

ENV MYSQL_ROOT_PASSWORD my-secret-pw
ENV MYSQL_DATABASE tagdemo

# Configure MySQL global settings
RUN echo "character_set_client = utf8" >> /etc/my.cf && \
    echo "character_set_connection = utf8" >> /etc/my.cf && \
    echo "character_set_database = utf8" >> /etc/my.cf && \
    echo "character_set_results = utf8" >> /etc/my.cf && \
    echo "character_set_server = utf8" >> /etc/my.cnf && \
    echo "sql_mode = STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION" >> /etc/my.cnf

# Initialize DB
ADD images-users.sql /docker-entrypoint-initdb.d/
