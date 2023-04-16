# polygon-fetcher

This program runs a static file server that serves HTML reports for stocks.

## Usage

The server listens to port `80` and serves HTML reports from `output_dir/html` folder. To request a report for `AAPL`, run the program and go to `localhost/AAPL.html` in a broswer. If the report does not exist, the program will pull data from `polygon.io` using REST APIs and generate a HTML report. The data will be stored in `output_dir/json` in JSON format, while the report will be stored in `output_dir/html` as HTML.

## Prerequisite

* `polygon.io` API key
* `just` >= 1.8.0
* `go` >= 1.19.6

## Run

### Build executable

Build the executable under `build` folder

```
just build
```

### Run the executabke

Run the server and set the output folder to `./output`, where the JSON data and html reports will locate in their sub-folders.

```
export POLYGON_API_KEY="Your API key"
polygon-fetcher -output_dir "output"
```

Or run it wil `just`, which set the output folder to `./build/output`

```
echo POLYGON_API_KEY="Your API key" > .env
just serve
```

### Use Docker

Define `POLYGON_API_KEY` in `.env`.

```
podman run -it --rm -p 3000:80 --env-file .env -v $PWD/output/:/data fjctp/polygon-fetcher:release
```
