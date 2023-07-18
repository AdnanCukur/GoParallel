package goparallel

import "testing"

func TestProcessSlice_EnsureAllItemsGetProcessed(t *testing.T) {
	testStructs := []*TestStruct{
		{IHaveBeenProcessed: false},
		{IHaveBeenProcessed: false},
		{IHaveBeenProcessed: false},
		{IHaveBeenProcessed: true},
		{IHaveBeenProcessed: false},
		{IHaveBeenProcessed: true},
		{IHaveBeenProcessed: false},
		{IHaveBeenProcessed: false},
		{IHaveBeenProcessed: false},
	}
	maxParallelism := 5
	ProcessSlice(testStructs, func(testStruct *TestStruct) {
		testStruct.IHaveBeenProcessed = true
	}, maxParallelism)

	for _, testStruct := range testStructs {
		if !testStruct.IHaveBeenProcessed {
			t.Errorf("TestStruct was not processed")
		}
	}
}

type TestStruct struct {
	IHaveBeenProcessed bool
}
