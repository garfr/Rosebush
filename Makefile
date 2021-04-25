BINARIES = false true yes basename

all: $(BINARIES)

$(BINARIES):
	go build -o bin $@/$@.go
