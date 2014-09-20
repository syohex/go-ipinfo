package ipinfo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

type Info struct {
	IP           net.IP `json:"ip"`
	HostName     string `json:"hostname"`
	City         string `json:"city"`
	Region       string `json:"region"`
	Country      string `json:"country"`
	Location     string `json:"loc"`
	Organization string `json:"org"`
}

func IPInfo(ip net.IP) *Info {
	url := fmt.Sprintf("http://ipinfo.io/%s/json", ip.String())
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	var info Info
	if err := json.Unmarshal(bytes, &info); err != nil {
		return nil
	}

	return &info
}
