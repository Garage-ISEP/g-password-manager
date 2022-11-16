build:
	$(foreach file, $(wildcard */*/*.go), GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../dist/$(basename $(file)) $(file);)