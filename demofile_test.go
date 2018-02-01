package demofile_test

import (
	"testing"

	demofile "github.com/MobalyticsGG/csgo-demofile"
)

func TestDemofileOpen(t *testing.T) {
	dem, err := demofile.NewDemofile("testdata/demos/cache_9-21_mm.dem", true)
	if err != nil {
		t.Error(err)
	}

	err = dem.Start()
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkDemofileOpen1(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		dem, err := demofile.NewDemofile("testdata/demos/cache_9-21_mm.dem", false)
		if err != nil {
			b.Error(err)
		}

		err = dem.Start()
		if err != nil {
			b.Error(err)
		}
	}

}
