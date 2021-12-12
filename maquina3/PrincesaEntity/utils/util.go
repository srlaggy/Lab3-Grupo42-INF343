package utils

import (
	"log"
	"fmt"
)

type Reloj struct {
	Server1 int64
	Server2 int64
	Server3 int64
}

type RegisterPrincesa struct{
	NamePlanet string
	NameCity string
	RelojPlanet Reloj
	Server int64
}

func CreateReloj(s1 int64, s2 int64, s3 int64) *Reloj{
	r := Reloj{Server1: s1, Server2: s2, Server3: s3}
	return &r
}

func CreateDir(protocolo string, address string, port string) string{
	if protocolo == ""{
		return address + ":" + port
	} else if address == ""{
		return protocolo + ":" + port
	} else {
		return protocolo + "://" + address + ":" + port
	}
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// retorna -1 si no existe el elemento
// retorna i si existe, con i = posicion
func ContainsPlanet(slice []RegisterPrincesa, planeta string, ciudad string) int{
    for i, elemento := range slice {
        if (elemento.NamePlanet == planeta) && (elemento.NameCity == ciudad){
            return i
        }
    }
    return -1
}

func PrintRegister(slice []RegisterPrincesa){
	for i, elemento := range slice{
		fmt.Println("\n--------- REGISTRO ", i+1, " -----------")
		fmt.Println("Planeta: ", elemento.NamePlanet)
		fmt.Println("Ciudad: ", elemento.NameCity)
		fmt.Println("Reloj: ", elemento.RelojPlanet.Server1, "-", elemento.RelojPlanet.Server2, "-", elemento.RelojPlanet.Server3)
		fmt.Println("Server: ", elemento.Server)
	}
}