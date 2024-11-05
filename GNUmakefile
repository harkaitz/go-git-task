.POSIX:
.SUFFIXES:
.PHONY: all clean install check

PROJECT    =git-task
VERSION    =1.0.5
PREFIX     =/usr/local
BUILDDIR  ?=.build
UNAME_S   ?=$(shell uname -s)
EXE       ?=$(shell uname -s | awk '/Windows/ || /MSYS/ || /CYG/ { print ".exe" }')
TOOLCHAINS =x86_64-w64-mingw32 x86_64-linux-musl

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
## -- BLOCK:release --
release:
	mkdir -p $(BUILDDIR)
	hrelease -w github -t "$(TOOLCHAINS)" -N $(PROJECT) -R $(VERSION) -o $(BUILDDIR)/Release
	gh release create v$(VERSION) $$(cat $(BUILDDIR)/Release)
## -- BLOCK:release --
