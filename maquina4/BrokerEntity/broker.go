package main

import (
	// ut "lab/broker/utils"
	sr "lab/broker/src/sendRebeldesBP"
	// "google.golang.org/grpc"
	// "fmt"
	// rr "lab/broker/src/requestRebelsFB"
)

// --------------- FUNCION MAIN --------------- //
func main(){
	sr.Grpc_func()
	// rr.RequestRebels("Venus", "Santiago", *ut.CreateReloj(1,0,1))
}