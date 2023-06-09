FROM golang:1.19-alpine3.17 AS builder
WORKDIR /app
RUN apk add --no-cache just gcc libc-dev

FROM builder as build1
COPY . /app
RUN go get
RUN just test
RUN just build

FROM alpine:3.17 AS app
WORKDIR /app
COPY --from=build1 /app/polygon-fetcher /app/polygon-fetcher
VOLUME [ "/data" ]
ENV POLYGON_API_KEY=""
CMD [ "/app/polygon-fetcher", "-output_dir", "/data", "-port", "80" ]
