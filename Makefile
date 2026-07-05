BINARY := runny
PREFIX ?= /usr/local

.PHONY: build install clean

build:
	go build -o $(BINARY) main.go

install: build
	@install -Dm755 $(BINARY) ${DESTDIR}${PREFIX}/bin/${BINARY}

run:
	go run main.go

clean:
	rm -f $(BINARY)
