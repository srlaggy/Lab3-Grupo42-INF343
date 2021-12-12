package requestRebelsFB

// NombreEntidad

import(
    "context"
    "log"
    "net"
    // "time"
    // "fmt"
	"os"
	"bufio"
	"strings"
    "strconv"
	"errors"
    "google.golang.org/grpc"
	ut "lab/fulcrum/utils"
	sp "lab/fulcrum/src/storePlanetaryRecords"
	sm "lab/fulcrum/proto/requestRebelsFB"
)

type Reloj struct {
    Server1 int64
    Server2 int64
    Server3 int64
}

type Register struct {
    NamePlanet string
    RelojPlanet Reloj
}

const (
	protocolo_grpc = "tcp"
	port_grpc = "60003"
)

// ----------------- FUNCIONES ------------------ //
// Funcion que retorna la cantidad de rebeldes
func GetRebels(planeta string, ciudad string) (int64, ut.Reloj){
	//Se abre el archivo
	var fileName string = "utils/RegistroPlanetario/" + planeta + ".txt"
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		return -1, *ut.CreateReloj(0,0,0)	// -1 : no existe el planeta
	}
	file, error1 := os.Open(fileName)
	ut.FailOnError(error1, "Failed to open file")
	scanner := bufio.NewScanner(file)
	//Recorrer el archivo para obtener la ciudad
	defer file.Close()
	for scanner.Scan(){
		registerCity := scanner.Text()
		s := strings.Split(registerCity, " ")
		if (s[0] == ciudad){
			// se iterará la lista de registros para obtener el reloj del planeta del presente servidor (1)
			for i := 0; i < len(sp.GetListRegister()); i++ {
				if (sp.GetListRegister()[i].NamePlanet == planeta){
					reloj_actual := sp.GetListRegister()[i].RelojPlanet     // MODIFICAR Server1 por el FULCRUM ACTUAL
					valor, _ := strconv.ParseInt(s[1], 10, 64)            // transformar a int64 la cantidad de rebeldes
					return valor, reloj_actual                                // retorno la cantidad de rebeldes, reloj del planeta
				}
			}
		}
	}
	return -2, *ut.CreateReloj(0,0,0)	// -2 : no existe la ciudad
}

// --------------- FUNCIONES GRPC --------------- //

// ----- FUNCIÓN: pedir cantidad de rebeldes a un Fulcrum ----- // --> Fulcrum actua como servidor
type server struct {
	sm.UnimplementedRequestRebelsServiceServer
}

func (s *server) RequestRebelsFB(ctx context.Context, in *sm.RebeldesReq) (*sm.RebeldesResp, error) {
	log.Printf("Received")
	cantReb, reloj := GetRebels(in.Planeta, in.Ciudad)
	return &sm.RebeldesResp{CantRebeldes: cantReb, Reloj: &sm.Reloj{Server1: reloj.Server1, Server2: reloj.Server2, Server3: reloj.Server3}}, nil	// envia la cantidad de rebeldes y el reloj del planeta
}

// --------------- FUNCION PRINCIPAL --------------- //
func Grpc_func() {
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc)
	ut.FailOnError(err, "Failed to listen")

	s := grpc.NewServer()
	sm.RegisterRequestRebelsServiceServer(s, &server{})
	log.Printf("Servidor grpc escuchando en el puerto %v", port_grpc)
	ut.FailOnError(s.Serve(lis), "Failed to serve")
}