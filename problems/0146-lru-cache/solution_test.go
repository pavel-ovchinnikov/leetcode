package lrucache

import (
	"testing"
)

func TestLRUCache_Get(t *testing.T) {
	capacity := 2
	tests := []struct {
		name string
		key  int
		want int
	}{
		{name: "Test Case 1", key: 1, want: -1},
		{name: "Test Case 2", key: 2, want: -1},
		{name: "Test Case 3", key: 3, want: 3},
		{name: "Test Case 4", key: 4, want: 4},
	}
	cache := Constructor(capacity)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cache.Get(tt.key)
			if got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
