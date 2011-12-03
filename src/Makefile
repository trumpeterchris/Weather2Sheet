include $(GOROOT)/src/Make.inc
PROGS = conditions forecast alerts slookup astronomy almanac wu
PREFIX = $(GOBIN)

all: $(PROGS)

wu:	wu.go
	$(GC) -o wu.$(O) wu.go
	$(LD) -o wu wu.$(O)

conditions:	conditions.go utils.$(O)
	$(GC) -o conditions.$(O) conditions.go
	$(LD) -o conditions conditions.$(O)

almanac:	almanac.go utils.$(O)
	$(GC) -o almanac.$(O) almanac.go
	$(LD) -o almanac almanac.$(O)

astronomy:	astronomy.go utils.$(O)
	$(GC) -o astronomy.$(O) astronomy.go
	$(LD) -o astronomy astronomy.$(O)

slookup:	slookup.go utils.$(O)
	$(GC) -o slookup.$(O) slookup.go
	$(LD) -o slookup slookup.$(O)

alerts:	alerts.go utils.$(O)
	$(GC) -o alerts.$(O) alerts.go
	$(LD) -o alerts alerts.$(O)

forecast:	forecast.go utils.$(O)
	$(GC) -o forecast.$(O) forecast.go
	$(LD) -o forecast forecast.$(O)

utils.$(O):	utils.go
	$(GC) utils.go 

format:
	gofmt -tabwidth=2 -w wu.go
	gofmt -tabwidth=2 -w conditions.go
	gofmt -tabwidth=2 -w forecast.go
	gofmt -tabwidth=2 -w alerts.go
	gofmt -tabwidth=2 -w slookup.go
	gofmt -tabwidth=2 -w astronomy.go
	gofmt -tabwidth=2 -w almanac.go

clean:
	rm -f $(PROGS)
	rm -f conditions.$(O) forecast.$(O) alerts.$(O) slookup.$(O) astronomy.$(O) almanac.$(O) wu.$(O)

install:
	cp conditions forecast alerts slookup astronomy almanac wu $(PREFIX)

uninstall:
	rm $(PREFIX)/conditions
	rm $(PREFIX)/forecast
	rm $(PREFIX)/alerts
	rm $(PREFIX)/slookup
	rm $(PREFIX)/astronomy
	rm $(PREFIX)/almanac
	rm $(PREFIX)/wu
