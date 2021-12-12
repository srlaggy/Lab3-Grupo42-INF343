###############
### WINDOWS ###
###############

build:
	@echo Construyendo fulcrum
	@del utils\RegistroPlanetario\"*.txt"
	@del utils\RecordsLog\"*.txt"
	@go build -o bin/fulcrum.exe fulcrum.go
	@echo Terminado

clean:
	@echo Eliminando archivos ejecutables
	@del utils\RegistroPlanetario\"*.txt"
	@del utils\RecordsLog\"*.txt"
	@del bin\"*.exe"
	@echo Terminado

run:
	@echo Iniciando fulcrum
	@del utils\RegistroPlanetario\"*.txt"
	@del utils\RecordsLog\"*.txt"
	@bin/fulcrum.exe

all:
	@echo Construyendo fulcrum
	@del utils\RegistroPlanetario\"*.txt"
	@del utils\RecordsLog\"*.txt"
	@go build -o bin/fulcrum.exe fulcrum.go
	@echo Iniciando fulcrum
	@bin/fulcrum.exe