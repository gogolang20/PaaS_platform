FROM golang:1.19.1 as builder

ADD . /workspace
WORKDIR /workspace

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN go build -o main ./cmd/main.go

# use a minimal alpine image
FROM alpine:3.17

# set working directory
WORKDIR /go/bin

COPY --from=builder /workspace/main .

USER 1001
# run the binary
CMD ["./main"]
