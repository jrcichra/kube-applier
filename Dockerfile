FROM golang:1.18-alpine as builder
WORKDIR /app
COPY . .
RUN apk add make
RUN make build

FROM alpine:3.15
WORKDIR /app/
ADD templates/* /templates/
ADD static/ /static/
RUN apk add git
RUN apk add kubectl --repository=http://dl-cdn.alpinelinux.org/alpine/edge/testing/
COPY --from=builder /app/kube-applier /kube-applier
