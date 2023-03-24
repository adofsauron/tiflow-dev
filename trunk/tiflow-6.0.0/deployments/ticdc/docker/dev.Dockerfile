FROM golang:1.18.0-alpine3.15 as builder
RUN apk add --no-cache git make bash
WORKDIR /go/src/sdbflow
COPY . .
ENV CDC_ENABLE_VENDOR=1
RUN go mod vendor
RUN make failpoint-enable
RUN make cdc
RUN make failpoint-disable

FROM alpine:3.15
RUN apk add --no-cache tzdata bash curl socat
COPY --from=builder /go/src/sdbflow/bin/cdc /cdc
EXPOSE 8300
CMD [ "/cdc" ]

