package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func testInitializeDataBase() *db {
	d := &db{
		state: map[string]string{
			"foo":  "bar",
			"span": "eggs",
			"go":   "lang",
			"test": "test",
		},
	}
	return d
}

func Test_db_Find(t *testing.T) {
	// setUp block
	d := testInitializeDataBase()

	type args struct {
		key string
	}

	type want struct {
		value string
		ok    bool
	}

	tests := []struct {
		name        string
		description string
		args        args
		want        want
	}{
		{
			name:        "Test db.Find method #1",
			description: "Try to find value that exists",
			args: args{
				key: "foo",
			},
			want: want{
				value: "bar",
				ok:    true,
			},
		},
		{
			name:        "Test db.Find method #2",
			description: "Try to find value that exists",
			args: args{
				key: "span",
			},
			want: want{
				value: "eggs",
				ok:    true,
			},
		},
		{
			name:        "Test db.Find method #3",
			description: "Try to find value that exists",
			args: args{
				key: "go",
			},
			want: want{
				value: "lang",
				ok:    true,
			},
		},
		{
			name:        "Test db.Find method #4",
			description: "Try to find value that exists",
			args: args{
				key: "test",
			},
			want: want{
				value: "test",
				ok:    true,
			},
		},
		{
			name:        "Test db.Find method #5",
			description: "Try to find value that not exists",
			args: args{
				key: "123",
			},
			want: want{
				value: "",
				ok:    false,
			},
		},
		{
			name:        "Test db.Find method #6",
			description: "Try to find value that not exists",
			args: args{
				key: "asdsad",
			},
			want: want{
				value: "",
				ok:    false,
			},
		},
		{
			name:        "Test db.Find method #7",
			description: "Try to find value with empty key",
			args: args{
				key: "",
			},
			want: want{
				value: "",
				ok:    false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			value, ok := d.Find(tt.args.key)
			require.Equal(t, value, tt.want.value, "Value undefined")
			require.Equal(t, ok, tt.want.ok)
		})
	}
}
