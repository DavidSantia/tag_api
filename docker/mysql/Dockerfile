FROM mysql/mysql-server:5.7

ENV MYSQL_ROOT_PASSWORD=my-secret-pw
ENV MYSQL_DATABASE=tagdemo

# SQL to Initialize DB
ADD grant.sql /docker-entrypoint-initdb.d/
ADD make-tables.sql /docker-entrypoint-initdb.d/
ADD sample-data.sql /docker-entrypoint-initdb.d/

ENV AGENT_VERSION=1.2.13-1

RUN curl -s -o /etc/yum.repos.d/newrelic-infra.repo http://apt.newrelic.com/infrastructure_agent/linux/yum/el/7/x86_64/newrelic-infra.repo
RUN yum makecache -y -q --disablerepo='*' --enablerepo='newrelic-infra'
RUN yum install -y newrelic-infra

ENV NEW_RELIC_LICENSE_KEY="" NEW_RELIC_LOG_LEVEL="info"

# Configure linux agent settings in /etc/newrelic-infra.yml
ADD newrelic-infra.yml /etc/

# Configure mysql integration settings in /etc/newrelic-infra/integrations.d
ADD mysql-config.yml /etc/newrelic-infra/integrations.d/

ADD nr-mysql /var/db/newrelic-infra/newrelic-integrations/bin/
ADD mysql-definition.yml /var/db/newrelic-infra/newrelic-integrations/
ADD start-agent.sh /
RUN sed -i '/^exec/i\sh /start-agent.sh &' /entrypoint.sh 

