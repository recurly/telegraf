package recurly

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/processors"
)

var sampleConfig = `
  ## Convert metrics to use a consistent standard.
  #
  ## By default, drop host names from metrics.
  # keep_hostnames = false
`

type Recurly struct {
	KeepHostNames bool
}

// SampleConfig descrbies the metrics to keep
func (r *Recurly) SampleConfig() string {
	return sampleConfig
}

// Description returns a description of the parser.
func (r *Recurly) Description() string {
	return "Apply metric modifications using override semantics."
}

// Apply parses and converts logs into metrics.
func (r *Recurly) Apply(in ...telegraf.Metric) []telegraf.Metric {
	for i, metric := range in {
		in[i] = r.prune(metric)
	}
	return in
}

func init() {
	processors.Add("recurly", func() telegraf.Processor {
		return &Recurly{}
	})
}
