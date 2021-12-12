package main

// NombreEntidad

import(
    // "context"
    // "log"
    // "net"
    // "time"
    // "fmt"
    // "google.golang.org/grpc"
	spr "lab/fulcrum/src/storePlanetaryRecords"
	rb "lab/fulcrum/src/requestRebelsFB"
	// "time"
)

func main() {
    // Pruebas
    spr.Commands("AddCity Mercurio Viña 10")
    spr.Commands("AddCity Venus Santiago 8")
    spr.Commands("AddCity Venus Rancagua")
    spr.Commands("AddCity Mercurio Valparaiso 8")
    spr.Commands("AddCity Mercurio Chiguayante")
    spr.Commands("UpdateName Mercurio Chiguayante Coquimbo")
    spr.Commands("UpdateNumber Mercurio Valparaiso 5")
    spr.Commands("DeleteCity Venus Santiago")

    // Conexión Grpc abierta para escuchar llamados del Broker
    rb.Grpc_func()

	// time.Sleep(3*time.Second)

    // spr.CleanRecordsLog()
}