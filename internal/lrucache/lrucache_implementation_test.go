package lrucache

import (
	"testing"
)

func TestLRUEntity_Add(t *testing.T) {
	type args struct {
		key   string
		value string
	}

	l := NewLRUCache(2)

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Add 1",
			args: args{key: "key1", value: "val1"},
			want: true,
		},
		{
			name: "Add 2",
			args: args{key: "key2", value: "val2"},
			want: true,
		},
		{
			name: "Add 3",
			args: args{key: "key3", value: "val3"},
			want: true,
		},
		{
			name: "Add 4",
			args: args{key: "key3", value: "val345"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := l.Add(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("LRUEntity.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRUEntity_Get(t *testing.T) {
	type args struct {
		key string
	}

	l := NewLRUCache(2)
	_ = l.Add("key1", "val1")
	_ = l.Add("key2", "val2")

	tests := []struct {
		name      string
		args      args
		wantValue string
		wantOk    bool
	}{
		{
			name:      "Get 1",
			args:      args{key: "key1"},
			wantValue: "val1",
			wantOk:    true,
		},
		{
			name:      "Get 2",
			args:      args{key: "key2"},
			wantValue: "val2",
			wantOk:    true,
		},
		{
			name:      "Get 3",
			args:      args{key: "key3"},
			wantValue: "",
			wantOk:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotOk := l.Get(tt.args.key)
			if gotValue != tt.wantValue {
				t.Errorf("LRUEntity.Get() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotOk != tt.wantOk {
				t.Errorf("LRUEntity.Get() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}

	_ = l.Add("key2", "val234")
	t.Run("Get with updated val", func(t *testing.T) {
		gotValue, gotOk := l.Get("key2")
		if gotValue != "val234" {
			t.Errorf("LRUEntity.Get() gotValue = %v, want %v", gotValue, "val234")
		}
		if !gotOk {
			t.Errorf("LRUEntity.Get() gotOk = %v, want %v", gotOk, true)
		}
	})

}

func TestLRUEntity_Remove(t *testing.T) {
	type args struct {
		key string
	}

	l := NewLRUCache(2)
	_ = l.Add("key1", "val1")
	_ = l.Add("key2", "val2")

	tests := []struct {
		name   string
		args   args
		wantOk bool
	}{
		{
			name:   "Remove 1",
			args:   args{key: "key1"},
			wantOk: true,
		},
		{
			name:   "Remove 2",
			args:   args{key: "key1"},
			wantOk: false,
		},
		{
			name:   "Remove 3",
			args:   args{key: "key2"},
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOk := l.Remove(tt.args.key); gotOk != tt.wantOk {
				t.Errorf("LRUEntity.Remove() = %v, want %v", gotOk, tt.wantOk)

			} else {
				if _, gotOk := l.Get(tt.args.key); gotOk {
					t.Errorf("LRUEntity.Remove(): element with key %s was not deleted", tt.args.key)
				}
			}

		})
	}
}
