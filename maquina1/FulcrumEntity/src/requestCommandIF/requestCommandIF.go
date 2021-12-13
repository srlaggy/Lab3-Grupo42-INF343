package requestCommandIF

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	ife "lab/fulcrum/proto/IF"
	spr "lab/fulcrum/src/storePlanetaryRecords"
	ut "lab/fulcrum/utils"
)

const (
	protocolo_grpc = "tcp"
	port_grpc = "41001"      // CAMBIAR SEGUN SERVIDOR FULCRUM
	port_grpc2 = "42001"
)

// --------------- FUNCIONES GRPC --------------- // -> FULCRUM actua como servidor
type server struct {
	ife.UnimplementedInformanteFulcrumServiceServer
}

// state puede ser:
// 1 -> si el comando se ejecuto bien
// 2 -> si el comando no pudo ser ejecutado
func (s *server) RequestWrite(ctx context.Context, in *ife.WriteReq) (*ife.WriteResp, error) {
	var strCompleto string = in.Comando + " " + in.Planeta + " " + in.Ciudad
	if (in.Valor != ""){
		strCompleto += " " + in.Valor
	}
	clock, state := spr.Commands(strCompleto)
	return &ife.WriteResp{Reloj: &ife.Reloj{Server1: clock.Server1, Server2: clock.Server2, Server3: clock.Server3}, Estado: state}, nil
}

// --------------- FUNCION PRINCIPAL --------------- //
func Grpc_func() {
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc)
	ut.FailOnError(err, "Failed to listen")

	s := grpc.NewServer()
	ife.RegisterInformanteFulcrumServiceServer(s, &server{})
	log.Printf("Servidor grpc escuchando en el puerto %v\n", port_grpc)
	ut.FailOnError(s.Serve(lis), "Failed to serve")
}

func Grpc_func2() {
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc2)
	ut.FailOnError(err, "Failed to listen")

	s := grpc.NewServer()
	ife.RegisterInformanteFulcrumServiceServer(s, &server{})
	log.Printf("Servidor grpc escuchando en el puerto %v\n", port_grpc2)
	ut.FailOnError(s.Serve(lis), "Failed to serve")
}