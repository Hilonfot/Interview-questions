package main


import (
	"reflect"
	"testing"
)

func TestSortByBubble(t *testing.T) {
	type args struct {
		link *SingleLink
	}
	s := NewSingleLink()
	s.Append(2)
	s.Append(1)
	s.Append(3)
	s1 := NewSingleLink()
	s1.Append(3)
	s1.Append(2)
	s1.Append(1)
	tests := []struct {
		name string
		args args
		want *SingleLink
	}{
		{
			name: "data1",
			args: args{
				link: s,
			},
			want: s1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortByBubble(tt.args.link); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortByBubble() = %v, want %v", got, tt.want)
			}
		})
	}
}
