set dotenv-load

build:
  go build -ldflags "-s -w" -o polygon-fetcher

test:
  go test -v ./utils

clean:
  rm polygon-fetcher

serve:
  ./polygon-fetcher -output_dir ./build/output -port 80
