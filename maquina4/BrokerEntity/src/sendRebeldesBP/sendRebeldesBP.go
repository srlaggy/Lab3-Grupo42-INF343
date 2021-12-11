package sendRebeldesBP

import (
	"context"
	"log"
	"net"
	"fmt"
	"google.golang.org/grpc"
	bp "lab/broker/proto/BP"
	ut "lab/broker/utils"
)

const (
	protocolo_grpc = "tcp"
	port_grpc = "50000"
)

// VARIABLE GLOBAL
var test string

// --------------- FUNCIONES GRPC --------------- //
type server struct {
	bp.UnimplementedBrokerPrincesaServiceServer
}

func (s *server) RequestRebeldes(ctx context.Context, in *bp.RebeldesReq) (*bp.RebeldesResp, error) {
	test = "\nEl juego ya comenzó. No puedes ingresar.\n"
	// if value!=0{
	// 	mensajeDeEntrada = fmt.Sprintf("\nEstas dentro del juego. Eres el jugador %d\n", value)
	// 	log.Printf("Entry Received")
	// }
	return &bp.RebeldesResp{GameMsg: mensajeDeEntrada, NroJugador: value}, nil
}

// --------------- FUNCION PRINCIPAL --------------- //
func Grpc_func() {
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc)
	ut.FailOnError(err, "Failed to listen")

	s := grpc.NewServer()
	bp.RegisterBrokerPrincesaServiceServer(s, &server{})
	log.Printf("Servidor grpc escuchando en el puerto %v\n", port_grpc)
	ut.FailOnError(s.Serve(lis), "Failed to serve")
}