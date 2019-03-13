GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)
LOCAL_BIN = $(shell pwd)/bin

default: build

bin/gox:
	GOBIN=$(LOCAL_BIN) go get github.com/mitchellh/gox
	
bin/redakt_linux_amd64 bin/redakt_darwin_amd64 bin/redakt_windows_amd64.exe: bin/gox $(GOFILES)
	$(LOCAL_BIN)/gox -osarch "linux/amd64 darwin/amd64 windows/amd64" -output "bin/redakt_{{.OS}}_{{.Arch}}" github.com/massiveco/redakt

build: bin/redakt_linux_amd64 bin/redakt_darwin_amd64 bin/redakt_windows_amd64.exe
