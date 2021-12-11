package utils

import (
	"log"
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