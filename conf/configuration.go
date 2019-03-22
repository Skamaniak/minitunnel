package conf

import (
	"encoding/json"
	"os"
)

type TunnelConfiguration struct {
	IpAddresses string
}

type ListenerConfiguration struct {
	Host           string
	Port           uint16
	ConnectionType string
}

type Configuration struct {
	Tunnel   TunnelConfiguration
	Listener ListenerConfiguration
}

func ReadConfiguration(fileName string) (*Configuration, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(file)

	configuration := &Configuration{}
	err = decoder.Decode(configuration)
	if err != nil {
		return nil, err
	}
	return configuration, nil
}
