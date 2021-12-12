package utils

import (
	"log"
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