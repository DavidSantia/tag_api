FROM debian:stretch

ENV AGENT_VERSION=newrelic-infra

RUN apt-get update
RUN apt-get install -y gnupg curl insserv
RUN echo "deb http://apt.newrelic.com/infrastructure_agent/linux/apt wheezy main" > /etc/apt/sources.list.d/newrelic.list

RUN curl -s http://apt.newrelic.com/infrastructure_agent/gpg/newrelic-infra.gpg | apt-key add -
RUN apt-get update
RUN apt-get install -y ${AGENT_VERSION}

ENV NEW_RELIC_LICENSE_KEY="" NEW_RELIC_LOG_LEVEL="info"

# Configure agent settings in /etc/newrelic-infra.yml
ADD newrelic-infra.yml /etc/
ADD start-agent.sh /usr/local/bin/
ADD nr-mysql /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/start-agent.sh"]
