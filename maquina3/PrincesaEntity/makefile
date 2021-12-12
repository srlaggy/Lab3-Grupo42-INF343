###############
### WINDOWS ###
###############

build:
	@echo Construyendo princesa
	@go build -o bin/princesa.exe princesa.go
	@echo Terminado

clean:
	@echo Eliminando archivos ejecutables
	@del bin\"*.exe"
	@echo Terminado

run:
	@echo Iniciando princesa
	@bin/princesa.exe

all:
	@echo Construyendo princesa
	@go build -o bin/princesa.exe princesa.go
	@echo Iniciando princesa
	@bin/princesa.exe