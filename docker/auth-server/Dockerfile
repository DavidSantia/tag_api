FROM scratch

ADD auth-server.tar /

ENTRYPOINT ["/auth-server", "-dbhost", "tagdemo-db", "-nhost", "nats://tagdemo-nats:4222"]
