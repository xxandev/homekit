package homekit

import "testing"

func TestConvertBool(t *testing.T) {
	var correctValues, errorValues []interface{}
	correctValues = append(correctValues, true, false)
	errorValues = append(errorValues, "val", 2.32, 0.1, nil, "true", "false", "123", 123, 1, 0)
	if toBool(correctValues[0], false) == false {
		t.Fatalf("error test - correctValues acchc.toBool(false, %v[%T])\n", correctValues[0], correctValues[0])
	}
	if toBool(correctValues[1], true) == true {
		t.Fatalf("error test - correctValues acchc.toBool(true, %v[%T])\n", correctValues[1], correctValues[1])
	}
	for i := range errorValues {
		if toBool(errorValues[i], true) != true {
			t.Fatal("error test - errorValues acchc.toBool(-127,", errorValues[i], ")")
		}
	}
}

func TestConvertInt(t *testing.T) {
	var correctValues, errorValues []interface{}
	correctValues = append(correctValues, 11, 12, 23, 4652, 11, 254, 1, 0, 125)
	errorValues = append(errorValues, "val", 2.32, 1.0, 0.1, 0.0, nil, true, false, "123")
	for i := range correctValues {
		if toInt(correctValues[i], -127) == -127 {
			t.Fatalf("error test - correctValues acchc.toInt(-127, %v[%T])\n", correctValues[i], correctValues[i])
		}
		if toInt32(correctValues[i], -127) == -127 {
			t.Fatalf("error test - correctValues acchc.toInt32(-127, %v[%T])\n", correctValues[i], correctValues[i])
		}
		if toInt64(correctValues[i], -127) == -127 {
			t.Fatalf("error test - correctValues acchc.toInt64(-127, %v[%T])\n", correctValues[i], correctValues[i])
		}
	}
	for i := range errorValues {
		if toInt(errorValues[i], -127) != -127 {
			t.Fatalf("error test - errorValues acchc.toInt(-127, %v[%T])\n", errorValues[i], errorValues[i])
		}
		if toInt32(errorValues[i], -127) != -127 {
			t.Fatalf("error test - errorValues acchc.toInt32(-127, %v[%T])\n", errorValues[i], errorValues[i])
		}
		if toInt64(errorValues[i], -127) != -127 {
			t.Fatalf("error test - errorValues acchc.toInt64(-127, %v[%T])\n", errorValues[i], errorValues[i])
		}
	}
}

func TestConvertFloat(t *testing.T) {
	var correctValues, errorValues []interface{}
	correctValues = append(correctValues, 0.0, 0.1, 1.0, 123.321, 1.001, 0, 1, 123, 6350453)
	errorValues = append(errorValues, "val", nil, true, false, "123", "123.123")
	for i := range correctValues {
		if toFloat32(correctValues[i], -0.127) == -0.127 {
			t.Fatalf("error test - correctValues acchc.convertFloat32(-0.127, %v[%T])\n", correctValues[i], correctValues[i])
		}
		if toFloat64(correctValues[i], -0.127) == -0.127 {
			t.Fatalf("test Fatal - correctValues acchc.convertFloat64(-0.127, %v[%T])\n", correctValues[i], correctValues[i])
		}
	}
	for i := range errorValues {
		if toFloat32(errorValues[i], -0.127) != -0.127 {
			t.Fatalf("error test - errorValues acchc.convertFloat32(-0.127, %v[%T])\n", errorValues[i], errorValues[i])
		}
		if toFloat64(errorValues[i], -0.127) != -0.127 {
			t.Fatalf("test Fatal - errorValues acchc.convertFloat64(-0.127, %v[%T])\n", errorValues[i], errorValues[i])
		}
	}
}
