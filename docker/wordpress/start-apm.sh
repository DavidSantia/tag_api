if ! [ -z $NEW_RELIC_LOG_LEVEL ]; then
   set -i -e "/^;newrelic.loglevel/s+info+$NEW_RELIC_LOG_LEVEL+" \
       /usr/local/etc/php/conf.d/newrelic.ini
fi

if ! [ -z $NEW_RELIC_LICENSE_KEY ]; then
   sed -i -e "s+REPLACE_WITH_REAL_KEY+$NEW_RELIC_LICENSE_KEY+" \
       /usr/local/etc/php/conf.d/newrelic.ini

   echo "Enabling agent $PLUGIN_VERSION"
fi
