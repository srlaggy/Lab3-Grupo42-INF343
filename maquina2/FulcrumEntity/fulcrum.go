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
    spr.Comands("AddCity Tierra Quellon 15")
    spr.Comands("AddCity Jupiter Castro 3")
    spr.Comands("AddCity Tierra Ancud")
    rb.Grpc_func()
}