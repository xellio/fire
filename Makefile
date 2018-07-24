TARGET = fire
GO      = go
BINDIR = ./bin/

all: $(TARGET)

$(TARGET): build
	upx --brute $(BINDIR)$@

build: clean $(BINDIR)
	$(GO) build -ldflags="-s -w" -o $(BINDIR)$(TARGET) ./cmd/main.go

clean:
	rm -f $(BINDIR)*

$(BINDIR):
	mkdir -p $(BINDIR)