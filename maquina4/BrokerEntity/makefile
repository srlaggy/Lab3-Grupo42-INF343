###############
### WINDOWS ###
###############

build:
	@echo Construyendo broker
	@go build -o bin/broker.exe broker.go
	@echo Terminado

clean:
	@echo Eliminando archivos ejecutables
	@del bin\"*.exe"
	@echo Terminado

run:
	@echo Iniciando broker
	@bin/broker.exe

all:
	@echo Construyendo broker
	@go build -o bin/broker.exe broker.go
	@echo Iniciando broker
	@bin/broker.exe