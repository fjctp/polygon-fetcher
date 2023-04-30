set dotenv-load

build:
  go build -ldflags "-s -w" -o polygon-fetcher

test opts="":
  go test {{opts}} ./utils
  go test {{opts}} ./tickerData

clean:
  rm polygon-fetcher

reset:
  rm -rf ./data

serve:
  ./polygon-fetcher -output_dir ./data -port 80
