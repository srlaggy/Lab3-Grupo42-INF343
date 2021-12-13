package main

import (
	// ut "lab/broker/utils"
	sr "lab/broker/src/sendRebeldesBP"
	ss "lab/broker/src/sendServerBI"
	// "google.golang.org/grpc"
	// "fmt"
	// rr "lab/broker/src/requestRebelsFB"
)

// --------------- FUNCION MAIN --------------- //
func main(){
	// servidor que escucha a informante 1 para entregar la direccion de un servidor fulcrum
	go ss.Grpc_func()
	// servidor que escucha a informante 2 para entregar la direccion de un servidor fulcrum
	go ss.Grpc_func2()
	// servidor que escucha las peticiones de la princesa leia
	sr.Grpc_func()
	// rr.RequestRebels("Venus", "Santiago", *ut.CreateReloj(1,0,1))
}