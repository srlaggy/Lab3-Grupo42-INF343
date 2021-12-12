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
)

func main() {
    spr.Comands("AddCity Mercurio Viña 10")
    spr.Comands("AddCity Venus Santiago 8")
    spr.Comands("AddCity Venus Rancagua")
    spr.Comands("AddCity Mercurio Valparaiso 8")
    spr.Comands("AddCity Mercurio Chiguayante")
    rb.Grpc_func()
}