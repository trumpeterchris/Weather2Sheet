include $(GOROOT)/src/Make.inc
PROGS =  wu
PREFIX = $(GOBIN)

all: $(PROGS)

wu: alerts almanac astro conditions forecast forecast7 lookup history 
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

forecast7:
	$(GC) -o forecast7.$(O) forecast7.go

lookup:
	$(GC) -o lookup.$(O) lookup.go

history:
	$(GC) -o history.$(O) history.go


format:
	  gofmt -tabwidth=2 -w wu.go
	  gofmt -tabwidth=2 -w alerts.go
	  gofmt -tabwidth=2 -w almanac.go
	  gofmt -tabwidth=2 -w astro.go
	  gofmt -tabwidth=2 -w conditions.go
	  gofmt -tabwidth=2 -w forecast.go
	  gofmt -tabwidth=2 -w forecast7.go
	  gofmt -tabwidth=2 -w lookup.go
	  gofmt -tabwidth=2 -w history.go

clean:
	rm -f $(PROGS)
	rm -f wu.$(O)
	rm -f alerts.$(O)
	rm -f almanac.$(O)
	rm -f astro.$(O)
	rm -f conditions.$(O)
	rm -f forecast.$(O)
	rm -f forecast7.$(O)
	rm -f lookup.$(O)
	rm -f history.$(O)

install:
	cp wu $(PREFIX)

uninstall:
	rm $(PREFIX)/wu
