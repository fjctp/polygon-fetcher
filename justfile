set dotenv-load

build:
  go build -ldflags "-s -w" -o polygon-fetcher

test opts="":
  go test {{opts}} ./utils
  go test {{opts}} ./tickerData

clean:
  rm polygon-fetcher

serve:
  ./polygon-fetcher -output_dir ./build/output -port 80
