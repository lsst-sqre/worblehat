FROM golang as build-stage
RUN mkdir -p /opt/worblehat
WORKDIR /opt/worblehat
COPY go/main.go go/Makefile /opt/worblehat/
RUN CGO_ENABLED=0 GOOS=linux make

FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build-stage /opt/worblehat/worblehat /worblehat
USER nonroot:nonroot
ENTRYPOINT [ "/worblehat" ]


