package generate

import (
	"os"
	"log"
	"github.com/Nygard-R1/utils/error"
)

func Generate(args []string) {
	schemaFile := args[0]

	if len(args) != 1 {
		log.Fatal("This program takes exactly 1 argument. Please provide schema file location")
	}
	log.Printf("About to parse schema from file: %s", schemaFile)

	f, err:= os.Open(args[0])
	error.PassOrExit(err)
}
