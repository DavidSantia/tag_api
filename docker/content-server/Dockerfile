FROM scratch

ADD content-server.tar /

ENTRYPOINT ["/content-server", "-dbhost", "tagdemo-db", "-nhost", "nats://tagdemo-nats:4222"]
