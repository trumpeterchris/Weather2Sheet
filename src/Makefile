include $(GOROOT)/src/Make.inc
PROGS =  wu
PREFIX = $(GOBIN)

all: $(PROGS)

wu:	wu.go
	$(GC) -o wu.$(O) wu.go
	$(LD) -o wu wu.$(O)

format:
	gofmt -tabwidth=2 -w wu.go

clean:
	rm -f $(PROGS)
	rm -f wu.$(O)

install:
	cp wu $(PREFIX)

uninstall:
	rm $(PREFIX)/wu
