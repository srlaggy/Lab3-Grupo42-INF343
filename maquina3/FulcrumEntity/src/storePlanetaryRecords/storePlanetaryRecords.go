package storePlanetaryRecords

// FULCRUM

import (
	// "bufio"
	ut "lab/fulcrum/utils"
	"os"
	"strings"
	"fmt"
	"log"
	"bytes"
	"bufio"
)

// Se tendrán las funciones:
// AddCity
// UpdateName
// UpdateNumber
// DeleteCity

// var global
var listRegister []ut.Register

func Comands(str string) {
	s := strings.Split(str, " ")
	if (s[0] == "AddCity"){
		AddCity(s[1:])
	} else if (s[0] == "UpdateName"){
		UpdateName(s[1:])
	} else if (s[0] == "UpdateNumber"){
		UpdateNumber(s[1:])
	} else if (s[0] == "DeleteCity"){
		DeleteCity(s[1:])
	}
}

// AddCity nombre_planeta nombre_ciudad [nuevo valor]
//              s[0]          s[1]           s[2]
func AddCity(s []string) {
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
			listRegister[i].RelojPlanet.Server1 = listRegister[i].RelojPlanet.Server3 + 1
		}
	}
	fmt.Println("lista: ", listRegister)
}


// UpdateName nombre_planeta nombre_ciudad [nuevo valor]
//                s[0]          s[1]           s[2]
func UpdateName(s []string) {
	var fileName string = "utils/RegistroPlanetario/" + s[0] + ".txt"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var bs []byte
	buf := bytes.NewBuffer(bs)

	scanner := bufio.NewScanner(f)
	var cantRebels string
	var scannerSplit []string
	for scanner.Scan() {
		scannerSplit = strings.Split(scanner.Text(), " ")
		if scannerSplit[0] != s[1] {
			_, err := buf.Write(scanner.Bytes())
			if err != nil {
				log.Fatal(err)
			}
			_, err = buf.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		} else{
			cantRebels = scannerSplit[1]
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(fileName, buf.Bytes(), 0666)
	if err != nil {
		log.Fatal(err)
	}
	
	// Agregar el dato de la ciudad
	file, error1 := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	ut.FailOnError(error1, "Failed to open file")
	_, error2 := file.WriteString(s[2] + " " + cantRebels + "\n")
	ut.FailOnError(error2, "Failed to write file")
	defer file.Close()
	
	// Cambiar el número en reloj vectorial del planeta	
	for i := 0; i < len(listRegister); i++ {
		if (listRegister[i].NamePlanet == s[0]){
			listRegister[i].RelojPlanet.Server1 = listRegister[i].RelojPlanet.Server3 + 1
		}
	}
	fmt.Println("lista: ", listRegister)
}


// UpdateNumber nombre_planeta nombre_ciudad [nuevo valor]
//                  s[0]          s[1]           s[2]
func UpdateNumber(s []string) {
	var fileName string = "utils/RegistroPlanetario/" + s[0] + ".txt"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var bs []byte
	buf := bytes.NewBuffer(bs)

	scanner := bufio.NewScanner(f)
	var scannerSplit []string
	for scanner.Scan() {
		scannerSplit = strings.Split(scanner.Text(), " ")
		if scannerSplit[0] != s[1] {
			_, err := buf.Write(scanner.Bytes())
			if err != nil {
				log.Fatal(err)
			}
			_, err = buf.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(fileName, buf.Bytes(), 0666)
	if err != nil {
		log.Fatal(err)
	}
	
	// Agregar el dato de la ciudad
	file, error1 := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	ut.FailOnError(error1, "Failed to open file")
	_, error2 := file.WriteString(s[1] + " " + s[2] + "\n")
	ut.FailOnError(error2, "Failed to write file")
	defer file.Close()
	
	// Cambiar el número en reloj vectorial del planeta	
	for i := 0; i < len(listRegister); i++ {
		if (listRegister[i].NamePlanet == s[0]){
			listRegister[i].RelojPlanet.Server1 = listRegister[i].RelojPlanet.Server3 + 1
		}
	}
	fmt.Println("lista: ", listRegister)
}


// DeleteCity nombre_planeta nombre_ciudad
//                 s[0]          s[1]
func DeleteCity(s []string) {
	var fileName string = "utils/RegistroPlanetario/" + s[0] + ".txt"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var bs []byte
	buf := bytes.NewBuffer(bs)

	scanner := bufio.NewScanner(f)
	var scannerSplit []string
	for scanner.Scan() {
		scannerSplit = strings.Split(scanner.Text(), " ")
		if scannerSplit[0] != s[1] {
			_, err := buf.Write(scanner.Bytes())
			if err != nil {
				log.Fatal(err)
			}
			_, err = buf.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(fileName, buf.Bytes(), 0666)
	if err != nil {
		log.Fatal(err)
	}
	
	// Cambiar el número en reloj vectorial del planeta	
	for i := 0; i < len(listRegister); i++ {
		if (listRegister[i].NamePlanet == s[0]){
			listRegister[i].RelojPlanet.Server1 = listRegister[i].RelojPlanet.Server3 + 1
		}
	}
	fmt.Println("lista: ", listRegister)
}


// Funcion GetListRegister
func GetListRegister() []ut.Register{
	return listRegister
}
