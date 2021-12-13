###############
### WINDOWS ###
###############

build:
	@echo Construyendo ahsoka
	@go build -o bin/ahsoka.exe ahsoka.go
	@echo Terminado

clean:
	@echo Eliminando archivos ejecutables
	@del bin\"*.exe"
	@echo Terminado

run:
	@echo Iniciando ahsoka
	@bin/ahsoka.exe

all:
	@echo Construyendo ahsoka
	@go build -o bin/ahsoka.exe ahsoka.go
	@echo Iniciando ahsoka
	@bin/ahsoka.exe