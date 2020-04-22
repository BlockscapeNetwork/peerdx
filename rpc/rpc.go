package rpc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/BlockscapeLab/peerdx/config"
)

// GetNetInfoAndCompare gets net_info via rpc from the provided addresses and compares the results
func GetNetInfoAndCompare(rpcAddrs []RPCAddr) {
	cfg := config.GetConfig()

	infos := make(map[string]NetInfo)
	for _, addr := range rpcAddrs {
		i, err := getNetInfo(addr)
		if err != nil {
			log.Printf("Couldn't get net info from %s check address and node setup: %s\n", fmt.Sprintf("%s:%d", addr.host, addr.port), err)
			continue
		}
		infos[addr.name] = i
	}

	peerList := compareNetInfo(infos)
	printResult(cfg, peerList)

}

func getNetInfo(addr RPCAddr) (NetInfo, error) {
	nodeURL := fmt.Sprintf("http://%s:%d/net_info", addr.host, addr.port)
	log.Println("Getting net info from", nodeURL)
	r, err := http.Get(nodeURL)
	if err != nil {
		return NetInfo{}, err
	}
	defer r.Body.Close()

	bz, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return NetInfo{}, err
	}
	ni := NetInfo{}
	err = json.Unmarshal(bz, &ni)
	return ni, err
}

func compareNetInfo(infosByNodes map[string]NetInfo) map[string][]string {
	peerList := make(map[string][]string)

	for name, info := range infosByNodes {
		for _, peer := range info.Result.Peers {
			nodeID := peer.NodeInfo.ID
			peerList[nodeID] = append(peerList[nodeID], name)
		}
	}
	return peerList
}

func printResult(info config.Config, ids map[string][]string) { //TODO put this somewhere else so addrbook and rpc can both use it
	// set size of name collum
	const idLen = 40
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

// ListDetailedPeerInfo prints detailed peer information to cmd line
func ListDetailedPeerInfo(addr RPCAddr) {
	ni, err := getNetInfo(addr)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	dp := createDetailedPeerInfo(&ni)
	printDetailedPeerInfo(dp)

}

func createDetailedPeerInfo(ni *NetInfo) displayData {
	dp := displayData{
		data: make([]detailData, 0, len(ni.Result.Peers)),
	}
	for _, peer := range ni.Result.Peers {

		direction := "IN"
		if peer.IsOutbound {
			direction = "OUT"
		}
		shortID := peer.NodeInfo.ID[:8]

		dd := detailData{
			rows: make([]string, 0, 3),
		}

		dd.label = fmt.Sprintf("%s %s [%s...]", direction, peer.NodeInfo.Moniker, shortID)

		dd.rows = append(dd.rows, fmt.Sprintf("IP %s", peer.RemoteIP))

		send := "inactive"
		recv := "inactive"
		if peer.ConnectionStatus.SendMonitor.Active {
			send = "active"
		}
		if peer.ConnectionStatus.RecvMonitor.Active {
			recv = "active"
		}
		dd.rows = append(dd.rows, fmt.Sprintf("send %s; receive %s", send, recv))
		dd.rows = append(dd.rows, fmt.Sprintf("tendermint version %s", peer.NodeInfo.Version))

		dp.AddData(dd)
	}
	return dp
}

func printDetailedPeerInfo(dp displayData) {
	for _, d := range dp.data {
		paddedLabel := addWhiteSpace(dp.maxLabelSize, d.label)
		fmt.Printf("%s: %s\n", paddedLabel, d.rows[0])
		if len(d.rows) < 2 {
			continue
		}
		for _, r := range d.rows[1:] {
			fmt.Printf("%s  %s\n", addWhiteSpace(dp.maxLabelSize, ""), r) // two additional spaces. One for colon, one for space after colon
		}
		fmt.Println("")
	}
}

type displayData struct {
	data         []detailData
	maxLabelSize int
}

func (dp *displayData) AddData(data detailData) {
	labelLen := len(data.label)
	if dp.maxLabelSize < labelLen {
		dp.maxLabelSize = labelLen
	}
	dp.data = append(dp.data, data)
}

type detailData struct {
	label string
	rows  []string
}
