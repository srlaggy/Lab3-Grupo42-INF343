package sendRebeldesBP

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	bp "lab/broker/proto/BP"
	rr "lab/broker/src/requestRebelsFB"
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
	reloj, cantRebeldes, servidor := rr.RequestRebels(in.Planeta, in.Ciudad, *ut.CreateReloj(in.Reloj.Server1, in.Reloj.Server2, in.Reloj.Server3), int(in.Server))
	return &bp.RebeldesResp{Planeta: in.Planeta, Ciudad: in.Ciudad, Reloj: &bp.Reloj{Server1: reloj.Server1, Server2: reloj.Server2, Server3: reloj.Server3}, Server: int64(servidor), CantRebeldes: cantRebeldes}, nil
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