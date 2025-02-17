# lambda.mk

# Variable para la ubicación del código de la Lambda
LAMBDA_SRC = .

# El archivo compilado que Lambda usará
LAMBDA_BIN = main

# Variables para el entorno de Go
GOOS=linux
GOARCH=amd64

# Regla para limpiar archivos generados
clean:
	rm -f $(LAMBDA_BIN)

# Regla para construir el binario de la Lambda
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build -o $(LAMBDA_BIN) $(LAMBDA_SRC)

# Regla para empaquetar el binario en un archivo zip
package: build
	zip $(LAMBDA_BIN).zip $(LAMBDA_BIN)

# Regla para desplegar la Lambda utilizando SAM
deploy: package
	sam deploy --guided

# Regla para invocar la Lambda localmente (usando SAM)
invoke:
	sam local invoke GoLambdaFunction --event event.json
