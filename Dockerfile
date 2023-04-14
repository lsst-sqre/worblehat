FROM golang as build-stage
RUN mkdir -p /opt/worblehat
WORKDIR /opt/worblehat
COPY go/main.go go/Makefile /opt/worblehat/
RUN CGO_ENABLED=0 GOOS=linux make

#FROM gcr.io/distroless/base-debian11
FROM busybox
WORKDIR /
COPY --from=build-stage /opt/worblehat/worblehat /worblehat
RUN mkdir -p /srv/worblehat-adam
RUN echo bar > /srv/worblehat-adam/foo
RUN chown -R 3000001:3000001 /srv/worblehat-adam

ENTRYPOINT [ "/worblehat" ]


