package requestRebelsFB

import (
	"context"
	"time"
	// "fmt"
	"math/rand"
	"google.golang.org/grpc"
	ut "lab/broker/utils"
	sm "lab/broker/proto/requestRebelsFB"
)

const (
	address = "localhost"
	protocolo_grpc = ""
)

// FUNCIÓN: verifica si el reloj entregado por el Fulcrum cumple con el Monotonic Reads de la princesa Leia
func VerificarReloj(relojPrincesa ut.Reloj, relojFulcrum ut.Reloj, server int) bool{
	if (server == 1){
		if (relojFulcrum.Server1 >= relojPrincesa.Server1){
			return true
		}else{
			return false
		}
	} else if (server == 2){
		if (relojFulcrum.Server2 >= relojPrincesa.Server2){
			return true
		} else{
			return false
		}
	} else{
		if (relojFulcrum.Server3 >= relojPrincesa.Server3){
			return true
		} else{
			return false
		}
	}
}

// --------------- FUNCIONES GRPC --------------- //

// ----- FUNCIÓN: pedir cantidad de rebeldes a un Fulcrum ----- // --> Broker actua como cliente

func RequestRebels(planeta string, ciudad string, reloj ut.Reloj, servidor int) (ut.Reloj, int64, int){
	rand.Seed(time.Now().UnixNano())
	var port_grpc string
	var server int
	lista := rand.Perm(3)
	for j := 0; j < 3; j++ {
		if (lista[j] == 0){
			port_grpc = "60001"
			server = 1
		} else if (lista[j] == 1){
			port_grpc = "60002"
			server = 2
		} else{
			port_grpc = "60003"
			server = 3
		}
		// Set up a connection to the server.
		conn1, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
		ut.FailOnError(err, "Failed to create a connection")
		defer conn1.Close()

		c := sm.NewRequestRebelsServiceClient(conn1)
		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.RequestRebelsFB(ctx, &sm.RebeldesReq{Planeta: planeta, Ciudad: ciudad})
		ut.FailOnError(err, "Failed to send Rebels")

		if ((r.GetCantRebeldes() < 0) && (j != 2)){
			continue
		} else if (!VerificarReloj(reloj, *ut.CreateReloj(r.GetReloj().GetServer1(), r.GetReloj().GetServer2(), r.GetReloj().GetServer3()), servidor)){
			continue
		} else {
			return *ut.CreateReloj(r.GetReloj().GetServer1(), r.GetReloj().GetServer2(), r.GetReloj().GetServer3()), r.GetCantRebeldes(), server
			// relojAux : reloj de la respuesta (Fulcrum)
			// relojAux := *ut.CreateReloj(r.GetReloj().GetServer1(), r.GetReloj().GetServer2(), r.GetReloj().GetServer3())
			// flag = VerificarReloj(reloj, relojAux, server)
			// impresión: Princesa, Fulcrum, flag, cantidad de rebeldes
			// fmt.Println("server: ", server, " -> ", reloj, relojAux, flag, r.GetCantRebeldes())
		}
	}
	return *ut.CreateReloj(0,0,0), 0, 1
}