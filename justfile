build: test
  go build -ldflags "-s -w" -o build

test:
  go test -v ./utils
