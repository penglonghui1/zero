package collection

import (
	"testing"
)

func TestNewCollection(t *testing.T) {
	col := New(Int, []int{2, 3, 5, 2, 7})
	t.Log(col.Count())
	t.Log(col.Avg().MustToFloat64())
	t.Log(col.Sum().MustToInt())
	t.Log(col.Insert(4, 9).Sum().MustToInt())
	t.Log(col.Sum().MustToInt())
	t.Log(col.Contains(90))
}
