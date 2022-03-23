# syntax=docker/dockerfile:1
FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk update --no-cache && apk add git

WORKDIR /build
ADD go.mod .
ADD cert.pem /app/cert.pem
ADD key.pem /app/key.pem
#ADD go.sum .

RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /app/go-poc . 

FROM alpine
RUN apk update --no-cache && apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/go-poc /app/key.pem /app/cert.pem /app/
#COPY --from=builder /app/go-poc /app/go-poc
#COPY --from=builder /app/key.pem /app/key.pem
#COPY --from=builder /app/cert.pem /app/cert.pem


EXPOSE 8080
# create and set non-root USER
RUN addgroup -g 1001 appuser && \
    adduser -S -u 1001 -G appuser appuser
RUN chown -R appuser:appuser /app && \
    chmod 755 /app
USER appuser

ENTRYPOINT ["/app/go-poc"]

#CMD [". /go-poc"]