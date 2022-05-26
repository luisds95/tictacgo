package tictacgo

import "testing"

func TestGetBestActionFromValues(t *testing.T) {
	tests := []struct {
		values            map[string]float64
		expectedMaxAction int
		expectedMaxValue  float64
		expectedMinAction int
		expectedMinValue  float64
	}{
		{
			values:            map[string]float64{"4": 1.0, "5": 0.0, "6": -1.0, "7": 0.0},
			expectedMaxAction: 4,
			expectedMaxValue:  1.0,
			expectedMinAction: 6,
			expectedMinValue:  -1.0,
		},
	}

	for _, test := range tests {
		maxAction, maxValue := getBestActionFromValues(test.values, true)
		minAction, minValue := getBestActionFromValues(test.values, false)
		if maxAction != test.expectedMaxAction ||
			maxValue != test.expectedMaxValue ||
			minAction != test.expectedMinAction ||
			minValue != test.expectedMinValue {
			t.Errorf(
				"Unexpected values found. Expected: [%v, %v, %v, %v], but got [%v, %v, %v, %v]",
				test.expectedMaxAction,
				test.expectedMaxValue,
				test.expectedMinAction,
				test.expectedMinValue,
				maxAction,
				maxValue,
				minAction,
				minValue,
			)
		}
	}
}
