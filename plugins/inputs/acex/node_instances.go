package acex

import (
	"fmt"
	"net/http"
)

type NodeInstance struct {
	LogicalNodeID int    `json:"logical_node_id"`
	AssetRefType  string `json:"asset_ref_type"`
	AssetRefID    int    `json:"asset_ref_id"`
	ID            int    `json:"id"`
	Hostname      string `json:"hostname"`
	Site          string `json:"site"`
}

func (a *AcexPlugin) getNodeInstances() ([]NodeInstance, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/inventory/node_instances/", a.URL), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	var res []NodeInstance
	if err := a.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
