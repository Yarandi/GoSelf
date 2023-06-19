package main

import "testing"


var tests = []struct {
	name string
	dividend float32
	divisor float32
	expected  float32
	isErr  bool
}{
	{
		name:     "valid-data",
		dividend: 100.0,
		divisor:  10.0,
		expected: 10.0,
		isErr:    false,
	},
	{
		name:     "invalid-data",
		dividend: 100.0,
		divisor:  0.0,
		expected: 0.0,
		isErr:    true,
	},
	{
		name:     "expected-5",
		dividend: 100.0,
		divisor:  20.0,
		expected: 5,
		isErr:    false,
	},
}

func TestDivision(t *testing.T){
	
	//itterate on tests struct
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		
		if tt.isErr{
			if err == nil{
				t.Error("Excpected an error but did not get one")
			}
		} else {
			if err != nil{
				t.Error("We Should not have an error but we have", err.Error() )
			}
		}

		if got != tt.expected {
			t.Errorf("Expected %f but we got %f", tt.expected, got)
		}

	}
}