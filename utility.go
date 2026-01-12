package odinfleet

import (
	"fmt"
	"odinfleet-nakama/api"
	"strconv"
	"strings"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

// ServerToInstanceInfo Current default implementation for ServerToInstanceInfoConverter
func ServerToInstanceInfo(fm *OdinFleetManager, server *api.Server) (*runtime.InstanceInfo, error) {
	id := strconv.Itoa(int(server.Id))
	metadata := server.Metadata
	for _, port := range server.Ports {
		if port.PublishedPort.IsSet() == false {
			continue
		}
		metadata[port.Name] = *port.PublishedPort.Get()
	}
	connectionInfo, err := nodeToConnectionInfo(server, "game_port")
	if err != nil {
		return nil, err
	}

	// TODO: Parameters
	status := server.Status

	var createdTime time.Time
	if t := server.CreatedAt.Get(); t != nil {
		createdTime = *t
	} else {
		createdTime = time.Now().UTC()
	}

	return &runtime.InstanceInfo{
		Id:             id,
		ConnectionInfo: connectionInfo,
		CreateTime:     createdTime,
		PlayerCount:    0,
		Status:         status,
		Metadata:       metadata,
	}, nil
}

func nodeToConnectionInfo(server *api.Server, portName string) (connection *runtime.ConnectionInfo, err error) {
	port, portExists := server.Metadata[portName]
	if portExists == false {
		return nil, fmt.Errorf("ODIN: No port with name %v was found", portName)
	}
	return &runtime.ConnectionInfo{
		IpAddress: server.Node.Address,
		Port:      (int)(port.(int32)),
	}, nil
}

func createQuery(metadata map[string]interface{}) string {
	builder := strings.Builder{}
	for k, v := range metadata {
		builder.WriteString(fmt.Sprintf("%s=\"%v\"", k, v))
	}
	return builder.String()
}
