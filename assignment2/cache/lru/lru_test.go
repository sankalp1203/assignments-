package lru

import (
	"reflect"
	"testing"
)

func TestNewLRUCache(t *testing.T) {
	type args struct {
		cap int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1", args: args{cap: 1}, want: 1,
		},
		{
			name: "test2", args: args{cap: 2}, want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLRUCache(tt.args.cap).(*LRUCache); !reflect.DeepEqual(got.cap, tt.want) {
				t.Errorf("NewLRUCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRUCache_Get(t *testing.T) {
	type args struct {
		cap int
		key string
		set []data
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "test1", args: args{cap: 1, key: "india", set: []data{{"india", "delhi"}}}, want: "delhi",
		},
		{
			name: "test2", args: args{cap: 1, key: "india", set: []data{{"usa", "dc"}}}, want: "generic answer",
		},
		{
			name: "test3", args: args{cap: -1, key: "usa", set: []data{{"usa", "dc"}}}, want: "generic answer",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewLRUCache(tt.args.cap)
			for _, data := range tt.args.set {
				c.Set(data.key, data.value)
			}
			if got := c.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LRUCache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRUCache_Set(t *testing.T) {
	type args struct {
		cap int
		key string
		val any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{cap: 1, key: "india", val: "delhi"}, want: true},
		{name: "test2", args: args{cap: -1, key: "india", val: "delhi"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewLRUCache(tt.args.cap)
			if got := c.Set(tt.args.key, tt.args.val); got != tt.want {
				t.Errorf("LRUCache.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}
