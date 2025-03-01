.PHONY: clean build

clean:	
	rm -f main

build-arm64:
	GOOS=linux GOARCH=arm64 go build -o main ./src/.

build:
	GOOS=linux GOARCH=amd64 go build -o main-x86 ./src/.

build-GoLambdaFunction:
	# Verificar si el directorio existe, y si no, crearlo
	[ -d "$(ARTIFACTS_DIR)" ] || mkdir -p "$(ARTIFACTS_DIR)"
	# Compilar y mover el archivo
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main ./src/.
	mv main "$(ARTIFACTS_DIR)"