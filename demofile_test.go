package demofile_test

import (
	"testing"

	demofile "github.com/MobalyticsGG/csgo-demofile"
)

func TestDemofileOpen(t *testing.T) {
	dem, err := demofile.NewDemofile("testdata/demos/cache_9-21_mm.dem")
	if err != nil {
		t.Error(err)
	}

	err = dem.Start()
	if err != nil {
		t.Error(err)
	}

}
