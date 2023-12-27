FROM alpine:3.17
WORKDIR /root

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.17.0/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

ADD ./migration.sh .
RUN chmod +x migration.sh

ADD ./migrations/*.sql ./migrations/

ENTRYPOINT ["bash", "migration.sh"]