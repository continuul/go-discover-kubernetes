CGO_ENABLED=0

help:
	@echo "build   - build go-discover"
	@echo "test    - test go-discover"
	@echo "clean   - remove temp files"

build: compile
	docker build -t continuul.io/discover:latest .

compile:
	@echo "*** Compiling continuul.io/go-discover/cmd/discover"
	GOOS=linux GOARCH=amd64 go build -i -ldflags '-s -w' continuul.io/go-discover/cmd/discover

test:
	@echo "*** Running go-discover test"

clean:
	rm -f discover
	rm -rf .terraform
	rm -f terraform.tfstate{,.backup}
	rm -f tf_rsa{,.pub}
	rm -f tf.log

.PHONY: compile build test clean
