package core

import "github.com/andig/evcc/core/soc"

type adapter struct {
	lp *LoadPoint
}

func (lp *LoadPoint) adapter() soc.Adapter {
	return &adapter{lp: lp}
}

func (a *adapter) Publish(key string, val interface{}) {
	a.lp.publish(key, val)
}

func (a *adapter) SocEstimator() *soc.Estimator {
	return a.lp.socEstimator
}

func (a *adapter) ActivePhases() int64 {
	return a.lp.activePhases
}

func (a *adapter) Voltage() float64 {
	return Voltage
}
