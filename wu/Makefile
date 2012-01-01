include $(GOROOT)/src/Make.inc
PROGS =  wu
PREFIX = $(GOBIN)

all: $(PROGS)

wu: alerts almanac astro conditions forecast lookup yesterday
	$(GC) -o wu.$(O) wu.go
	$(LD) -o wu wu.$(O)

alerts:
	$(GC) -o alerts.$(O) alerts.go

almanac:
	$(GC) -o almanac.$(O) almanac.go

astro:
	$(GC) -o astro.$(O) astro.go

conditions:
	$(GC) -o conditions.$(O) conditions.go

forecast:
	$(GC) -o forecast.$(O) forecast.go

lookup:
	$(GC) -o lookup.$(O) lookup.go

yesterday:
	$(GC) -o yesterday.$(O) yesterday.go


format:
	  gofmt -tabwidth=2 -w wu.go
	  gofmt -tabwidth=2 -w alerts.go
	  gofmt -tabwidth=2 -w almanac.go
	  gofmt -tabwidth=2 -w astro.go
	  gofmt -tabwidth=2 -w conditions.go
	  gofmt -tabwidth=2 -w forecast.go
	  gofmt -tabwidth=2 -w lookup.go
	  gofmt -tabwidth=2 -w yesterday.go

clean:
	rm -f $(PROGS)
	rm -f wu.$(O)
	rm -f alerts.$(O)
	rm -f almanac.$(O)
	rm -f astro.$(O)
	rm -f conditions.$(O)
	rm -f forecast.$(O)
	rm -f lookup.$(O)
	rm -f yesterday.$(O)

install:
	cp wu $(PREFIX)

uninstall:
	rm $(PREFIX)/wu
