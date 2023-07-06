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
	name := "normal item"

	tbl := []struct {
		name     string
		input    gildedrose.Item
		expected gildedrose.Item
	}{
		{
			"normal item: q > 0, s > 0",
			gildedrose.Item{name, 10, 10},
			gildedrose.Item{name, 9, 9},
		},
		{
			"normal item: q = 0, s > 0",
			gildedrose.Item{name, 10, 0},
			gildedrose.Item{name, 9, 0},
		},
		{
			"normal item: q > 0, s < 0",
			gildedrose.Item{name, -1, 10},
			gildedrose.Item{name, -2, 8},
		},
		{
			"normal item: q > 0, s = 0",
			gildedrose.Item{name, 0, 10},
			gildedrose.Item{name, -1, 8},
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

func Test_AgedBrie(t *testing.T) {
	name := "Aged Brie"

	tbl := []struct {
		name     string
		input    gildedrose.Item
		expected gildedrose.Item
	}{
		{
			"aged brie: q > 0, s > 0",
			gildedrose.Item{name, 10, 10},
			gildedrose.Item{name, 9, 11},
		},
		{
			"aged brie: q = 0, s > 0",
			gildedrose.Item{name, 10, 0},
			gildedrose.Item{name, 9, 1},
		},
		{
			"aged brie: q > 0, s < 0",
			gildedrose.Item{name, -1, 10},
			gildedrose.Item{name, -2, 12},
		},
		{
			"aged brie: q > 0, s = 0",
			gildedrose.Item{name, 0, 10},
			gildedrose.Item{name, -1, 12},
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

func Test_BackstagePass(t *testing.T) {
	name := "Backstage passes to a TAFKAL80ETC concert"

	tbl := []struct {
		name     string
		input    gildedrose.Item
		expected gildedrose.Item
	}{
		{
			"backstage pass: q > 0, s > 0",
			gildedrose.Item{name, 10, 10},
			gildedrose.Item{name, 9, 12},
		},
		{
			"backstage pass: q = 0, s > 0",
			gildedrose.Item{name, 10, 0},
			gildedrose.Item{name, 9, 2},
		},
		{
			"backstage pass: q > 0, s < 0",
			gildedrose.Item{name, -1, 10},
			gildedrose.Item{name, -2, 0},
		},
		{
			"backstage pass: q > 0, s = 0",
			gildedrose.Item{name, 0, 10},
			gildedrose.Item{name, -1, 0},
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

func Test_Sulfuras(t *testing.T) {
	name := "Sulfuras, Hand of Ragnaros"

	tbl := []struct {
		name     string
		input    gildedrose.Item
		expected gildedrose.Item
	}{
		{
			"sulfuras",
			gildedrose.Item{name, 0, 80},
			gildedrose.Item{name, 0, 80},
		},
		{
			"sulfuras s<0",
			gildedrose.Item{name, -1, 80},
			gildedrose.Item{name, -1, 80},
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
