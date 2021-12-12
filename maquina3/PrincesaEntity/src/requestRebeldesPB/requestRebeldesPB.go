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

// --------------- FUNCIONES GRPC --------------- //
func GetNumberRebelds(planet string, city string, reloj ut.Reloj, server int, posicion int) {
	// Set up a connection to the server.
	conn1, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn1.Close()

	c := bp.NewBrokerPrincesaServiceClient(conn1)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.RequestRebeldes(ctx, &bp.RebeldesReq{Planeta: planet, Ciudad: city, Reloj: &bp.Reloj{Server1: reloj.Server1, Server2: reloj.Server2, Server3: reloj.Server3}, Server: int64(server)})
	ut.FailOnError(err, "Failed to send a entry")

	if (r.GetCantRebeldes() == -1){
		fmt.Println("El planeta no existe")
	} else if (r.GetCantRebeldes() == -2){
		fmt.Println("La ciudad no existe")
	}

	registros[posicion].RelojPlanet = *ut.CreateReloj(r.GetReloj().GetServer1(), r.GetReloj().GetServer2(), r.GetReloj().GetServer3())
	registros[posicion].Server = r.GetServer()

	fmt.Println("La cantidad de rebeldes es ", r.GetCantRebeldes())
}

// --------------- EXECUTE --------------- //
func Execute(){
	registros = make([]ut.RegisterPrincesa, 0)
	var eleccion int64
	var planeta string
	var ciudad string
	var pos int
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

			pos = ut.ContainsPlanet(registros, planeta, ciudad)
			if (pos == -1){
				pos = len(registros)
				registros = append(registros, ut.RegisterPrincesa{NamePlanet: planeta, NameCity: ciudad, RelojPlanet: ut.Reloj{Server1: 0, Server2: 0, Server3: 0}, Server: 1})
			}

			GetNumberRebelds(planeta, ciudad, registros[pos].RelojPlanet, int(registros[pos].Server), pos)
			time.Sleep(time.Second)
			ut.PrintRegister(registros)
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
	ut.PrintRegister(registros)
}

func GetRegistros() []ut.RegisterPrincesa{
	return registros
}