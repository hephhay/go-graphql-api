FROM golang:1.17.5-alpine as debug

RUN mkdir -p /server \
    mkdir -p /var/log/app && \
    chmod 755 -R /var/log && \
    apk update && \
    apk add curl \
            git \
            bash \
            alpine-sdk \
            make && \
    rm -rf /var/cache/apk/*

RUN go get -u golang.org/x/tools/go/packages
#RUN go get github.com/go-delve/delve/cmd/dlv
### Run the Delve debugger ###
#COPY ./dlv.sh /
#RUN chmod +x /dlv.sh
#ENTRYPOINT [ "/dlv.sh"]

###########PRODUCTION IMAGE###################

FROM golang:1.15-alpine as prod
