FROM scratch

ADD api-server /

ENTRYPOINT ["/api-server", "-dbhost", "tagdemo-db", "-nhost", "tagdemo-nats"]
