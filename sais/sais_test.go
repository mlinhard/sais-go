package sais

import (
	"fmt"
	"testing"
)

func ExampleSuffixArray() {
	T := "ABRACADABRA"
	SA, err := TextSuffixArray(T)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return
	}
	for i, val := range SA {
		fmt.Printf("SA[%2v]=%2v  %v\n", i, val, T[SA[i]:])
	}
}

func TestSuffixArray(t *testing.T) {
	T := "ABRACADABRA"
	expectedSA := []int{10, 7, 0, 3, 5, 8, 1, 4, 6, 9, 2}
	SA, err := TextSuffixArray(T)
	if err != nil {
		t.Error(err)
		return
	}
	if len(T) != len(SA) {
		t.Errorf("Length of text %v is not equal to length of the suffix array %v", len(T), len(SA))
		return
	}
	for i, val := range SA {
		if val != expectedSA[i] {
			t.Errorf("SA[%v] should be %v but was %v", i, expectedSA[i], val)
		}
	}
}
