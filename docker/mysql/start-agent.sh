#!/bin/sh

if ! [ -z "$NEW_RELIC_LOG_LEVEL" ]; then
   if [ "$NEW_RELIC_LOG_LEVEL" = "debug" ]; then
       sed -i -e "/^verbose/s+0+1+" \
           /etc/newrelic-infra.yml
   fi
fi

if ! [ -z $NEW_RELIC_LICENSE_KEY ]; then
   sed -i -e "/^license_key/s+your_license_key+$NEW_RELIC_LICENSE_KEY+" \
       /etc/newrelic-infra.yml

   echo "[Entrypoint] Enabling agent newrelic-infra $AGENT_VERSION"
   newrelic-infra start
fi
