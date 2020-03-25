package lru

import (
	"reflect"
	"testing"
)

func TestDefaultLRU_Set(t *testing.T) {
	tests := map[string]struct {
		keys   []string
		vals   []int
		maxCap int
		res    []int
	}{
		"normal": {
			keys:   []string{"one", "two", "three"},
			vals:   []int{1, 2, 3},
			maxCap: 4,
			res:    []int{1, 2, 3},
		},
		"overflow": {
			keys:   []string{"one", "two", "three"},
			vals:   []int{1, 2, 3},
			maxCap: 1,
			res:    []int{3},
		},
		"emtpy": {
			keys:   []string{"one", "two", "three"},
			vals:   []int{1, 2, 3},
			maxCap: 0,
			res:    []int{},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			lru := NewLRU()
			lru.SetMaxCap(tc.maxCap)
			var got = []int{}
			for i, v := range tc.keys {
				lru.Set(v, tc.vals[i])
			}

			for _, v := range tc.keys {
				if lru.Get(v) != nil {
					got = append(got, lru.Get(v).(int))
				}
			}

			if !reflect.DeepEqual(got, tc.res) {
				t.Fatalf("want: %v, got: %v", tc.res, got)
			}
		})
	}
}

func TestDefaultLRU_Clear(t *testing.T) {
	tests := map[string]struct {
		keys   []string
		vals   []int
		maxCap int
		res    []int
	}{
		"normal": {
			keys:   []string{"one", "two", "three"},
			vals:   []int{1, 2, 3},
			maxCap: 4,
			res:    []int{},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			lru := NewLRU()
			lru.SetMaxCap(tc.maxCap)
			var got = []int{}
			for i, v := range tc.keys {
				lru.Set(v, tc.vals[i])
			}

			lru.Clear()
			for _, v := range tc.keys {
				if lru.Get(v) != nil {
					got = append(got, lru.Get(v).(int))
				}
			}

			if !reflect.DeepEqual(got, tc.res) {
				t.Fatalf("want: %v, got: %v", tc.res, got)
			}
		})
	}
}

func TestDefaultLRU_Remove(t *testing.T) {
	tests := map[string]struct {
		keys   []string
		vals   []int
		remove []string
		maxCap int
		res    []int
	}{
		"normal": {
			keys:   []string{"one", "two", "three"},
			vals:   []int{1, 2, 3},
			remove: []string{"two", "three"},
			maxCap: 4,
			res:    []int{1},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			lru := NewLRU()
			lru.SetMaxCap(tc.maxCap)
			var got = []int{}
			for i, v := range tc.keys {
				lru.Set(v, tc.vals[i])
			}

			for _, v := range tc.remove {
				lru.Remove(v)
			}

			for _, v := range tc.keys {
				if lru.Get(v) != nil {
					got = append(got, lru.Get(v).(int))
				}
			}

			if !reflect.DeepEqual(got, tc.res) {
				t.Fatalf("want: %v, got: %v", tc.res, got)
			}
		})
	}
}
