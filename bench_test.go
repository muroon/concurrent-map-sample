package main

import (
	"sync"
	"testing"

	"github.com/google/uuid"
	cmap "github.com/orcaman/concurrent-map/v2"
)

func BenchmarkSyncMap(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	m := sync.Map{}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			uuidv4, _ := uuid.NewRandom()
			uuid := uuidv4.String()

			m.Store(uuid, uuid)
		}
	})
}

func BenchmarkConcurrentMap(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	m := cmap.New[string]()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			uuidv4, _ := uuid.NewRandom()
			uuid := uuidv4.String()

			m.Set(uuid, uuid)
		}
	})
}
