package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_NormalItem(t *testing.T) {
	var items = []*gildedrose.Item{
		{"normal item", 10, 10},
	}

	expected := gildedrose.Item{"normal item", 9, 9}

	gildedrose.UpdateQuality(items)

	if *items[0] != expected {
		t.Errorf("Expected %v but got %v ", expected, *items[0])
	}
}

func Test_NormalItems(t *testing.T) {

	tbl := []struct {
		name     string
		input    gildedrose.Item
		expected gildedrose.Item
	}{
		{
			"normal item: q > 0, s > 0",
			gildedrose.Item{"normal item", 10, 10},
			gildedrose.Item{"normal item", 9, 9},
		},
		{
			"normal item: q = 0, s > 0",
			gildedrose.Item{"normal item", 10, 0},
			gildedrose.Item{"normal item", 9, 0},
		},
		{
			"normal item: q > 0, s < 0",
			gildedrose.Item{"normal item", -1, 10},
			gildedrose.Item{"normal item", -2, 8},
		},
		{
			"normal item: q > 0, s = 0",
			gildedrose.Item{"normal item", 0, 10},
			gildedrose.Item{"normal item", -1, 8},
		},
	}

	for _, tc := range tbl {
		t.Run(tc.name, func(t *testing.T) {

			items := []*gildedrose.Item{
				&tc.input,
			}

			gildedrose.UpdateQuality(items)

			if *items[0] != tc.expected {
				t.Errorf("Expected %v but got %v ", tc.expected, *items[0])
			}
		})
	}
}
