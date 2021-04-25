BINARIES = false true yes basename cat

all: $(BINARIES)

$(BINARIES):
	go build -o bin $@/$@.go
