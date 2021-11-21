package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkTable(b *testing.B) {
	type requests struct {
		fn, ln string
	}

	benchmarks := []struct {
		name    string
		request requests
	}{
		{
			name: "Fajar",
			request: requests{
				fn: "Fajar",
				ln: "Sylvana",
			},
		},
		{
			name: "John",
			request: requests{
				fn: "John",
				ln: "Doe",
			},
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request.fn, benchmark.request.ln)
			}
		})
	}
}

/* func BenchmarkSub(b *testing.B) {
	b.Run("Fajar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Fajar", "Sylvana")
		}
	})
	b.Run("John", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("John", "Doe")
		}
	})
}*/

func TestHelloWorld(t *testing.T) {
	type args struct {
		firstName, lastName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Fajar Sylvana",
			args: args{
				firstName: "Fajar",
				lastName:  "Sylvana",
			},
			want: "Hello Fajar Sylvana",
		},
		{
			name: "Ninja Naruto",
			args: args{
				firstName: "Ninja",
				lastName:  "Naruto",
			},
			want: "Hello Ninja Naruto",
		},
		{
			name: "Bruce Wayne",
			args: args{
				firstName: "Bruce",
				lastName:  "Wayne",
			},
			want: "Hello Bruce Wayne",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HelloWorld(tt.args.firstName, tt.args.lastName)
			assert.Equal(t, tt.want, result)
		})
	}
}
