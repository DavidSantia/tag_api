FROM wordpress

ENV MYSQL_ROOT_PASSWORD=my-secret-pw
ENV MYSQL_DATABASE=tagdemo
ENV PLUGIN_VERSION=newrelic-php5

RUN apt-get update
RUN apt-get install -y gnupg
RUN echo 'deb http://apt.newrelic.com/debian/ newrelic non-free' > /etc/apt/sources.list.d/newrelic.list
RUN curl -s https://download.newrelic.com/548C16BF.gpg | apt-key add -
RUN apt-get update
RUN apt-get install -y ${PLUGIN_VERSION}
RUN echo | newrelic-install install

ENV NEW_RELIC_LICENSE_KEY="" NEW_RELIC_LOG_LEVEL="info"

# Configure agent settings in /usr/local/etc/php/conf.d/newrelic.ini
ADD start-apm.sh /usr/local/bin/
RUN sed -i '/^exec/i\sh /usr/local/bin/start-apm.sh &' /usr/local/bin/docker-entrypoint.sh 
