package healthcheck

import (
	"github.com/innogames/yacht/logger"
	"sync"
)

// HCScript stores all properties of a script healthcheck.
type HCScript struct {
	HealthCheckBase
	Script string
}

// NewHCScript creates new script healthcheck struct and populates it with data from Json config.
func newHCScript(json JSONMap) *HCScript {
	hc := new(HCScript)
	hc.hcType = json["type"].(string)
	hc.Script = json["script"].(string)
	logger.Info.Printf("healthcheck: %s", hc.hcType)
	return hc
}

func (this HCScript) Run(wg *sync.WaitGroup) {
	this.HealthCheckBase.Run(wg)
}

func (this HCScript) Stop() {
	this.HealthCheckBase.Stop()
}
