###############
### WINDOWS ###
###############

build:
	@echo Construyendo almirante
	@go build -o bin/almirante.exe almirante.go
	@echo Terminado

clean:
	@echo Eliminando archivos ejecutables
	@del bin\"*.exe"
	@echo Terminado

run:
	@echo Iniciando almirante
	@bin/almirante.exe

all:
	@echo Construyendo almirante
	@go build -o bin/almirante.exe almirante.go
	@echo Iniciando almirante
	@bin/almirante.exe