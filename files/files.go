package files

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// LoadAddressBooks loads address books in directory, will ignores nested directories,
// ignores files it can't open or marshal,
// stops the program if there is an issue while inspecting the directory.
func LoadAddressBooks(dir string) map[string]AddrBook {
	log.Println("Looking for address book json files at", dir)
	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println("Couldn't read files in", dir, ":", err)
		os.Exit(1)
	}

	abs := make(map[string]AddrBook)

	for _, fi := range fis {
		if fi.IsDir() { // disregard nested directories
			continue
		}

		fp := filepath.Join(dir, fi.Name())
		addrBook := &AddrBook{}
		if err := loadAndUnmarshal(fp, addrBook); err != nil {
			log.Printf("Couldn't read file %s: %s \nSkipping\n", fp, err.Error())
			continue
		}

		abs[fi.Name()] = *addrBook
	}

	return abs
}

// LoadFile loads file into provided pointer
func LoadFile(filepath string, out interface{}) error {
	return loadAndUnmarshal(filepath, out)
}

// loads the specified file and unmarshals its contents. out must be a pointer
func loadAndUnmarshal(filepath string, out interface{}) error {
	bz, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	return json.Unmarshal(bz, out)
}

// WriteJSONFile marshals the provided struct and writes it to filepath
func WriteJSONFile(filepath string, file interface{}) error {
	bz, err := json.MarshalIndent(file, "", " ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath, bz, 0644)
}
