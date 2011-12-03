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
<<<<<<< HEAD
	rm -f wu.$(O)
=======
	rm -f conditions.$(O) forecast.$(O) alerts.$(O) slookup.$(O) astronomy.$(O) almanac.$(O) utils.$(O) wu.$(O)
>>>>>>> 40238e9f0daa296b5234b55113bf046f8a54caaa

install:
	cp wu $(PREFIX)

uninstall:
	rm $(PREFIX)/wu
