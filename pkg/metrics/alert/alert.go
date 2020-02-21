package alert

type Threshold struct {
	Min float64
	Max float64
}

func (t *Threshold) DoesMatch(value float64) bool {
	return value >= t.Min && value <= t.Max
}
