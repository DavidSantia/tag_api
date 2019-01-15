FROM scratch

ADD api-server /

ENTRYPOINT ["/api-server", "-nhost", "tagdemo-nats", "-dbhost", "tagdemo-db", "-dbload"]
