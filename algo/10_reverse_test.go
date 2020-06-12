package main

import (
	"reflect"
	"testing"
)

func TestReverseLinked(t *testing.T) {
	type args struct {
		head *SingleLinkNode
	}

	s1 := NewSingleLink()
	s1.Append("a")
	s1.Append("b")
	s1.Append("c")

	s2 := NewSingleLink()
	s2.Append("c")
	s2.Append("b")
	s2.Append("a")

	tests := []struct {
		name string
		args args
		want *SingleLinkNode
	}{
		{
			name: "ReverseLinked01",
			args: args{
				head: s1.head,
			},
			want: s2.head,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseLinked(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseLinked() = %v, want %v", got, tt.want)
			}
		})
	}
}
