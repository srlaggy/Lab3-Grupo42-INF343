package requestCommandIF

import (
	"context"
	"time"
	"fmt"
	"google.golang.org/grpc"
	ut "lab/ahsoka/utils"
	ib "lab/ahsoka/proto/IB"
	ife "lab/ahsoka/proto/IF"
)

const (
	address = "localhost"
	protocolo_grpc = ""
	port_grpc = "50502"
)

// variables globales
var registros []ut.RegisterInformante
var addressF string
var port_grpcF string
var caso int64
var serverNuevo int64

// --------------- FUNCIONES GRPC --------------- //
// entrega la direccion, puerto y el caso que se da
// caso 1 -> registro no modificado por nadie mas (reloj igual)
// caso 2 -> registro modificado por un externo (reloj adelantado)
func RequestServer(planet string, city string, reloj ut.Reloj, server int) (string, string, int64, int64) {
	// Set up a connection to the server.
	conn1, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn1.Close()

	c := ib.NewInformanteBrokerServiceClient(conn1)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.RequestServer(ctx, &ib.ServerReq{Planeta: planet, Ciudad: city, Reloj: &ib.Reloj{Server1: reloj.Server1, Server2: reloj.Server2, Server3: reloj.Server3}, Server: int64(server)})
	ut.FailOnError(err, "Failed to send a entry")

	return r.GetDireccion(), r.GetPuerto(), r.GetCaso(), r.GetServerNuevo()
}

