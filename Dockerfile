FROM golang:1.18-alpine as builder
WORKDIR /app
COPY . .
RUN apk update && apk add make git
RUN make build

FROM alpine:3.16.1
WORKDIR /app/
ADD templates/* /app/templates/
ADD static/ /app/static/
RUN apk update && apk add git
RUN apk add kubectl --repository=http://dl-cdn.alpinelinux.org/alpine/edge/testing/
COPY --from=builder /app/kube-applier /app/kube-applier
ENTRYPOINT [ "/app/kube-applier" ]
