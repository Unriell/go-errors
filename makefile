
SRC_PATH=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
VERSION=0.0.1

test:
	go test -v -tags=unit -coverprofile $(SRC_PATH)/errors/cover.out stash.bq.com/pdtdev/go-errors.git/errors
	go tool cover -html=$(SRC_PATH)/errors/cover.out -o $(SRC_PATH)/errors/cover.html

clean:
	rm -rf\
	 $(SRC_PATH)/dist\
	 $(SRC_PATH)/debug\
	 $(SRC_PATH)/*/cover.out\
	 $(SRC_PATH)/*/cover.html

all: sync test

docker_test:
	docker run -t --rm -v $(SRC_PATH):/go/src/stash.bq.com/pdtdev/go-errors.git/ golang:1.7 make -f /go/src/stash.bq.com/pdtdev/go-errors.git/makefile all 

version:
	@echo $(VERSION)-v$$(basename $(JOB_NAME))$(BUILD_NUMBER)	

sync:
	- go get -v github.com/kardianos/govendor
	- cd $(SRC_PATH) && govendor sync

