
include $(GOROOT)/src/Make.inc
PROGS = conditions forecast alerts slookup astronomy almanac
PREFIX = /usr/local/bin

all: $(PROGS)

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
	gofmt -tabwidth=2 -w=true conditions.go
	gofmt -tabwidth=2 -w=true forecast.go
	gofmt -tabwidth=2 -w=true alerts.go
	gofmt -tabwidth=2 -w=true slookup.go
	gofmt -tabwidth=2 -w=true astronomy.go
	gofmt -tabwidth=2 -w=true almanac.go

clean:
	rm -f $(PROGS)
	rm -f *.8
	rm -f *.6

install:
	cp conditions $(PREFIX)
	cp forecast $(PREFIX)
	cp alerts $(PREFIX)
	cp slookup $(PREFIX)
	cp astronomy $(PREFIX)
	cp almanac $(PREFIX)

uninstall:
	rm $(PREFIX)/conditions
	rm $(PREFIX)/forecast
	rm $(PREFIX)/alerts
	rm $(PREFIX)/slookup
	rm $(PREFIX)/astronomy
	rm $(PREFIX)/almanac
