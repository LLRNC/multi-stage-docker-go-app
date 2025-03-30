FROM golang:1.21 AS builder

WORKDIR /app
COPY main.go .
ARG VERSION=1.0.0
ENV VERSION=$VERSION

RUN go mod init app && go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o app .
RUN ./app > index.html

FROM nginx:alpine
COPY --from=builder /app/index.html /usr/share/nginx/html/index.html

HEALTHCHECK --interval=30s --timeout=5s --start-period=5s   CMD wget --spider -q http://localhost:80 || exit 1
