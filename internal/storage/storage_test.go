package storage_test

import (
	"../storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetOriginalUrl(t *testing.T) {
	type args struct {
		shortURL string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "DoesntURL",
			args: args{shortURL: "www.google.com"},
			want: "www.google.com",
		},
		{
			name: "Some text ",
			args: args{shortURL: "stylish exterior belies a growing darkness that has long been smoldering within, revealing that true monsters are made, not born"},
			want: "www.google.con",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			shortURL := storage.GetUrlShort(tt.args.shortURL)

			got, err := storage.GetOriginalUrl(shortURL)

			require.NoError(t, err)
			require.Equal(t, tt.args.shortURL, got)

		})
	}
}

func TestGetOriginalUrlErr(t *testing.T) {
	type args struct {
		shortURL string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Std str",
			args: args{shortURL: storage.RandString(5)},
		},
		{
			name: "big str",
			args: args{shortURL: storage.RandString(50)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := storage.GetOriginalUrl(tt.args.shortURL)

			require.Error(t, err)

		})
	}
}

func TestGetUrlShort(t *testing.T) {
	type args struct {
		originalURL string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "std url",
			args: args{originalURL: "www.google.com"},
			want: "[[:alpha:]]{6}$",
		},
		{
			name: "bad url",
			args: args{originalURL: "bla bla bla"},
			want: "[[:alpha:]]{6}$",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := storage.GetUrlShort(tt.args.originalURL)
			assert.Regexp(t, tt.want, got, "Random Str doesn't match the pattern")
		})
	}
}

func TestRandString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "std string ",
			args: args{
				n: 6,
			},
			want: "[[:alpha:]]{6}$",
		},
		{
			name: "zero sting",
			args: args{
				n: 0,
			},
			want: "[[:alpha:]]{0}$",
		},

		{
			name: "big string",
			args: args{
				n: 200,
			},
			want: "[[:alpha:]]{2ffdf00}$",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := storage.RandString(tt.args.n)

			assert.Regexp(t, tt.want, got, "Random Str doesn't match the pattern")
		})
	}
}
