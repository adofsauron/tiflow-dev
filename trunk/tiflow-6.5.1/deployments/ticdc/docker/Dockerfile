FROM golang:1.19.1-alpine3.15 as builder
RUN apk add --no-cache git make bash
WORKDIR /go/src/sdbflow
COPY . .
ENV CDC_ENABLE_VENDOR=0
RUN make cdc

FROM alpine:3.15
RUN apk add --no-cache tzdata bash curl socat
COPY --from=builder /go/src/sdbflow/bin/cdc /cdc
EXPOSE 8300
CMD [ "/cdc" ]
