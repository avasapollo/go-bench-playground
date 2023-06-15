package compare_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/faceit/go-bench-playground/compare"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func buildRandomSlice(howMany, howLong int) []string {
	sl := make([]string, howMany)
	for i := 0; i < howMany; i++ {
		sl[i] = randStringBytes(howLong)
	}
	return sl
}

func TestStringSliceByDeepEqual(t *testing.T) {
	t.Parallel()

	howLong := 8
	howMany := 10
	tests := []struct {
		name string
		init func() (sl1, sl2 []string)
		want bool
	}{
		{
			name: "equal",
			init: func() (sl1, sl2 []string) {
				sl1 = buildRandomSlice(howMany, howLong)
				sl2 = make([]string, len(sl1))
				copy(sl2, sl1)
				return sl1, sl2
			},
			want: true,
		},
		{
			name: "not equal",
			init: func() (sl1, sl2 []string) {
				sl1 = buildRandomSlice(howMany, howLong)
				sl2 = make([]string, len(sl1))
				copy(sl2, sl1)
				sl1 = append(sl1, randStringBytes(howLong))
				return sl1, sl2
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			sl1, sl2 := tt.init()
			got := compare.StringSliceByDeepEqual(sl1, sl2)
			require.Equal(t, tt.want, got)
		})
	}
}
