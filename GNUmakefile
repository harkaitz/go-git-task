.POSIX:
.SUFFIXES:
.PHONY: all clean install check

PROJECT   =git-task
VERSION   =1.0.0
PREFIX    =/usr/local
BUILDDIR ?=.build
UNAME_S  ?=$(shell uname -s)
EXE      ?=$(shell uname -s | awk '/Windows/ || /MSYS/ || /CYG/ { print ".exe" }')

all:
clean:
install:
check:

## -- BLOCK:go --
.PHONY: all-go install-go clean-go $(BUILDDIR)/git-task$(EXE)
all: all-go
install: install-go
clean: clean-go
all-go: $(BUILDDIR)/git-task$(EXE)
install-go:
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp  $(BUILDDIR)/git-task$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f $(BUILDDIR)/git-task$(EXE)
##
$(BUILDDIR)/git-task$(EXE): $(GO_DEPS)
	mkdir -p $(BUILDDIR)
	go build -o $@ $(GO_CONF) ./cmd/git-task
## -- BLOCK:go --
