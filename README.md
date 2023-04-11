# polygon-fetcher

This program fetches financial data from `polgon.io` and generate a html report.

## How-to

Build the executable under `build` folder

```
just build
```

Run the server and set the output folder to `./output`, where the JSON data and html reports will locate in their sub-folders.

```
export POLYGON_API_KEY="Your API key"
polygon-fetcher -output_dir "output"
```

If running from docker, define `POLYGON_API_KEY` in `.env` and call 

```
podman run -it --rm -p 3000:80 --env-file .env -v $PWD/build/output/:/data fjctp/test:latest
```
