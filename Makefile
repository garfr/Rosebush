BINARIES = false true yes

all: $(BINARIES)

$(BINARIES):
	go build -o bin $@/$@.go
