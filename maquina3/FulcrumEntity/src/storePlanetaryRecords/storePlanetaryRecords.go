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
// (0) Commands
// (1) AddCity			-> Se agrega la ciudad al registro de su planeta
// (2) UpdateName		-> Se actualiza el nombre de la ciudad del registro de su planeta
// (3) UpdateNumber		-> Se actualiza la cantidad de rebeldes de la ciudad del registro de su planeta
// (4) DeleteCity		-> Se elimina la ciudad del registro de su planeta
// (5) CreateRecordsLog	-> Se crea el archivo de Log de Registro
// (6) CleanRecordsLog	-> Se limpia el contenido de los logs de cada planeta
// (7) RecordsLog		-> Se agrega el comando enviado por los informantes al Log de Registro
// (8) GetListRegister	-> Se retorna la lista de registros de los planetas

// var global
var listRegister []ut.Register

// (0) Commands: se verifica qué tipo de comando envió el informante y se registra tal comando
func Commands(str string) {
	s := strings.Split(str, " ")
	if (s[0] == "AddCity"){
		CreateRecordsLog(s[1])
		RecordsLog(str, s[1])
		AddCity(s[1:])
	} else if (s[0] == "UpdateName"){
		RecordsLog(str, s[1])
		UpdateName(s[1:])
	} else if (s[0] == "UpdateNumber"){
		RecordsLog(str, s[1])
		UpdateNumber(s[1:])
	} else if (s[0] == "DeleteCity"){
		RecordsLog(str, s[1])
		DeleteCity(s[1:])
	}
}


// ----------------------------------------------------------------------- //
// ------------------ Comandos enviados por Informantes ------------------ //
// ----------------------------------------------------------------------- //

// (1) AddCity: 		nombre_planeta nombre_ciudad [nuevo valor]
//              			s[0]          s[1]           s[2]
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
	}

	// Agregar el dato de la ciudad
	file, error1 := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	ut.FailOnError(error1, "Failed to open file")
	var cityRegister string
	if (len(s) == 3){
		cityRegister = s[0] + " " +  s[1] + " " + s[2] + "\n"
	} else{
		cityRegister = s[0] + " " + s[1] + " " + "0" + "\n"
	}
	_, error2 := file.WriteString(cityRegister)
	ut.FailOnError(error2, "Failed to write file")
	defer file.Close()
	
	// Cambiar el número en reloj vectorial del planeta	
	for i := 0; i < len(listRegister); i++ {
		if (listRegister[i].NamePlanet == s[0]){
			listRegister[i].RelojPlanet.Server3 = listRegister[i].RelojPlanet.Server3 + 1
		}
	}
	fmt.Println("lista: ", listRegister)
}


// (2) UpdateName: 		nombre_planeta nombre_ciudad [nuevo valor]
//                			s[0]          s[1]           s[2]
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
	_, error2 := file.WriteString(s[0] + " " + s[2] + " " + cantRebels + "\n")
	ut.FailOnError(error2, "Failed to write file")
	defer file.Close()
	
	// Cambiar el número en reloj vectorial del planeta	
	for i := 0; i < len(listRegister); i++ {
		if (listRegister[i].NamePlanet == s[0]){
			listRegister[i].RelojPlanet.Server3 = listRegister[i].RelojPlanet.Server3 + 1
		}
	}
	fmt.Println("lista: ", listRegister)
}


// (3) UpdateNumber: 	nombre_planeta nombre_ciudad [nuevo valor]
//                  		s[0]          s[1]           s[2]
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
	_, error2 := file.WriteString(s[0] + " " + s[1] + " " + s[2] + "\n")
	ut.FailOnError(error2, "Failed to write file")
	defer file.Close()
	
	// Cambiar el número en reloj vectorial del planeta	
	for i := 0; i < len(listRegister); i++ {
		if (listRegister[i].NamePlanet == s[0]){
			listRegister[i].RelojPlanet.Server3 = listRegister[i].RelojPlanet.Server3 + 1
		}
	}
	fmt.Println("lista: ", listRegister)
}


// (4) DeleteCity: 		nombre_planeta nombre_ciudad
//                 			s[0]          s[1]
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
			listRegister[i].RelojPlanet.Server3 = listRegister[i].RelojPlanet.Server3 + 1
		}
	}
	fmt.Println("lista: ", listRegister)
}



// ----------------------------------------------------------------------- //
// --------------------------- Log de Registro --------------------------- //
// ----------------------------------------------------------------------- //

// (5) CreateRecordsLog: se crea el archivo de Log de registro por cada planeta que va ingresando
func CreateRecordsLog(planeta string) {
	var fileName string = "utils/RecordsLog/" + planeta + ".txt"
	// Condición para verificar si el planeta existe o no en utils/RecordsLog/:
	if _, err := os.Stat(fileName); err == nil {
		// Si existe, no se vuelve a crear
	} else {
		// No existe, se crea entonces
		file, err := os.Create(fileName)
		ut.FailOnError(err, "Failed to create file")
		defer file.Close()
	}
}


// (6) CleanRecordsLog: se elimina el contenido de todos los log de los planetas
func CleanRecordsLog() {
	for i := 0; i < len(listRegister); i++ {
		var fileName string = "utils/RecordsLog/" + listRegister[i].NamePlanet + ".txt"
		f, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		var bs []byte
		buf := bytes.NewBuffer(bs)

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if scanner.Text() == "\n" {
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
	}
}


// (7) RecordsLog: se agrega cada Command en el archivo de Log de Registro
func RecordsLog(str string, planeta string) {
	var fileName string = "utils/RecordsLog/" + planeta + ".txt"
	file, error1 := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	ut.FailOnError(error1, "Failed to open file")
	_, error2 := file.WriteString(str + "\n")
	ut.FailOnError(error2, "Failed to write file")
	defer file.Close()
}



// ----------------------------------------------------------------------- //
// --------------------------- Otras funciones --------------------------- //
// ----------------------------------------------------------------------- //

// (8) Funcion GetListRegister
func GetListRegister() []ut.Register{
	return listRegister
}
