set dotenv-load

run:
  ./build/polygon-fetcher

build:
  go build -ldflags "-s -w" -o build

