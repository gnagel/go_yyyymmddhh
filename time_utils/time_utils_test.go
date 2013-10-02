package time_utils

import "time"
import "strconv"
import "testing"
import "github.com/orfjackal/gospec/src/gospec"

func TestTimeUtilsSpecs(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in benchmark mode.")
		return
	}
	r := gospec.NewRunner()
	r.AddSpec(TimeUtilsSpecs)
	gospec.MainGoTest(r, t)
}

func TimeUtilsSpecs(c gospec.Context) {

	c.Specify("YYYY", func() {
		actual := YYYY()
		c.Expect(len(actual), gospec.Equals, 4)

		year, err := strconv.Atoi(actual)
		c.Expect(err, gospec.Equals, nil)
		c.Expect(year, gospec.Equals, time.Now().UTC().Year())
	})

	c.Specify("YYYYMM", func() {
		actual := YYYYMM()
		c.Expect(len(actual), gospec.Equals, 6)

		month, err := strconv.Atoi(actual[4:])
		c.Expect(err, gospec.Equals, nil)
		c.Expect(month, gospec.Equals, int(time.Now().UTC().Month()))
	})

	c.Specify("YYYYMMDD", func() {
		actual := YYYYMMDD()
		c.Expect(len(actual), gospec.Equals, 8)

		day, err := strconv.Atoi(actual[6:])
		c.Expect(err, gospec.Equals, nil)
		c.Expect(day, gospec.Equals, time.Now().UTC().Day())
	})

	c.Specify("YYYYMMDDHH", func() {
		actual := YYYYMMDDHH()
		c.Expect(len(actual), gospec.Equals, 10)

		hour, err := strconv.Atoi(actual[8:])
		c.Expect(err, gospec.Equals, nil)
		c.Expect(hour, gospec.Equals, time.Now().UTC().Hour())
	})

}

func Benchmark_YYYY(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YYYY()
	}
}

func Benchmark_YYYYMM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YYYYMM()
	}
}

func Benchmark_YYYYMMDD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YYYYMMDD()
	}
}

func Benchmark_YYYYMMDDHH(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YYYYMMDDHH()
	}
}
