GO := go

BINARY := runny
PREFIX ?= /usr/local

.PHONY: build
build:
	$(GO) build -o $(BINARY) .

.PHONY: install
install: build
	install -Dm755 $(BINARY) $(DESTDIR)$(PREFIX)/bin/$(BINARY)

.PHONY: fmt
fmt:
	$(GO) fmt ./...

.PHONY: vet
vet:
	$(GO) vet ./...

.PHONY: clean
clean:
	$(RM) $(BIN)
