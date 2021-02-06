# this is just a common layer to use local and builds
FROM golang:1.15-alpine3.13 as base
RUN apk --no-cache update && apk add --no-cache git ca-certificates
ADD https://s3.amazonaws.com/dft-live-us-east-1-ca-root-certificates/dafiti_local.crt /etc/ssl/certs/dafiti_local.crt
RUN update-ca-certificates

# this layer is reponsable to execute tests in cicle-ci
FROM base as ci
WORKDIR /app/
COPY . .

# the build layer responsable to create the entrypoint
FROM ci as builder
WORKDIR /app/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o entrypoint

# the shipment layer
FROM scratch
WORKDIR /
COPY --from=base /etc/ssl/certs /etc/ssl/certs/
COPY --from=builder /app/entrypoint .
COPY --from=builder /app/rev.txt .
EXPOSE 8080
ENTRYPOINT [ "/entrypoint" ]
