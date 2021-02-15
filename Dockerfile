# layer simple to only get dafiti certificat
FROM gfgit/dafiti-certificates:latest as certs

# this is just a common layer to use local and builds
FROM golang:1.15-alpine3.13 as base
RUN apk --no-cache update && apk add --no-cache git ca-certificates
# ADD https://s3.amazonaws.com/dft-live-us-east-1-ca-root-certificates/dafiti_local.crt /etc/ssl/certs/dafiti_local.crt
COPY --from=certs /dafiti_local.crt /etc/ssl/certs/
RUN update-ca-certificates

# exclusive layer to vscode devcontainer
FROM base as vscodedevcontainer
ARG EXTERNAL_USERGROUP=1001
ARG EXTERNAL_USERID=1001
ENV BASH_FILE=/home/vscode/.bash_profile
RUN echo "uid: ${EXTERNAL_USERID} and gid: ${EXTERNAL_USERGROUP}"
RUN apk add --no-cache openssh
RUN addgroup -S vscode -g ${EXTERNAL_USERGROUP} && adduser -S vscode -u ${EXTERNAL_USERID} -G vscode
RUN mv /etc/profile.d/color_prompt /etc/profile.d/color_prompt.sh
USER vscode
RUN mkdir -p /tmp/gotools && cd /tmp/gotools && go mod init tmp/tools
COPY .devcontainer/go-tools-install.sh /tmp/gotools
RUN cd /tmp/gotools && sh go-tools-install.sh && rm go-tools-install.sh


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
