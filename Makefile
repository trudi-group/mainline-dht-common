default: build

build:
	go build

test:
	go build
	go test -v

debug:
	go build -gcflags="all=-N -l" 

