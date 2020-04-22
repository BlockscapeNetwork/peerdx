package config

type Config struct {
	KnownIds []KnownId `json:"known_ids"`
}
type KnownId struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetNameForID will return the name alias for an ID and true if known, or else "" and false
func (c Config) GetNameForID(id string) (string, bool) {
	for _, ki := range c.KnownIds {
		if ki.ID == id {
			return ki.Name, true
		}
	}
	return "", false
}

// GetMaxNameLength returns the length of the longest name
func (c Config) GetMaxNameLength() int {
	l := 0
	for _, ki := range c.KnownIds {
		nl := len(ki.Name)
		if nl > l {
			l = nl
		}
	}
	return l
}
