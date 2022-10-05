all: clean compile-darwin-arm64 compile-darwin-amd64 compile-linux-arm64 compile-linux-amd64 compile-freebsd-arm64 compile-freebsd-amd64

compile-darwin-arm64:
	CGO_ENABLED=0 && GOOS=darwin GOARCH=arm64 go build -o ./bin/gvm_darwin_arm64 --tags netgo -ldflags '-extldflags "-static" -s -w' main.go

compile-darwin-amd64:
		CGO_ENABLED=0 && GOOS=darwin GOARCH=amd64 go build -o ./bin/gvm_darwin_amd64 --tags netgo -ldflags '-extldflags "-static" -s -w' main.go

compile-linux-arm64:
		CGO_ENABLED=0 && GOOS=linux GOARCH=arm64 go build -o ./bin/gvm_linux_arm64 --tags netgo -ldflags '-extldflags "-static" -s -w' main.go

compile-linux-amd64:
		CGO_ENABLED=0 && GOOS=linux GOARCH=amd64 go build -o ./bin/gvm_linux_amd64 --tags netgo -ldflags '-extldflags "-static" -s -w' main.go

compile-freebsd-arm64:
		CGO_ENABLED=0 && GOOS=freebsd GOARCH=arm64 go build -o ./bin/gvm_freebsd_arm64 --tags netgo -ldflags '-extldflags "-static" -s -w' main.go

compile-freebsd-amd64:
		CGO_ENABLED=0 && GOOS=freebsd GOARCH=amd64 go build -o ./bin/gvm_freebsd_amd64 --tags netgo -ldflags '-extldflags "-static" -s -w' main.go

clean:
	if [ -d "./bin" ]; then rm -rf ./bin; fi