package requestRebeldesPB

import (
	"context"
	"time"
	"fmt"
	"google.golang.org/grpc"
	ut "lab/princesa/utils"
	bp "lab/princesa/proto/BP"
)

const (
	address = "localhost"
	protocolo_grpc = ""
	port_grpc = "50000"
)

// variables globales
var registros []ut.RegisterPrincesa
var vivosSlice []bool

// --------------- FUNCIONES GRPC --------------- //
func GetNumberRebelds(entrada string) {
	// Set up a connection to the server.
	conn1, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn1.Close()

	c := lj.NewLiderJugadorServiceClient(conn1)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	nroJugadorBot = make([]int64, maxJug)
	vivosSlice = make([]bool, maxJug)
	for i:=0; i<maxJug; i++{
		vivosSlice[i] = true
	}

	r, err := c.RequestGame(ctx, &lj.GameReq{EntryMsg: entrada})
	ut.FailOnError(err, "Failed to send a entry")
	// almacenamos numero del jugador
	nroJugador = r.GetNroJugador()
	nroJugadorBot[nroJugador-1] = nroJugador
	fmt.Printf("%s", r.GetGameMsg())
}

// --------------- EXECUTE --------------- //
func Execute(){
	registros = make([]ut.RegisterPrincesa, 0)
	var eleccion int64
	var planeta string
	var ciudad string
	var continuar bool = true
	fmt.Println("\nBienvenida Princesa Leia")
	fmt.Println("Sabemos que quieres estar informada de todo lo que acontece en los planetas vecinos")
	fmt.Println("Hemos creado un menu para facilitarte esta informacion")

	for continuar{
		fmt.Println("\n|-----------Menu------------|")
		fmt.Println("1) Consultar por los rebeldes")
		fmt.Println("2) Salir del menu\n")
		fmt.Print("Ingrese una opcion: ")
		fmt.Scanln(&eleccion)

		switch eleccion{
		case 1:
			time.Sleep(time.Second)
			fmt.Println("\n|-----------Rebeldes------------|")
			fmt.Print("Ingrese nombre del planeta: ")
			fmt.Scanln(&planeta)
			fmt.Print("Ingrese nombre de la ciudad: ")
			fmt.Scanln(&ciudad)
			time.Sleep(time.Second)
		case 2:
			time.Sleep(time.Second)
			continuar = false
			fmt.Println("Adios Princesa Leia\n")
			time.Sleep(time.Second)
		default:
			time.Sleep(time.Second)
			fmt.Println("Ingrese una opcion valida\n")
			time.Sleep(time.Second)
		}
	}
}

func GetRegistros() []RegisterPrincesa{
	return registros
}

func GetVivosSlice() []bool{
	return vivosSlice
}