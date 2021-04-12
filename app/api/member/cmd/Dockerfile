FROM golang:alpine AS builder

LABEL stage=gobuilder

ARG git_user_name
ARG git_user_token
ARG project_path

WORKDIR /project

# Copy go mod and sum files
COPY go.mod go.sum ./

RUN git config --global url."https://${git_user_name}:${git_user_token}@git.private.tv/".insteadOf "https://git.private.tv/" \
    && GO111MODULE="on" \
    && GOPRIVATE="git.private.tv" \
    && GOSUMDB="off" \
    && GOPROXY="https://goproxy.cn,direct" \
    && time go mod download

COPY .. .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o main ${project_path}

FROM alpine:3.10

RUN apk update --no-cache
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder  /project/main /app/
RUN mkdir -p /data/

CMD ["/app/main", "-f", "etc/config.yaml"]
