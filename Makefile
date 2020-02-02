include build-ldflags.properties
export $(shell sed 's/=.*//' build-ldflags.properties)

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOCOVCMD=gocov
BINARY_NAME=chaos-monkey
BINARY_NAME_LINUX=chaos-monkey-linux
BINARY_NAME_LINUX=chaos-monkey-windows
RELEASE_FOLDER=chaosmonkey
ASSET_FOLDER=assets
RELEASE_BALL="$(RELEASE_FOLDER)-$(VERSION)-$(BUILD_DATE)-darwin-x86_64.tgz"



BUILD_DATE := $(shell date +%Y%m%d)
VERSION := $(main.Version)
AUTHOR := $(main.Author)
RELEASE := DEV

LDFLAGS=-ldflags "-X=main.BuiltDate=${BUILD_DATE} -X=main.Version=${VERSION} -X=main.Author=${AUTHOR} -X=main.Release=${RELEASE}"

all: clean deps $(BINARY_NAME) 
release: clean deps $(BINARY_NAME)  darwin-tarball
$(BINARY_NAME): 
	$(GOBUILD) -o $(BINARY_NAME)  $(LDFLAGS) -v 
darwin-tarball: 
	mkdir $(RELEASE_FOLDER)
	cp $(BINARY_NAME) $(RELEASE_FOLDER)
	cp -r $(ASSET_FOLDER) $(RELEASE_FOLDER)
	tar czvf $(RELEASE_BALL) $(RELEASE_FOLDER)
# bin-windows:
#  	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME_WIN) $(LDFLAGS) -v
# bin-linux:
# 	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME_LINUX) $(LDFLAGS) -v

clean: 
	rm -rf vendor
	rm -f $(BINARY_NAME)
	rm -rf $(RELEASE_FOLDER)
	rm -f $(RELEASE_BALL)
	# rm -f $(BINARY_NAME_LINUX)
deps:
	# go get github.com/go-gl/glfw/v3.2/glfw
	go mod vendor

.PHONY: clean deps all
# Cross compilation
