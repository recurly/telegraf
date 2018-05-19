package recurly

import (
	"strings"

	"github.com/influxdata/telegraf"
)

// keepTags are tags that we are kept. Otherwise, they are pruned off.
// This could use regex's to make things easy, however, they are slower in Go.
// As a general rule, regexs are avoided unless it makes sense.
var keepTags = map[string]string{
	"begin_time":    "begin_time",
	"begintime":     "begin_time",
	"event":         "event",
	"eventtype":     "event",
	"method":        "method",
	"rendertime":    "render_time",
	"request_time":  "request_time",
	"response":      "response",
	"response_time": "response_time",
	"responsetime":  "response_time",
	"source_type":   "source_type",
	"sourcetype":    "source_type",
	"status":        "status",
	"status_code":   "status",
	"wait_time":     "wait_time",
}

// prune removes tags that are not needed.
func (r *Recurly) prune(in telegraf.Metric) telegraf.Metric {
	for _, tag := range in.Tags() {
		value, _ := in.GetTag(tag)
		l := strings.ToLower(tag)
		l = strings.Replace(l, "-", "_", -1)
		in.RemoveTag(tag)

		kAs, keep := keepTags[l]
		if r.KeepHostNames {
			if strings.HasPrefix(l, "host") || strings.HasSuffix(l, "host") {
				keep = true
				kAs = "host"
			} else if strings.HasPrefix(l, "server") || strings.HasSuffix(l, "server") {
				keep = true
				kAs = "server"
			}
		}

		if !keep {
			continue
		}

		in.AddTag(kAs, value)
	}
	return in
}
