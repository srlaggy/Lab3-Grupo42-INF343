package requestRebelsFB

import (
	"context"
	"time"
	"net"
	"log"
	"fmt"
	"math/rand"
	"google.golang.org/grpc"
	ut "lab/broker/utils"
	bf "lab/broker/proto/BF"
	ib "lab/broker/proto/IB"
)

const (
	protocolo_grpc = "tcp"
	port_grpc = "50501"
	port_grpc2 = "50502"
)

// FUNCIÓN: verifica si el reloj entregado por el Fulcrum cumple con el reads your write del informante
func VerificarReloj(relojInformante ut.Reloj, relojFulcrum ut.Reloj, server int) int{
	if (server == 1){
		if (relojFulcrum.Server1 == relojInformante.Server1){
			return 1
		}else if (relojFulcrum.Server1 > relojInformante.Server1){
			return 2
		} else {
			return 0
		}
	} else if (server == 2){
		if (relojFulcrum.Server2 == relojInformante.Server2){
			return 1
		}else if (relojFulcrum.Server2 > relojInformante.Server2){
			return 2
		} else {
			return 0
		}
	} else {
		if (relojFulcrum.Server3 == relojInformante.Server3){
			return 1
		}else if (relojFulcrum.Server3 > relojInformante.Server3){
			return 2
		} else {
			return 0
		}
	}
}

// --------------- FUNCIONES GRPC --------------- //
type server struct {
	ib.UnimplementedInformanteBrokerServiceServer
}

func (s *server) RequestServer(ctx context.Context, in *ib.ServerReq) (*ib.ServerResp, error) {
	fmt.Println("Recibi mensaje de informante")
	direccion, puerto, caso, server := RequestServerFulcrum(in.Planeta, in.Ciudad, *ut.CreateReloj(in.Reloj.Server1, in.Reloj.Server2, in.Reloj.Server3), int(in.Server))
	return &ib.ServerResp{Direccion: direccion, Puerto: puerto, Caso: caso, ServerNuevo: server}, nil
}

func Grpc_func() {
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc)
	ut.FailOnError(err, "Failed to listen")

	s := grpc.NewServer()
	ib.RegisterInformanteBrokerServiceServer(s, &server{})
	log.Printf("Servidor grpc escuchando en el puerto %v\n", port_grpc)
	ut.FailOnError(s.Serve(lis), "Failed to serve")
}

func Grpc_func2() {
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc2)
	ut.FailOnError(err, "Failed to listen")

	s := grpc.NewServer()
	ib.RegisterInformanteBrokerServiceServer(s, &server{})
	log.Printf("Servidor grpc escuchando en el puerto %v\n", port_grpc2)
	ut.FailOnError(s.Serve(lis), "Failed to serve")
}

// ----- FUNCIÓN: pedir direccion a fulcrum
// casoSave
// caso 1 -> registro no modificado por nadie mas (reloj igual)
// caso 2 -> registro modificado por un informante (reloj adelantado)
func RequestServerFulcrum(planeta string, ciudad string, reloj ut.Reloj, servidor int) (string, string, int64, int64){
	rand.Seed(time.Now().UnixNano())
	var addressAux string
	var portAux string
	var portAuxSend string
	var serverAux int64
	var addressSave string
	var portSave string
	var serverSave int64
	lista := rand.Perm(3)
	for j := 0; j < 3; j++ {
		if (lista[j] == 0){
			addressAux = "localhost"
			portAux = "40001"
			portAuxSend = "41001"
			serverAux = 1
		} else if (lista[j] == 1){
			addressAux = "localhost"
			portAux = "40002"
			portAuxSend = "41002"
			serverAux = 2
		} else{
			addressAux = "localhost"
			portAux = "40003"
			portAuxSend = "41003"
			serverAux = 3
		}
		// Set up a connection to the server.
		conn1, err := grpc.Dial(ut.CreateDir("", addressAux, portAux), grpc.WithInsecure(), grpc.WithBlock())
		ut.FailOnError(err, "Failed to create a connection")
		defer conn1.Close()

		c := bf.NewBrokerFulcrumServiceClient(conn1)
		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.RequestServer(ctx, &bf.ServerReq{Planeta: planeta})
		ut.FailOnError(err, "Failed to send server")

		switch VerificarReloj(reloj, *ut.CreateReloj(r.GetReloj().GetServer1(), r.GetReloj().GetServer2(), r.GetReloj().GetServer3()), servidor){
		case 1:
			return addressAux, portAuxSend, 1, serverAux
		case 2:
			addressSave = addressAux
			portSave = portAuxSend
			serverSave = serverAux
		}
	}
	return addressSave, portSave, 2, serverSave
}

