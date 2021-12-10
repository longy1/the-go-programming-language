package main

import (
	"math"
	"testing"
)

func TestCreate(t *testing.T) {
	testIntSet := IntSet{}
	if &testIntSet == nil {
		t.Errorf("Can not create IntSet")
	}
	if len(testIntSet.words) != 0 {
		t.Errorf("Init IntSet has %d word, expected 0", len(testIntSet.words))
	}
}

func TestUpdateAndRetrieve(t *testing.T) {
	set := IntSet{}
	AddTests := struct {
		nums []int
	}{
		[]int{0, 100, math.MaxUint32},
	}
	for i := range AddTests.nums {
		set.Add(AddTests.nums[i])
		if !set.Has(AddTests.nums[i]) {
			t.Errorf("Expect Has(%d) == true", AddTests.nums[i])
		}
	}
}
