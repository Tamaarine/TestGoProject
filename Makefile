EXEC := hello
SRC := src
BIN := bin
SOURCES := $(shell find $(SRC) -name '*.go')

all: $(BIN)/$(EXEC)
	mkdir -p $(BIN)

$(BIN)/$(EXEC): $(SOURCES)
	go build -o $(BIN)/$(EXEC) $^
