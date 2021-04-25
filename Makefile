BINARIES = false true yes basename cat head

all: $(BINARIES)

$(BINARIES):
	go build -o bin $@/$@.go
