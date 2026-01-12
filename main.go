package odinfleet

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"odinfleet-nakama/api"
	"strconv"

	"github.com/heroiclabs/nakama-common/runtime"
)

// ServerToInstanceInfoConverter Converts server struct info instance info struct required by fleet api
type ServerToInstanceInfoConverter func(fm *OdinFleetManager, server *api.Server) (instanceInfo *runtime.InstanceInfo, err error)

type OdinFleetManager struct {
	config                            *OdinFleetManagerConfig
	serverToInstanceInfoConverterFunc ServerToInstanceInfoConverter

	client *api.APIClient
	nk     runtime.NakamaModule
	logger runtime.Logger
}

type OdinFleetManagerConfig struct {
	AppId     int32
	ApiToken  string
	ClientCfg *api.Configuration
}

// GetDefaultConfiguration Generate configuration for ODIN Fleet API
func GetDefaultConfiguration(appId int32, apiToken string) *OdinFleetManagerConfig {
	cfg := api.NewConfiguration()
	cfg.Host = "fleet.4players.io"
	cfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %v", apiToken))

	return &OdinFleetManagerConfig{
		AppId:     appId,
		ClientCfg: cfg,
	}
}

// CreateOdinFleetManager Create ready to use fleet manager communicating with ODIN Fleet API
func CreateOdinFleetManager(logger runtime.Logger, config *OdinFleetManagerConfig, converterFunc ServerToInstanceInfoConverter) (fm *OdinFleetManager, err error) {
	if config == nil {
		return nil, errors.New("ODIN: Fleet manager config is missing")
	}

	if converterFunc == nil {
		logger.Warn("ODIN: No implementation for ServerToInstanceInfoConverter given, will use default implementation")
		converterFunc = ServerToInstanceInfo
	}
	client := api.NewAPIClient(config.ClientCfg)
	return &OdinFleetManager{
		serverToInstanceInfoConverterFunc: converterFunc,
		client:                            client,
		logger:                            logger,
		config:                            config,
	}, nil
}

// TODO: Is callback needed?
func (fm OdinFleetManager) Init(nk runtime.NakamaModule, callbackHandler runtime.FmCallbackHandler) error {
	fm.nk = nk
	return nil
}

func (fm OdinFleetManager) Update(ctx context.Context, id string, playerCount int, metadata map[string]any) error {
	serviceId, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("ODIN: Invalid service id given")
	}

	request := fm.client.DockerAPI.DockerServicesMetadataUpdate(ctx, int32(serviceId)).PatchMetadataRequest(api.PatchMetadataRequest{Metadata: metadata})
	_, _, err = request.PatchMetadataRequest(api.PatchMetadataRequest{Metadata: metadata}).Execute()
	if err != nil {
		return err
	}
	return nil
}

// TODO: Odin configuration call for deleting instance directly (removing them from deployment)
func (fm OdinFleetManager) Delete(ctx context.Context, id string) error {
	return errors.New("ODIN: Deleting specific instance not yet implemented")
}

func (fm OdinFleetManager) Get(ctx context.Context, id string) (instance *runtime.InstanceInfo, err error) {
	serviceId, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("ODIN: Invalid instance id given")
	}
	server, _, err := fm.client.DockerAPI.GetServerById(ctx, fm.config.AppId, int32(serviceId)).Execute()
	if err != nil {
		return nil, errors.New("ODIN: Could not fetch ODIN fleet instance")
	}
	return fm.serverToInstanceInfoConverterFunc(&fm, server)
}

func (fm OdinFleetManager) List(ctx context.Context, query string, limit int, previousCursor string) (
	instances []*runtime.InstanceInfo,
	nextCursor string,
	err error,
) {
	params := &GetServersParams{}
	if err := json.Unmarshal([]byte(query), params); err != nil {
		params = nil
	}
	servers, nextPage, err := fm.GetServers(ctx, params)
	if err != nil {
		return nil, "", err
	}

	instances = []*runtime.InstanceInfo{}
	for _, s := range servers {
		instance, err := fm.serverToInstanceInfoConverterFunc(&fm, &s)
		if err != nil {
			continue
		}
		instances = append(instances, instance)
	}
	if nextPage != 0 {
		nextCursor = strconv.Itoa(int(nextPage))
	}

	return instances, nextCursor, nil
}

// GetServersParams Expected as argument in `query` parameter in `List` method
type GetServersParams struct {
	Pages   int32
	PerPage int32

	Metadata map[string]any
	Status   *string
	Location *FleetLocation
}

type FleetLocation struct {
	Continent *string
	City      *string
	Country   *string
}

func (fm OdinFleetManager) GetServers(ctx context.Context, params *GetServersParams) (
	servers []api.Server,
	nextPage int32,
	err error,
) {

	request := fm.generateGetServersRequest(ctx, params)
	response, _, err := request.Execute()
	if err != nil {
		return nil, 0, err
	}
	if p := params; p != nil && p.Pages != 0 {
		nextPage = p.Pages + 1
	}
	return response.Data, nextPage, nil
}

func (fm OdinFleetManager) generateGetServersRequest(ctx context.Context, params *GetServersParams) api.DockerAPIGetServersRequest {
	request := fm.client.DockerAPI.
		GetServers(ctx, fm.config.AppId)

	if params == nil {
		return request
	}

	if perPage := params.PerPage; perPage > 0 {
		request = request.PerPage(perPage)
	}
	if pages := params.Pages; pages > 0 {
		request = request.Page(pages)
	}

	if m := params.Metadata; len(m) > 0 {
		metadata := createQuery(m)
		request = request.FilterMetadata(metadata)
	}
	if l := params.Location; l != nil {
		if city := l.City; city != nil {
			request = request.FilterLocationCity(*city)
		}
		if country := l.Country; country != nil {
			request = request.FilterLocationCountry(*country)
		}
		if continent := l.Continent; continent != nil {
			request = request.FilterLocationContinent(*continent)
		}
	}

	if status := params.Status; status != nil {
		request = request.FilterStatus(*status)
	}
	return request
}

func (fm OdinFleetManager) Create(ctx context.Context, maxPlayers int, userIds []string, latencies []runtime.FleetUserLatencies, metadata map[string]any, callback runtime.FmCreateCallbackFn) (err error) {
	return errors.New("ODIN: Creating specific ODIN instance not supported yet")
}

func (fm OdinFleetManager) Join(ctx context.Context, id string, userIds []string, metadata map[string]string) (joinInfo *runtime.JoinInfo, err error) {
	return nil, errors.New("ODIN: Join workflow not defined yet")
}
