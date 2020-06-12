package main

import (
	"testing"
)

func TestLRUCache_Get(t *testing.T) {
	c := NewLRUCache(5)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	t.Log(c)
	type args struct {
		key int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "数据一",
			args: args{1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := c.Get(tt.args.key); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
