###############
### WINDOWS ###
###############

build:
	@echo Construyendo fulcrum
	@go build -o bin/fulcrum.exe fulcrum.go
	@echo Terminado

clean:
	@echo Eliminando archivos ejecutables
	@del bin\"*.exe"
	@echo Terminado

run:
	@echo Iniciando fulcrum
	@bin/fulcrum.exe

all:
	@echo Construyendo fulcrum
	@go build -o bin/fulcrum.exe fulcrum.go
	@echo Iniciando fulcrum
	@bin/fulcrum.exe