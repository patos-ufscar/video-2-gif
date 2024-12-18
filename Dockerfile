# ARG GO_VERSION

FROM docker.io/golang:1.23.3-alpine3.20 as builder
RUN apk --no-cache add make bash
WORKDIR /app
COPY . /app
RUN make service

FROM docker.io/alpine:3.20
RUN apk --no-cache add ca-certificates ffmpeg
COPY --from=builder /app/cmd/gif/gif /usr/bin/gif
CMD ["gif"]
