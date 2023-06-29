package compare_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"fmt"

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

func TestStringSliceCustom1(t *testing.T) {
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
			got := compare.StringSliceCustom1(sl1, sl2)
			require.Equal(t, tt.want, got)
		})
	}
}

func BenchmarkCompare_with2M(b *testing.B) {
	sl1 := buildRandomSlice(2e6, 8)
	sl2 := make([]string, len(sl1))
	copy(sl2, sl1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = compare.StringSlice(sl1, sl2)
	}
}

func BenchmarkCompare_with5(b *testing.B) {
	sl1 := buildRandomSlice(2e6, 5)
	sl2 := make([]string, len(sl1))
	copy(sl2, sl1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = compare.StringSlice(sl1, sl2)
	}
}

func BenchmarkSCompare(b *testing.B) {
	for _, tcase := range []struct {
		howMany int
		howLong int
	}{
		{
			howMany: 3,
			howLong: 8,
		},
		{
			howMany: 5,
			howLong: 8,
		},
		{
			howMany: 10,
			howLong: 8,
		},
		{
			howMany: 20,
			howLong: 8,
		},
		{
			howMany: 100,
			howLong: 8,
		},
	} {
		b.Run(fmt.Sprintf("lines-%d", tcase.howMany), func(b *testing.B) {
			sl1 := buildRandomSlice(tcase.howMany, tcase.howLong)
			sl2 := make([]string, len(sl1))
			copy(sl2, sl1)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = compare.StringSlice(sl1, sl2)
			}
		})
	}
}
