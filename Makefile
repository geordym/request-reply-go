.PHONY: clean build

clean:	
	rm -f bootstrap
	rm -f build/bootstrap.zip

build-arm64:
	GOOS=linux GOARCH=arm64 go build -o bootstrap ./src/.

build:
	GOOS=linux GOARCH=amd64 go build -o main-x86 ./src/.

build-GoLambdaFunction:
	[ -d "$(ARTIFACTS_DIR)" ] || mkdir -p "$(ARTIFACTS_DIR)"
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main ./src/.
	cp targets.json "$(ARTIFACTS_DIR)"
	mv main "$(ARTIFACTS_DIR)"

zip:
	zip build/bootstrap.zip bootstrap targets.json

deploy: clean copyresources build-arm64 zip