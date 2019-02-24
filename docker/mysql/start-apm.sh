if ! [ -z $NEW_RELIC_LICENSE_KEY ]; then
   cd /usr/local/newrelic_mysql_plugin-$PLUGIN_VERSION
   sed config/newrelic.template.json \
       -e "s+YOUR_LICENSE_KEY_HERE+$NEW_RELIC_LICENSE_KEY+" > config/newrelic.json
   sed config/plugin.template.json \
       -e "s+Localhost+$AGENT_NAME+" \
       -e "s+USER_NAME_HERE+$AGENT_USER+" \
       -e "s+USER_PASSWD_HERE+$AGENT_PASSWD+" > config/plugin.json

   echo "[Entrypoint] Starting newrelic_mysql_plugin-$PLUGIN_VERSION"
   java -Xmx128m -jar plugin.jar
fi
