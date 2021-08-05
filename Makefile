default: build

build:
	go build

debug:
	go build -gcflags="all=-N -l" 

