package recurly

import (
	"testing"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/metric"
)

func createTestMetric(k, v string) telegraf.Metric {
	metric, _ := metric.New(
		"recurly",
		map[string]string{k: v},
		map[string]interface{}{"value": int64(1)},
		time.Now(),
	)
	return metric
}

func calculateProcessedTags(processor Recurly, metric telegraf.Metric) map[string]string {
	processed := processor.Apply(metric)
	return processed[0].Tags()
}

type ts struct {
	tag             string
	wantReplacement string
	wantRemoved     bool
	keepHosts       bool
}

func TestPruner(t *testing.T) {
	tC := []ts{
		{
			tag:         "hostname",
			wantRemoved: true,
		},
		{
			tag:         "hostName",
			wantRemoved: true,
		},
		{
			tag:         "host_name",
			wantRemoved: true,
		},
		{
			tag:             "host_name",
			wantReplacement: "host",
			keepHosts:       true,
		},
	}

	for k, v := range keepTags {
		var wr string
		if k != v {
			wr = v
		}
		tC = append(tC, ts{
			tag:             k,
			wantReplacement: wr,
		})

	}

	for _, tc := range tC {
		processor := Recurly{
			KeepHostNames: tc.keepHosts,
		}
		tags := calculateProcessedTags(processor, createTestMetric(tc.tag, tc.tag))
		_, remained := tags[tc.tag]

		if tc.wantRemoved && remained {
			t.Errorf("pruner: expected tag '%s' to be removed", tc.tag)
		}
		if tc.wantReplacement != "" && remained {
			t.Errorf("pruner: expected tag '%s' to be reaplced, it remained", tc.tag)
		}
		vn, okay := tags[tc.wantReplacement]
		if tc.wantReplacement != "" && !okay {
			t.Errorf("pruner: expected tag '%s' to be replaced with '%s'", tc.tag, tc.wantReplacement)
		}
		if tc.wantReplacement != "" && tc.tag != vn {
			t.Errorf("pruner: value not assigned for replacement tag '%s'", tc.tag)
		}
	}
}
