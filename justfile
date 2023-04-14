build:
  go build -ldflags "-s -w" -o polygon-fetcher

test:
  go test -v ./utils

clean:
  rm polygon-fetcher