// r retorna un estado que puede ser:
// 1 -> si el comando se ejecuto bien
// 2 -> si el comando no pudo ser ejecutado
func ChangeRegister(command string, planet string, city string, value string, caso int64, posicion int, addressNew string, portNew string, servidorNuevo int64) {
	// Set up a connection to the server.
	conn2, err := grpc.Dial(ut.CreateDir(protocolo_grpc, addressNew, portNew), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn2.Close()

	c := ife.NewInformanteFulcrumServiceClient(conn2)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.RequestWrite(ctx, &ife.WriteReq{Comando: command, Planeta: planet, Ciudad: city, Valor: value, Caso: caso})
	ut.FailOnError(err, "Failed to send a entry")

	switch command{
	case "AddCity":
		if (r.GetEstado() == 1){
			registros[posicion].Server = servidorNuevo
			registros[posicion].RelojPlanet = *ut.CreateReloj(r.GetReloj().GetServer1(), r.GetReloj().GetServer2(), r.GetReloj().GetServer3())
			fmt.Println("\nEl comando fue ejecutado con exito\n")
		} else {
			fmt.Println("\nEl comando no pudo ser ejecutado\n")
		}
	case "UpdateName":
		if (r.GetEstado() == 1){
			registros[posicion].NameCity = value
			registros[posicion].Server = servidorNuevo
			registros[posicion].RelojPlanet = *ut.CreateReloj(r.GetReloj().GetServer1(), r.GetReloj().GetServer2(), r.GetReloj().GetServer3())
			fmt.Println("\nEl comando fue ejecutado con exito\n")
		} else {
			fmt.Println("\nEl comando no pudo ser ejecutado\n")
		}
	case "UpdateNumber":
		if (r.GetEstado() == 1){
			registros[posicion].Server = servidorNuevo
			registros[posicion].RelojPlanet = *ut.CreateReloj(r.GetReloj().GetServer1(), r.GetReloj().GetServer2(), r.GetReloj().GetServer3())
			fmt.Println("\nEl comando fue ejecutado con exito\n")
		} else {
			fmt.Println("\nEl comando no pudo ser ejecutado\n")
		}
	case "DeleteCity":
		if (r.GetEstado() == 1){
			registros[posicion].Server = servidorNuevo
			registros[posicion].RelojPlanet = *ut.CreateReloj(r.GetReloj().GetServer1(), r.GetReloj().GetServer2(), r.GetReloj().GetServer3())
			fmt.Println("\nEl comando fue ejecutado con exito\n")
		} else {
			fmt.Println("\nEl comando no pudo ser ejecutado\n")
		}
	}
}

// --------------- EXECUTE --------------- //
func Execute(){
	registros = make([]ut.RegisterInformante, 0)
	var eleccion int64
	var comando string
	var planeta string
	var ciudad string
	var valor string
	var pos int
	var continuar bool = true
	fmt.Println("\nBienvenido querido Informante")
	fmt.Println("Sabemos que tu mision es informarnos sobre los acontecimientos de los planetas vecinos")
	fmt.Println("Por lo mismo, hemos creado un menu para facilitar tu mision y quitarte tanto papeleo")

	for continuar{
		fmt.Println("\n|-----------Menu------------|")
		fmt.Println("1) Agregar nueva ciudad a los registros")
		fmt.Println("2) Actualizar nombre de una ciudad en los registros")
		fmt.Println("3) Actualizar numero de rebeldes en una ciudad")
		fmt.Println("4) Eliminar ciudad en los registros")
		fmt.Println("5) Salir del menu\n")
		fmt.Print("Ingrese una opcion: ")
		fmt.Scanln(&eleccion)

		switch eleccion{
		case 1:
			comando = "AddCity"
			time.Sleep(time.Second)
			fmt.Println("\n|-----------AddCity------------|")
			fmt.Print("Ingrese nombre del planeta: ")
			fmt.Scanln(&planeta)
			fmt.Print("Ingrese nombre de la ciudad: ")
			fmt.Scanln(&ciudad)
			fmt.Print("Ingrese numero de rebeldes en la ciudad (puede omitirlo si escribe omitir): ")
			fmt.Scanln(&valor)
			if (valor == "omitir"){
				valor = ""
			}
			time.Sleep(time.Second)

			pos = ut.ContainsPlanet(registros, planeta, ciudad)
			if (pos == -1){
				pos = len(registros)
				registros = append(registros, ut.RegisterInformante{NamePlanet: planeta, NameCity: ciudad, RelojPlanet: ut.Reloj{Server1: 0, Server2: 0, Server3: 0}, Server: 1})
			}

			addressF, port_grpcF, caso, serverNuevo = RequestServer(planeta, ciudad, registros[pos].RelojPlanet, int(registros[pos].Server))
			ChangeRegister(comando, planeta, ciudad, valor, caso, pos, addressF, port_grpcF, serverNuevo)
			time.Sleep(time.Second)
			ut.PrintRegister(registros)
		case 2:
			comando = "UpdateName"
			time.Sleep(time.Second)
			fmt.Println("\n|-----------UpdateName------------|")
			fmt.Print("Ingrese nombre del planeta: ")
			fmt.Scanln(&planeta)
			fmt.Print("Ingrese nombre de la ciudad a actualizar: ")
			fmt.Scanln(&ciudad)
			fmt.Print("Ingrese nombre del nuevo nombre de la ciudad (debe ser distinto al actual): ")
			fmt.Scanln(&valor)
			time.Sleep(time.Second)

			pos = ut.ContainsPlanet(registros, planeta, ciudad)
			if (pos == -1){
				pos = len(registros)
				registros = append(registros, ut.RegisterInformante{NamePlanet: planeta, NameCity: ciudad, RelojPlanet: ut.Reloj{Server1: 0, Server2: 0, Server3: 0}, Server: 1})
			} else {
				registros[pos].NameCity = valor
			}

			addressF, port_grpcF, caso, serverNuevo = RequestServer(planeta, ciudad, registros[pos].RelojPlanet, int(registros[pos].Server))
			ChangeRegister(comando, planeta, ciudad, valor, caso, pos, addressF, port_grpcF, serverNuevo)
			time.Sleep(time.Second)
			ut.PrintRegister(registros)
		case 3:
			comando = "UpdateNumber"
			time.Sleep(time.Second)
			fmt.Println("\n|-----------UpdateNumber------------|")
			fmt.Print("Ingrese nombre del planeta: ")
			fmt.Scanln(&planeta)
			fmt.Print("Ingrese nombre de la ciudad a actualizar: ")
			fmt.Scanln(&ciudad)
			fmt.Print("Ingrese nuevo numero de rebeldes (debe ser distinto al actual): ")
			fmt.Scanln(&valor)
			time.Sleep(time.Second)

			pos = ut.ContainsPlanet(registros, planeta, ciudad)
			if (pos == -1){
				pos = len(registros)
				registros = append(registros, ut.RegisterInformante{NamePlanet: planeta, NameCity: ciudad, RelojPlanet: ut.Reloj{Server1: 0, Server2: 0, Server3: 0}, Server: 1})
			}

			addressF, port_grpcF, caso, serverNuevo = RequestServer(planeta, ciudad, registros[pos].RelojPlanet, int(registros[pos].Server))
			ChangeRegister(comando, planeta, ciudad, valor, caso, pos, addressF, port_grpcF, serverNuevo)
			time.Sleep(time.Second)
			ut.PrintRegister(registros)
		case 4:
			comando = "DeleteCity"
			time.Sleep(time.Second)
			fmt.Println("\n|-----------DeleteCity------------|")
			fmt.Print("Ingrese nombre del planeta: ")
			fmt.Scanln(&planeta)
			fmt.Print("Ingrese nombre de la ciudad a eliminar: ")
			fmt.Scanln(&ciudad)
			valor = ""
			time.Sleep(time.Second)

			pos = ut.ContainsPlanet(registros, planeta, ciudad)
			if (pos == -1){
				pos = len(registros)
				registros = append(registros, ut.RegisterInformante{NamePlanet: planeta, NameCity: ciudad, RelojPlanet: ut.Reloj{Server1: 0, Server2: 0, Server3: 0}, Server: 1})
			}

			addressF, port_grpcF, caso, serverNuevo = RequestServer(planeta, ciudad, registros[pos].RelojPlanet, int(registros[pos].Server))
			ChangeRegister(comando, planeta, ciudad, valor, caso, pos, addressF, port_grpcF, serverNuevo)
			time.Sleep(time.Second)
			ut.PrintRegister(registros)
		case 5:
			time.Sleep(time.Second)
			continuar = false
			fmt.Println("Adios Informante\n")
			time.Sleep(time.Second)
		default:
			time.Sleep(time.Second)
			fmt.Println("Ingrese una opcion valida\n")
			time.Sleep(time.Second)
		}
	}
	ut.PrintRegister(registros)
}

func GetRegistros() []ut.RegisterInformante{
	return registros
}