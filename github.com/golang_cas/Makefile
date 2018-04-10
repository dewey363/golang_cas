run: deps
	GOPATH=$(CURDIR)/.go go run examples/examples.go 

deps: 
	cd examples; GOPATH=$(CURDIR)/.go go get -d
	rm -Rf $(CURDIR)/.go/src/github.com/golang_cas
	ln -s $(CURDIR) $(CURDIR)/.go/src/github.com/golang_cas
