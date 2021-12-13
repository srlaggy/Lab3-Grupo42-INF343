package main

// Fulcrum

import (
	rb "lab/fulcrum/src/requestRebelsFB"
	ss "lab/fulcrum/src/sendServerFB"
    rc "lab/fulcrum/src/requestCommandIF"
)

func main() {
    // Conexión Grpc abierta para escuchar llamados del Broker desde la princesa
    go rb.Grpc_func()
    // Servidor grpc para escuchar las llamadas directas del informante 1
    go rc.Grpc_func()
    // Servidor grpc para escuchar las llamadas directas del informante 2
    go rc.Grpc_func2()
    // Servidor grpc para llamadas de broker desde informante
    ss.Grpc_func()
}