package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

// timeTrack is used to debug each function by measuring how long it takes to execute.
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	if *debug == true {
		log.Printf("%s took %s", name, elapsed)
	}
}

// stringToSha256 converts a string to sha 256.
func stringToSha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	str := hex.EncodeToString(bs)
	return str
}

// fileToString Reads a file into a sting.
func fileToString(fileName string) string {
	defer timeTrack(time.Now(), "fileToString")
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Println("ERROR: The file/path", fileName, "does not exist here")
		os.Exit(-1)
	}
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(dat)
}

// searchInStringToMap Reads a string and returns all lowercased matches in the regular expression as map keys
func searchInStringToMap(total string, expression string) map[string]bool {
	defer timeTrack(time.Now(), "searchInStringToMap")
	r, err := regexp.Compile(expression)
	if err != nil {
		panic(err)
	}
	allMatches := r.FindAllString(total, -1)
	a := make(map[string]bool)
	for _, v := range allMatches {
		a[strings.ToLower(v)] = false
	}
	return a
}

// searchInStringToMapUC Reads a string and returns all matches upper case in the regular expression as map keys
func searchInStringToMapUC(total string, expression string) map[string]bool {
	defer timeTrack(time.Now(), "searchInStringToMapCS")
	r, err := regexp.Compile(expression)
	if err != nil {
		panic(err)
	}
	allMatches := r.FindAllString(total, -1)
	a := make(map[string]bool)
	for _, v := range allMatches {
		a[strings.ToUpper(v)] = false
	}
	return a
}

// Compare Compares 2 maps with words as what to search and boleans false value. Transforms in true when the key exists in the other map.
func Compare(myMap map[string]bool, encMap map[string]bool) (map[string]bool, map[string]bool) {
	defer timeTrack(time.Now(), "Compare")
	var y bool
	for key := range myMap {
		_, y = encMap[stringToSha256(key)]
		if y == true {
			myMap[key] = true
		}
	}
	return myMap, encMap
}

// mapKeysToSlice Adds the keys with val to a slice
func mapKeysToSlice(m map[string]bool, val bool) []string {
	defer timeTrack(time.Now(), "mapKeysToSlice")
	var result []string
	for k, v := range m {
		if v == val {
			result = append(result, k)
		}
	}
	return result
}

// stringToFile writes a string to a file.
func stringToFile(fileName string, dat string) {
	defer timeTrack(time.Now(), "stringToFile")
	err := ioutil.WriteFile(fileName, []byte(dat), 0644)
	if err != nil {
		panic(err)
	}
}

// trashFiles deletes files created by ecompare
func trashFiles() {
	os.Remove("was-not-found.txt")
	os.Remove("was-found.txt")
}
