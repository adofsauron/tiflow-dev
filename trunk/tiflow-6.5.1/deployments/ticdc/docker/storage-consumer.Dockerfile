FROM golang:1.19.1-alpine3.15 as builder
RUN apk add --no-cache make bash
WORKDIR /go/src/sdbflow
COPY . .

RUN make storage_consumer

FROM alpine:3.15
COPY --from=builder  /go/src/sdbflow/bin/cdc_storage_consumer /cdc_storage_consumer

ENTRYPOINT ["tail", "-f", "/dev/null"]
