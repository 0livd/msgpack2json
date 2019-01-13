package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/vmihailenco/msgpack"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("%s takes a list of message pack encoded files and converts them to json\n", os.Args[0])
		fmt.Printf("Usage: %s in-file...\n", os.Args[0])
	}
	for _, msgpackFilename := range os.Args[1:] {
		data, err := ioutil.ReadFile(msgpackFilename)
		if err != nil {
			fmt.Printf("Cannot read input file %s with error %v\n", msgpackFilename, err)
			continue
		}
		var unpackedData interface{}
		err = msgpack.Unmarshal(data, &unpackedData)
		if err != nil {
			fmt.Printf("Cannot unmarshal msgpack file %s with error %v\n", msgpackFilename, err)
			continue
		}
		jsonFilename := msgpackFilename[0:len(msgpackFilename)-len(filepath.Ext(msgpackFilename))] + ".json"
		jsonFile, err := os.Create(jsonFilename)
		if err != nil {
			fmt.Printf("Cannot create output file %s for input file %s with error %v\n",
				jsonFilename, msgpackFilename, err)
			continue
		}
		defer jsonFile.Close()
		jsonOut, err := json.MarshalIndent(unpackedData, "", "    ")
		if err != nil {
			fmt.Printf("Cannot convert msgpack to json for input file %s with error %v\n",
				msgpackFilename, err)
			continue
		}
		_, err = jsonFile.Write(jsonOut)
		if err != nil {
			fmt.Printf("Cannot write to output file %s for input file %s with error %v\n",
				jsonFilename, msgpackFilename, err)
			continue
		}
		fmt.Printf("Successfully converted %s to %s\n", msgpackFilename, jsonFilename)
	}
}
