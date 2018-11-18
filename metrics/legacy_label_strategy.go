package metrics

import "github.com/lightwirelimited/bird_exporter/protocol"

type LegacyLabelStrategy struct {
}

func (*LegacyLabelStrategy) LabelNames() []string {
	return []string{"name"}
}

func (*LegacyLabelStrategy) LabelValues(p *protocol.Protocol) []string {
	return []string{p.Name}
}
