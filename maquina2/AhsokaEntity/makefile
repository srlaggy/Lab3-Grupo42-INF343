###############
### WINDOWS ###
###############

# build:
# @echo Construyendo ahsoka
# @go build -o bin/ahsoka.exe ahsoka.go
# @echo Terminado

# clean:
# @echo Eliminando archivos ejecutables
# @del bin\"*.exe"
# @echo Terminado

# run:
# @echo Iniciando ahsoka
# @bin/ahsoka.exe

# all:
# @echo Construyendo ahsoka
# @go build -o bin/ahsoka.exe ahsoka.go
# @echo Iniciando ahsoka
# @bin/ahsoka.exe

#############
### LINUX ###
#############

build:
	@echo Construyendo ahsoka
	@go build -o bin/ahsoka ahsoka.go
	@echo Terminado

clean:
	@echo Eliminando archivos ejecutables
	@rm -f bin/*
	@echo Terminado

run:
	@echo Iniciando ahsoka
	@./bin/ahsoka

all:
	@echo Construyendo ahsoka
	@go build -o bin/ahsoka ahsoka.go
	@echo Iniciando ahsoka
	@./bin/ahsoka