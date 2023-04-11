FROM alpine as builder

RUN apk add --no-cache just go
COPY . /app
WORKDIR /app
RUN just build

FROM alpine 
COPY --from=builder /app/build/polygon-fetcher /app/polygon-fetcher
WORKDIR /app
VOLUME [ "/data" ]
ENV POLYGON_API_KEY=""
ENTRYPOINT [ "/app/polygon-fetcher", "-output_dir", "/data", "-port", "80" ]
