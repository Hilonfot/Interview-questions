package main

import (
	"testing"
)

func TestSingleLink_String(t *testing.T) {
	type fields struct {
		head *SingleLinkNode
		tail *SingleLinkNode
		len  int
	}
	s := NewSingleLink()
	s.Append("a")
	s.Append("b")
	s.Append("c")
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "data1",
			fields: fields{
				head: s.head,
				tail: s.tail,
				len:  s.len,
			},
			want: "链表结果:a=>b=>c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleLink{
				head: tt.fields.head,
				tail: tt.fields.tail,
				len:  tt.fields.len,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
