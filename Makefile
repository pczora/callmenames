GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARYNAME=callmenames

all: build
build:
	$(GOBUILD) -o $(BINARYNAME) -v
build-docker: build
	docker build . -t $(BINARYNAME)
clean:
	$(GOCLEAN)
