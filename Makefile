build:
	go build -o build/smerch-cli ./cmd

run: build
	./build/smerch-cli