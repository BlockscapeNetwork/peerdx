package addrbook

import (
	"fmt"
	"log"

	"github.com/BlockscapeLab/peerdx/config"

	"github.com/BlockscapeLab/peerdx/files"
)

const idLen = 40

// LoadAndCompare loads addrBooks from dir and compares them
func LoadAndCompare(addrBookDir string) {
	abs := files.LoadAddressBooks(addrBookDir)
	cfg := config.GetConfig()
	ids := compareAddrBooks(abs)
	printResult(cfg, ids)
}

// returns a map of ids and the address books that contain them
func compareAddrBooks(abs map[string]files.AddrBook) map[string][]string {
	log.Println("Analyzing address books")

	ids := make(map[string][]string)

	for name, adressBook := range abs {
		for _, addr := range adressBook.Addrs {
			id := addr.Addr.ID
			if names, ok := ids[id]; ok {
				ids[id] = append(names, name)
			} else {
				ids[id] = []string{name}
			}
		}
	}

	return ids
}

func printResult(info config.Config, ids map[string][]string) { //TODO put this somewhere else so addrbook and rpc can both use it
	// set size of name collum
	nameColSize := info.GetMaxNameLength()
	if idLen > nameColSize {
		nameColSize = idLen
	}
	log.Printf("A total of %d different addresses:\n", len(ids))

	for id, names := range ids {
		namelist := ""
		for _, n := range names {
			if namelist == "" {
				namelist = n
			} else {
				namelist = fmt.Sprintf("%s, %s", namelist, n)
			}
		}
		if name, ok := info.GetNameForID(id); ok {
			log.Printf("%s: %s\n", addWhiteSpace(nameColSize, name), namelist)
		} else {
			log.Printf("%s: %s\n", addWhiteSpace(nameColSize, id), namelist)
		}
	}
}

func addWhiteSpace(targetLength int, original string) string {
	diff := targetLength - len(original)
	if diff < 1 {
		return original
	}

	for i := 0; i < diff; i++ {
		original = original + " "
	}

	return original
}
