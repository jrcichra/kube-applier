FROM golang:1.21.1-alpine as builder
WORKDIR /app
COPY . .
RUN apk update && apk add make git
RUN make build

FROM alpine:3.18.3
WORKDIR /app/
ADD templates/* /templates/
ADD static/ /static/
RUN apk update && apk add git
RUN apk add kubectl --repository=http://dl-cdn.alpinelinux.org/alpine/edge/community/
COPY --from=builder /app/kube-applier /app/kube-applier
ENTRYPOINT [ "/app/kube-applier" ]
