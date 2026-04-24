//go:build smoke

package main

import "testing"

func TestFunctions(t *testing.T) {
	///AAA

	//Arrange
	expectedResult := 3

	//Act
	actualResult := addInts(1, 2)

	//Assert
	if actualResult != expectedResult {
		t.Errorf("addInts(1, 2) = %d; want %d", actualResult, expectedResult)
	} else {
		t.Log("Test Passed")
	}
}

type Tests struct {
	input1         int
	input2         int
	expectedOutput int
}

func TestAdd(t *testing.T) {
	tests := []Tests{
		{1, 1, 2},
		{2, 2, 4},
		{3, 3, 6},
		{0, 1, 1},
	}

	for _, test := range tests {
		actualOutput := addInts(test.input1, test.input2)
		if actualOutput != test.expectedOutput {
			t.Errorf("expected %d, got %d", test.expectedOutput, actualOutput)
		}
	}
}
