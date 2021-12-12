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
    spr.Comands("AddCity Mercurio Viña 2")
    spr.Comands("AddCity Tierra Quellon 3")
    rb.Grpc_func()
}