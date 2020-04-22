package rpc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// RPCAddr contains the url and a name for the node
type RPCAddr struct {
	host string
	port int
	name string
}

// CreateRPCAddr from string address and name
func CreateRPCAddr(addr, name string) (RPCAddr, error) {

	host, port, err := validateAddress(addr)
	if err != nil {
		return RPCAddr{}, err
	}

	if name == "" {
		name = addr
	}

	return RPCAddr{
		host: host,
		port: port,
		name: name,
	}, nil
}

func validateAddress(addr string) (host string, port int, err error) {
	parts := strings.Split(addr, ":")
	if len(parts) != 2 {
		return host, port, errors.New("Please provide addresses in the form of <ip or hostname>:<port>")
	}
	port, err = strconv.Atoi(parts[1])
	if err != nil {
		return host, port, fmt.Errorf("%s is not a valid port", parts[1])
	}
	host = parts[0]
	return
}

type NetInfo struct {
	Jsonrpc string `json:"jsonrpc"`
	// ID      string `json:"id"` // removed because sometimes it's "" and other times it's -1
	Result Result `json:"result"`
}
type ProtocolVersion struct {
	P2P   string `json:"p2p"`
	Block string `json:"block"`
	App   string `json:"app"`
}
type Other struct {
	TxIndex    string `json:"tx_index"`
	RPCAddress string `json:"rpc_address"`
}
type NodeInfo struct {
	ProtocolVersion ProtocolVersion `json:"protocol_version"`
	ID              string          `json:"id"`
	ListenAddr      string          `json:"listen_addr"`
	Network         string          `json:"network"`
	Version         string          `json:"version"`
	Channels        string          `json:"channels"`
	Moniker         string          `json:"moniker"`
	Other           Other           `json:"other"`
}
type SendMonitor struct {
	Start    time.Time `json:"Start"`
	Bytes    string    `json:"Bytes"`
	Samples  string    `json:"Samples"`
	InstRate string    `json:"InstRate"`
	CurRate  string    `json:"CurRate"`
	AvgRate  string    `json:"AvgRate"`
	PeakRate string    `json:"PeakRate"`
	BytesRem string    `json:"BytesRem"`
	Duration string    `json:"Duration"`
	Idle     string    `json:"Idle"`
	TimeRem  string    `json:"TimeRem"`
	Progress int       `json:"Progress"`
	Active   bool      `json:"Active"`
}
type RecvMonitor struct {
	Start    time.Time `json:"Start"`
	Bytes    string    `json:"Bytes"`
	Samples  string    `json:"Samples"`
	InstRate string    `json:"InstRate"`
	CurRate  string    `json:"CurRate"`
	AvgRate  string    `json:"AvgRate"`
	PeakRate string    `json:"PeakRate"`
	BytesRem string    `json:"BytesRem"`
	Duration string    `json:"Duration"`
	Idle     string    `json:"Idle"`
	TimeRem  string    `json:"TimeRem"`
	Progress int       `json:"Progress"`
	Active   bool      `json:"Active"`
}
type Channels struct {
	ID                int    `json:"ID"`
	SendQueueCapacity string `json:"SendQueueCapacity"`
	SendQueueSize     string `json:"SendQueueSize"`
	Priority          string `json:"Priority"`
	RecentlySent      string `json:"RecentlySent"`
}
type ConnectionStatus struct {
	Duration    string      `json:"Duration"`
	SendMonitor SendMonitor `json:"SendMonitor"`
	RecvMonitor RecvMonitor `json:"RecvMonitor"`
	Channels    []Channels  `json:"Channels"`
}
type Peers struct {
	NodeInfo         NodeInfo         `json:"node_info"`
	IsOutbound       bool             `json:"is_outbound"`
	ConnectionStatus ConnectionStatus `json:"connection_status"`
	RemoteIP         string           `json:"remote_ip"`
}
type Result struct {
	Listening bool     `json:"listening"`
	Listeners []string `json:"listeners"`
	NPeers    string   `json:"n_peers"`
	Peers     []Peers  `json:"peers"`
}
