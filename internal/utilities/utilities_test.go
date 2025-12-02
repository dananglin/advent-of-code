package utilities_test

import (
	"reflect"
	"testing"

	"codeflow.dananglin.me.uk/apollo/advent-of-code/internal/utilities"
)

func TestParseIntList(t *testing.T) {
	table := []struct {
		list string
		want []int
	}{
		{
			list: "87 2 435 89 100 1 0 16534",
			want: []int{87, 2, 435, 89, 100, 1, 0, 16534},
		},
		{
			list: "145, 834, 66, 2, 90, 123",
			want: []int{145, 834, 66, 2, 90, 123},
		},
	}

	for i := range table {
		got, err := utilities.ParseIntList(table[i].list)
		if err != nil {
			t.Fatalf("received an error after running ParseIntList(); %v", err)
		}

		if !reflect.DeepEqual(got, table[i].want) {
			t.Errorf("unexpected result received from ParseIntList(); want %v, got %v", table[i].want, got)
		}
	}
}
