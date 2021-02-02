package homekit

import "testing"

func TestConvertBool(t *testing.T) {
	var correctValues, errorValues []interface{}
	correctValues = append(correctValues, true, false)
	errorValues = append(errorValues, "val", 2.32, 0.1, nil, "true", "false", "123", 123, 1, 0)
	if argToBool(correctValues[0], false) == false {
		t.Fatalf("error test - correctValues acchc.argToBool(false, %v[%T])\n", correctValues[0], correctValues[0])
	}
	if argToBool(correctValues[1], true) == true {
		t.Fatalf("error test - correctValues acchc.argToBool(true, %v[%T])\n", correctValues[1], correctValues[1])
	}
	for i := range errorValues {
		if argToBool(errorValues[i], true) != true {
			t.Fatal("error test - errorValues acchc.argToBool(-127,", errorValues[i], ")")
		}
	}
}

func TestConvertInt(t *testing.T) {
	var correctValues, errorValues []interface{}
	correctValues = append(correctValues, 11, 12, 23, 4652, 11, 254, 1, 0, 125)
	errorValues = append(errorValues, "val", 2.32, 1.0, 0.1, 0.0, nil, true, false, "123")
	for i := range correctValues {
		if argToInt(correctValues[i], -127) == -127 {
			t.Fatalf("error test - correctValues acchc.argToInt(-127, %v[%T])\n", correctValues[i], correctValues[i])
		}
		if argToInt32(correctValues[i], -127) == -127 {
			t.Fatalf("error test - correctValues acchc.argToInt32(-127, %v[%T])\n", correctValues[i], correctValues[i])
		}
		if argToInt64(correctValues[i], -127) == -127 {
			t.Fatalf("error test - correctValues acchc.argToInt64(-127, %v[%T])\n", correctValues[i], correctValues[i])
		}
	}
	for i := range errorValues {
		if argToInt(errorValues[i], -127) != -127 {
			t.Fatalf("error test - errorValues acchc.argToInt(-127, %v[%T])\n", errorValues[i], errorValues[i])
		}
		if argToInt32(errorValues[i], -127) != -127 {
			t.Fatalf("error test - errorValues acchc.argToInt32(-127, %v[%T])\n", errorValues[i], errorValues[i])
		}
		if argToInt64(errorValues[i], -127) != -127 {
			t.Fatalf("error test - errorValues acchc.argToInt64(-127, %v[%T])\n", errorValues[i], errorValues[i])
		}
	}
}

func TestConvertFloat(t *testing.T) {
	var correctValues, errorValues []interface{}
	correctValues = append(correctValues, 0.0, 0.1, 1.0, 123.321, 1.001, 0, 1, 123, 6350453)
	errorValues = append(errorValues, "val", nil, true, false, "123", "123.123")
	for i := range correctValues {
		if argToFloat32(correctValues[i], -0.127) == -0.127 {
			t.Fatalf("error test - correctValues acchc.convertFloat32(-0.127, %v[%T])\n", correctValues[i], correctValues[i])
		}
		if argToFloat64(correctValues[i], -0.127) == -0.127 {
			t.Fatalf("test Fatal - correctValues acchc.convertFloat64(-0.127, %v[%T])\n", correctValues[i], correctValues[i])
		}
	}
	for i := range errorValues {
		if argToFloat32(errorValues[i], -0.127) != -0.127 {
			t.Fatalf("error test - errorValues acchc.convertFloat32(-0.127, %v[%T])\n", errorValues[i], errorValues[i])
		}
		if argToFloat64(errorValues[i], -0.127) != -0.127 {
			t.Fatalf("test Fatal - errorValues acchc.convertFloat64(-0.127, %v[%T])\n", errorValues[i], errorValues[i])
		}
	}
}
