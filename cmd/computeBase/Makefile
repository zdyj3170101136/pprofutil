# vendor dependencies:
.PHONY: internal/pprof
internal/pprof:
	rm -rf internal/pprof
	rm -rf tmp
	git clone https://github.com/google/pprof tmp/pprof
	mkdir -p internal
	cp -r tmp/pprof/internal internal/pprof
    # in mac brew install brew install gnu-sed, and replace sed with gsed
	find internal/pprof -type f -exec gsed -i 's/github.com\/google\/pprof\/internal/github.com\/zdyj3170101136\/pprofutil\/cmd\/computeBase\/internal\/pprof/g' {} +
	rm -rf tmp
