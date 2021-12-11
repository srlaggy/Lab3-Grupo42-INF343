package storePlanetaryRecords

// FULCRUM

import (
	// "bufio"
	ut "lab/fulcrum/utils"
	"os"
	"strings"
)

// Se tendrán las funciones:
// AddCity
// UpdateName nombre planeta nombre ciudad nuevo valor
// UpdateNumber nombre planeta nombre ciudad nuevo valor
// DeleteCity nombre planeta nombre ciudad

func main() {
    
}


// AddCity nombre planeta nombre ciudad [nuevo valor]
func AddCity(register string){
	s := strings.Split(register, " ")
	var fileName string = "utils/RegistroPlanetario/" + s[0] + ".txt"
	file, err := os.Create(fileName)
	ut.FailOnError(err, "Failed to create file")
	//FIXME: undeclared name: failOnError
	defer file.Close()
}


