### Build using the Alpine Golang container.
FROM golang:alpine as builder

ENV CGO=0
ENV GOOS=linux

ADD . /go/src/github.com/influxdata/telegraf
WORKDIR /go/src/github.com/influxdata/telegraf

RUN set -eux; \
        apk update && \
        apk add make git

RUN make deps && \
    make test && \
    make telegraf

### Metrics container definition.
FROM alpine
RUN mkdir -p /etc/telegraf /metrics

RUN apk update && \
    apk add bash

COPY --from=builder /go/src/github.com/influxdata/telegraf/telegraf /bin
COPY etc/telegraf_recurly.conf /etc/telegraf/recurly.conf
COPY entrypoint.sh /entrypoint.sh
RUN chmod 0755 /entrypoint.sh

# These variables are used by entrypoint.sh to configure telegraf
# at run time.
ENV LOGGING_NAMESPACE      "dev-mode"
ENV GRAPHITE_SRV1          "metrics-relay2.ewr1.recurly.net:2004"
ENV GRAPHITE_SRV2          "metrics-relay1.ewr1.recurly.net:2004"
ENV TELEGRAF_PLUGIN_CONFIG ""
ENV TELEGRAF_CONFIGURE     "yes"
ENV TELEGRAF_CONFIG_PATH   "/metrics/telegraf.conf"

ENTRYPOINT ["/entrypoint.sh"]
