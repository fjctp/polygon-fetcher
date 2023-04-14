FROM golang:1.20-alpine3.17 AS builder
WORKDIR /app
RUN apk add --no-cache just

FROM builder as build1
COPY . /app
RUN go get
RUN just build

FROM alpine:3.17 AS app
WORKDIR /app
COPY --from=build1 /app/build/polygon-fetcher /app/polygon-fetcher
VOLUME [ "/data" ]
ENV POLYGON_API_KEY=""
CMD [ "/app/polygon-fetcher", "-output_dir", "/data", "-port", "80" ]
