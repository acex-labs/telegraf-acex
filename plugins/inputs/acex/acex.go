package acex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/config"
	common_http "github.com/influxdata/telegraf/plugins/common/http"
	"github.com/influxdata/telegraf/plugins/inputs"
)

var Version = "dev"

type Metric struct {
	Name   string
	Tags   map[string]string
	Fields map[string]any
}

type AcexPlugin struct {
	URL string `toml:"url"`

	// Auth
	Username config.Secret `toml:"username"`
	Password config.Secret `toml:"password"`

	Token     config.Secret `toml:"token"`
	TokenFile string        `toml:"token_file"`

	Headers map[string]*config.Secret `toml:"headers"`

	SuccessStatusCodes []int `toml:"success_status_codes"`

	Query string `toml:"query"`

	Log *StderrLogger `toml:"-"`

	common_http.HTTPClientConfig

	client *http.Client
}

func (a *AcexPlugin) Description() string {
	return "Acex Input plugin"
}

func (a *AcexPlugin) SampleConfig() string {
	return `
  # AcexPlugin minimal config
  [[inputs.Acex]]
    ## Optional URL if you plan to use HTTP later
    url = "http://localhost:8080"
`
}

func (a *AcexPlugin) Init() error {
	// Validate token config
	if a.TokenFile != "" && !a.Token.Empty() {
		return errors.New("either use 'token_file' or 'token', not both")
	}

	// Create HTTP client
	ctx := context.Background()
	client, err := a.HTTPClientConfig.CreateClient(ctx, a.Log)
	if err != nil {
		return err
	}
	a.client = client

	// Default success codes
	if len(a.SuccessStatusCodes) == 0 {
		a.SuccessStatusCodes = []int{200}
	}
	a.Log.Infof("Starting external Acex Input plugin '%s' version by Acebit", Version)

	return nil
}

func (a *AcexPlugin) Gather(acc telegraf.Accumulator) error {
	nodeInstances, err := a.getNodeInstances()
	if err != nil {
		return err
	}

	ts := time.Now()
	var wg sync.WaitGroup

	for _, nodeInstance := range nodeInstances {
		ni := nodeInstance

		wg.Add(1)
		go func() {
			defer wg.Done()

			a.Log.Debugf("Fetching metrics for Node ID: %d", ni.ID)

			err := a.gatherComplianceMetrics(ni, acc, ts)
			if err != nil {
				a.Log.Error(fmt.Sprintf("Error fetching metrics for node %d: %v", ni.ID, err))
			}
		}()
	}

	wg.Wait()
	return nil
}

func init() {
	inputs.Add("acex", func() telegraf.Input {
		return &AcexPlugin{}
	})
}
