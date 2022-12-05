package homekit

import "testing"

func TestToBool(t *testing.T) {
	type tester struct {
		val interface{}
		def bool
		res bool
	}
	testik := []tester{
		{val: "val", def: true, res: true},
		{val: "val", def: false, res: false},
		{val: "false", def: true, res: true},
		{val: "false", def: false, res: false},
		{val: "true", def: true, res: true},
		{val: "true", def: false, res: false},
		{val: 1, def: true, res: true},
		{val: 1, def: false, res: false},
		{val: 0, def: true, res: true},
		{val: 0, def: false, res: false},
		{val: 0.1, def: true, res: true},
		{val: 0.1, def: false, res: false},
		{val: -1, def: true, res: true},
		{val: -1, def: false, res: false},
		{val: -0.1, def: true, res: true},
		{val: -0.1, def: false, res: false},
		{val: nil, def: true, res: true},
		{val: nil, def: false, res: false},
		{val: 0x1, def: true, res: true},
		{val: 0x1, def: false, res: false},
		{val: 0x0, def: true, res: true},
		{val: 0x0, def: false, res: false},
		{val: true, def: false, res: true},
		{val: false, def: true, res: false},
	}
	for i := range testik {
		if tob(testik[i].val, testik[i].def) != testik[i].res {
			t.Fatalf("error test func homekit.tob(%[1]v[%[1]T], %[2]v[%[2]T]) != %[3]v[%[3]T]\n", testik[i].val, testik[i].def, testik[i].res)
		}
	}
}

func TestToInt(t *testing.T) {
	type tester struct {
		val interface{}
		def int
		res int
	}
	testik := []tester{
		{val: "val", def: 256, res: 256},
		{val: "32", def: 256, res: 256},
		{val: "123.8", def: 256, res: 256},
		{val: true, def: 256, res: 256},
		{val: false, def: 256, res: 256},
		{val: 1, def: 256, res: 1},
		{val: 0, def: 256, res: 0},
		{val: 0.1, def: 256, res: 256},
		{val: -1, def: 256, res: -1},
		{val: -0.1, def: 256, res: 256},
		{val: 128.64, def: 256, res: 256},
		{val: -128.64, def: 256, res: 256},
		{val: nil, def: 256, res: 256},
		{val: 0x1, def: 256, res: 1},
		{val: 0x0, def: 256, res: 0},
	}
	for i := range testik {
		if toi(testik[i].val, testik[i].def) != testik[i].res {
			t.Fatalf("error test func homekit.toi(%[1]v[%[1]T], %[2]v[%[2]T]) != %[3]v[%[3]T]\n", testik[i].val, testik[i].def, testik[i].res)
		}
	}
}

func TestToFloat64(t *testing.T) {
	type tester struct {
		val interface{}
		def float64
		res float64
	}
	testik := []tester{
		{val: "val", def: 256.00, res: 256.00},
		{val: "32", def: 256.00, res: 256.00},
		{val: "123.8", def: 256.00, res: 256.00},
		{val: true, def: 256.00, res: 256.00},
		{val: false, def: 256.00, res: 256.00},
		{val: 1, def: 256, res: 1.00},
		{val: 0, def: 256, res: 0.00},
		{val: 0.1, def: 256.00, res: 0.1},
		{val: -1, def: 256.00, res: -1.00},
		{val: -0.1, def: 256.00, res: -0.1},
		{val: 128.64, def: 256.00, res: 128.64},
		{val: -128.64, def: 256.00, res: -128.64},
		{val: nil, def: 256.00, res: 256.00},
		{val: 0x1, def: 256.00, res: 1.00},
		{val: 0x0, def: 256.00, res: 0.00},
	}
	for i := range testik {
		if tof64(testik[i].val, testik[i].def) != testik[i].res {
			t.Fatalf("error test func homekit.tof64(%[1]v[%[1]T], %[2]v[%[2]T]) != %[3]v[%[3]T]\n", testik[i].val, testik[i].def, testik[i].res)
		}
	}
}
