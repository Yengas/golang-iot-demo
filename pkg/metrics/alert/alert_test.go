package alert

import "testing"

func TestThreshold_DoesMatch(t1 *testing.T) {
	type fields struct {
		Min float64
		Max float64
	}
	type args struct {
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"does match if min", fields{5, 10}, args{5}, true},
		{"does match if max", fields{5, 10}, args{10}, true},
		{"does match if between", fields{5, 10}, args{7}, true},
		{"does not match if lesser than min", fields{5, 10}, args{4}, false},
		{"does not match if bigger than max", fields{5, 10}, args{11}, false},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Threshold{
				Min: tt.fields.Min,
				Max: tt.fields.Max,
			}
			if got := t.DoesMatch(tt.args.value); got != tt.want {
				t1.Errorf("DoesMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
