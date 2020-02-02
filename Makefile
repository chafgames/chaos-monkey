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


BUILD_DATE := $(shell date +%Y%m%d)
VERSION := $(main.Version)
AUTHOR := $(main.Author)
RELEASE := DEV

LDFLAGS=-ldflags "-X=main.BuiltDate=${BUILD_DATE} -X=main.Version=${VERSION} -X=main.Author=${AUTHOR} -X=main.Release=${RELEASE}"

all: clean deps $(BINARY_NAME)
$(BINARY_NAME): 
	$(GOBUILD) -o $(BINARY_NAME)  $(LDFLAGS) -v 
# bin-windows:
#  	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME_WIN) $(LDFLAGS) -v
# bin-linux:
# 	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME_LINUX) $(LDFLAGS) -v

clean: 
	rm -rf vendor
	rm -f $(BINARY_NAME)
	# rm -f $(BINARY_NAME_LINUX)
deps:
	# go get github.com/go-gl/glfw/v3.2/glfw
	go mod vendor

.PHONY: clean deps all
# Cross compilation
