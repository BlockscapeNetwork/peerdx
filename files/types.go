package files

type AddrBook struct {
	Key   string `json:"key"`
	Addrs []Addr `json:"addrs"`
}

type Addr struct {
	Addr        AddrDetail `json:"addr"`
	Src         AddrDetail `json:"src"`
	Buckets     []int      `json:"buckets"`
	Attempts    int        `json:"attempts"`
	BucketType  int        `json:"bucket_type"`
	LastAttempt string     `json:"last_attempt"`
	LastSuccess string     `json:"last_success"`
}

type AddrDetail struct {
	ID   string `json:"id"`
	IP   string `json:"ip"`
	Port int    `json:"port"`
}
