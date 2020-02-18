CC=go
CFLAGS=-o ../program1

program1: src/program1.go src/types.go src/funcs.go
	cd src; \
	$(CC) build $(CFLAGS)