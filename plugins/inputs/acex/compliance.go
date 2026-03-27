package acex

import (
	"fmt"
	"net/http"
	"time"

	"github.com/influxdata/telegraf"
)

type ComplianceMetric struct {
	TotalDesired      int     `json:"total_desired"`
	TotalObserved     int     `json:"total_observed"`
	CompliantCount    int     `json:"compliant_count"`
	CompliancePercent float64 `json:"compliance_percentage"`
}

func (a *AcexPlugin) gatherComplianceMetrics(ni NodeInstance, acc telegraf.Accumulator, ts time.Time) error {

	metricName := "acex_compliance"
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/operations/diff/%d", a.URL, ni.ID), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	var res ComplianceMetric
	if err := a.sendRequest(req, &res); err != nil {
		return err
	}

	tags := map[string]string{
		"node_instance_id": fmt.Sprintf("%d", ni.ID),
		"hostname":         ni.Hostname,
	}

	fields := map[string]any{
		"compliance_percentage": res.CompliancePercent,
		"total_desired":         res.TotalDesired,
		"total_observed":        res.TotalObserved,
		"compliant_count":       res.CompliantCount,
	}
	acc.AddFields(metricName, fields, tags, ts)

	return nil
}
