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

   echo "Enabling agent $AGENT_VERSION"
   /etc/init.d/newrelic-infra start

   echo "Enabling integration nr-mysql"
   /usr/local/bin/nr-mysql -username newrelic -password welcome1 -hostname tagdemo-mysql
fi

while :; do sleep 3600; done
