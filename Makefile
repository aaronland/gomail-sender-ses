GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

cli:
	go build -mod $(VENDOR) -ldflags="$(LDFLAGS)" -o bin/send cmd/send/main.go
