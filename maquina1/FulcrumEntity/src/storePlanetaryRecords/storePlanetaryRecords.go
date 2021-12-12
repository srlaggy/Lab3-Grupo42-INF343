package storePlanetaryRecords

// FULCRUM

import (
	// "bufio"
	ut "lab/fulcrum/utils"
	"os"
	"strings"
	"fmt"
)

// Se tendrán las funciones:
// AddCity
// UpdateName nombre planeta nombre ciudad nuevo valor
// UpdateNumber nombre planeta nombre ciudad nuevo valor
// DeleteCity nombre planeta nombre ciudad

// var global
var listRegister []ut.Register

func Comands(str string) {
	s := strings.Split(str, " ")
	if (s[0] == "AddCity"){
		AddCity(s[1:])
	}
}

// AddCity nombre_planeta nombre_ciudad [nuevo valor]
//              s[0]          s[1]           s[2]
func AddCity(s []string){
	var fileName string = "utils/RegistroPlanetario/" + s[0] + ".txt"

	// Condición para verificar si el planeta existe o no:
	if _, err := os.Stat(fileName); err == nil {
	} else {
		// No existe
		file, err := os.Create(fileName)
		ut.FailOnError(err, "Failed to create file")
		defer file.Close()
		listRegister = append(listRegister, ut.Register{NamePlanet: s[0], RelojPlanet: ut.Reloj{Server1: 0, Server2: 0, Server3: 0}})
		fmt.Println("lista: ", listRegister)
	}

	// Agregar el dato de la ciudad
	file, error1 := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	ut.FailOnError(error1, "Failed to open file")
	var cityRegister string
	if (len(s) == 3){
		cityRegister =  s[1] + " " + s[2] + "\n"
	} else{
		cityRegister =  s[1] + " " + "0" + "\n"
	}
	_, error2 := file.WriteString(cityRegister)
	ut.FailOnError(error2, "Failed to write file")
	defer file.Close()
	
	// Cambiar el número en reloj vectorial del planeta	
	for i := 0; i < len(listRegister); i++ {
		if (listRegister[i].NamePlanet == s[0]){
			listRegister[i].RelojPlanet.Server1 = listRegister[i].RelojPlanet.Server1 + 1
		}
	}
	fmt.Println("lista: ", listRegister)
}

func GetListRegister() []ut.Register{
	return listRegister
}
