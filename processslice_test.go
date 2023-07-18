package GoParallel

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

	ProcessSlice(testStructs, func(testStruct *TestStruct) {
		testStruct.IHaveBeenProcessed = true
	}, 5)

	for _, testStruct := range testStructs {
		if !testStruct.IHaveBeenProcessed {
			t.Errorf("TestStruct was not processed")
		}
	}
}

type TestStruct struct {
	IHaveBeenProcessed bool
}
