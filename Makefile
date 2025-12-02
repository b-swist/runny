BINARY := runny
PREFIX ?= /usr/local

.PHONY: build run install clean

build:
	go build -o $(BINARY) runny.go

run: build
	./$(BINARY)

install: build
	@install -Dm755 $(BINARY) ${DESTDIR}${PREFIX}/bin/${BINARY}

clean:
	rm -f $(BINARY)
