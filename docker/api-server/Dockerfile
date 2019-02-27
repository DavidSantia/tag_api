FROM scratch

ADD api-server /
ADD rootca.tar /

ENTRYPOINT ["/api-server", "-nhost", "tagdemo-nats", "-dbhost", "tagdemo-mysql", "-dbload"]
