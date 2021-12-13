package sendServerFB

import (
	"context"
	"log"
	"net"
	"fmt"
	"google.golang.org/grpc"
	bf "lab/fulcrum/proto/BF"
	spr "lab/fulcrum/src/storePlanetaryRecords"
	ut "lab/fulcrum/utils"
)

const (
	protocolo_grpc = "tcp"
	port_grpc = "40001"
)

// VARIABLE GLOBAL

// --------------- FUNCIONES GRPC --------------- // -> FULCRUM actua como servidor
type server struct {
	bf.UnimplementedBrokerFulcrumServiceServer
}

func (s *server) RequestServer(ctx context.Context, in *bf.ServerReq) (*bf.ServerResp, error) {
	fmt.Println("Recibi mensaje de broker desde informante")
	for i := 0; i < len(spr.GetListRegister()); i++ {
		if (spr.GetListRegister()[i].NamePlanet == in.Planeta){
			return &bf.ServerResp{Reloj: &bf.Reloj{Server1: spr.GetListRegister()[i].RelojPlanet.Server1, Server2: spr.GetListRegister()[i].RelojPlanet.Server2, Server3: spr.GetListRegister()[i].RelojPlanet.Server3}, Existe: 1}, nil
		}
	}
	return &bf.ServerResp{Reloj: &bf.Reloj{Server1: 0, Server2: 0, Server3: 0}, Existe: 0}, nil
}

// --------------- FUNCION PRINCIPAL --------------- //
func Grpc_func() {
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc)
	ut.FailOnError(err, "Failed to listen")

	s := grpc.NewServer()
	bf.RegisterBrokerFulcrumServiceServer(s, &server{})
	log.Printf("Servidor grpc escuchando en el puerto %v\n", port_grpc)
	ut.FailOnError(s.Serve(lis), "Failed to serve")
}