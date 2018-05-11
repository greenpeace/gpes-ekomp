package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

const emailRegex string = `([a-zA-Z0-9_\-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([a-zA-Z0-9\-]+\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\]?)`

const shaRegex string = `[A-Fa-f0-9]{64}`

const dninieRegex string = `[A-z]?\d{7,8}[TRWAGMYFPDXBNJZSQVHLCKEtrwagmyfpdxbnjzsqvhlcke]`

var debug *bool

func main() {

	defer timeTrack(time.Now(), "main")

	help := flag.Bool("help", false, "Display help")
	trash := flag.Bool("trash", false, "Delete files created by ekomp")
	data := flag.String("data", "emails", "What type of data")
	myFilename := flag.String("file", "myfile.txt", "File to do the operations")
	encFilename := flag.String("list", "encrypted.txt", "Encrypted list")
	debug = flag.Bool("debug", false, "Debug the script")
	flag.Parse()

	if *help == true {
		helpMe()
	} else if *trash == true {
		trashFiles()
	} else {

		// Read the files to strings
		myFile := fileToString(*myFilename)
		encFile := fileToString(*encFilename)

		// Create a map from the emails in each file with value as false
		var myMap, encMap map[string]bool
		switch *data {
		case "emails":
			myMap = searchInStringToMap(myFile, emailRegex)
			encMap = searchInStringToMap(encFile, shaRegex)
		case "dni":
			myMap = searchInStringToMapUC(myFile, dninieRegex)
			encMap = searchInStringToMap(encFile, shaRegex)
		default:
			myMap = searchInStringToMap(myFile, emailRegex)
			encMap = searchInStringToMap(encFile, shaRegex)
		}

		// Transforms the values in the map to true when the key exits in the other map
		Compare(myMap, encMap)

		wasNotFound := mapKeysToSlice(myMap, false)
		wasFound := mapKeysToSlice(myMap, true)

		stringToFile("was-not-found.txt", strings.Join(wasNotFound, "\n"))
		stringToFile("was-found.txt", strings.Join(wasFound, "\n"))

		fmt.Printf("\nWHAT HAPPENED?\n\n")
		fmt.Println("Your file:", *myFilename)
		fmt.Println("Encrypted file:", *encFilename)
		fmt.Println("Parsed", *data, "in", *myFilename, ":", len(myMap))
		fmt.Println("Parsed sha256 in", *encFilename, ":", len(encMap))
		fmt.Println("Was not found in", *myFilename, "when comparing to", *encFilename, ":", len(wasNotFound), *data)
		fmt.Println("Was found in", *myFilename, "when comparing to", *encFilename, ":", len(wasFound), *data)
		fmt.Printf("\nCHECK THE FILES:\nwas-found.txt\nwas-not-found.txt\nfor more information.\n\n")
	}

}
