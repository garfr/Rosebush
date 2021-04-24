BINARIES = false \
					 true 

all: $(BINARIES)

$(BINARIES):
	go build -o bin $@/$@.go
