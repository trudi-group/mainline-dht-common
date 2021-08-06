package common

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/anacrolix/dht/v2/krpc"
	log "github.com/sirupsen/logrus"
)

func ReportToFile(report *Output, path string) {
	var nodes []*MonitoredNodeJSON
	for _, node := range report.Nodes {
		jsonFormatted := MonitoredNodeJSON{UDPAddrs: node.UDPAddr, Reachable: node.Reachable, IPVersion: node.IPVersion.String()}
		nodes = append(nodes, &jsonFormatted)
	}

	MonitorOutput := MonitorOutputJSON{StartDate: report.StartDate, EndDate: report.EndDate, Nodes: nodes}

	vf, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
			log.Errorf("Error creating file %s\n", err)
			os.Exit(1)
		}
		vf, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.Errorf("Error opening file %s\n", err)
			os.Exit(1)
		}
	} else if err != nil {
		panic(err)
	}
	defer vf.Close()

	err = json.NewEncoder(vf).Encode(MonitorOutput)
	if err != nil {
		log.WithField("err", err).Error("Could not encode JSON and/or write to output file.")
	}
}

func CreateDirIfNotExists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.MkdirAll(path, 0777)
		return nil
	}
	return err
}

// Parses a file containing bootstrap peers. It assumes a text file with an IP:Port address on each line.
// It will ignore lines starting with a comment "//"
func ReadBootstrapListFromFile(path string) ([]*krpc.NodeInfo, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file line by line and parse the multiaddress string
	var bootstrapNI []*krpc.NodeInfo
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Ignore lines that are commented out
		if strings.HasPrefix(line, "//") {
			continue
		}
		ainfo, err := ParseAddrString(line)
		if err != nil {
			log.WithField("err", err).Error("Error parsing bootstrap peers.")
			return nil, err
		}
		bootstrapNI = append(bootstrapNI, ainfo)
	}

	return bootstrapNI, nil
}
