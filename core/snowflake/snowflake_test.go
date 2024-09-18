package sonyflake

import (
	"fmt"
	"testing"
	"time"
)

func TestSonyflake_NextID(t *testing.T) {
	var now, _ = time.Parse("2006-01-02", "2009-02-25")
	snow := NewSonyflake(Settings{
		StartTime: now,
	})

	generateID(t, snow)
}

func BenchmarkSonyflake_NextID(b *testing.B) {
	var now, _ = time.Parse("2006-01-02", "2009-02-25")
	snow := NewSonyflake(Settings{
		StartTime: now,
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id, err := snow.NextID()
		if err != nil {
			b.Error(err)
			return
		}
		b.Log(id)
	}
}

func TestDecompose(t *testing.T) {
	var now, _ = time.Parse("2006-01-02", "2009-02-25")
	snow := NewSonyflake(Settings{
		StartTime: now,
	})

	t.Log(snow.startTime)
	id, err := snow.NextID()
	if err != nil {
		t.Error(err)
		return
	}
	m := Decompose(id)
	t.Log(m)
	// 8位机器码

	fmt.Println(m)
}

func generateID(t *testing.T, snow *Sonyflake) {
	for i := 0; i < 100000; i++ {
		id, err := snow.NextID()
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(id)
	}
}
