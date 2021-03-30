package level

import "testing"

func TestMergeNibbles(t *testing.T) {
	a := byte(0x10)
	b := byte(0x10)
	result := mergeNibbles(a, b)

	if result != 0x11 {
		t.Logf("result should be 0x11 but was %x", result)
		t.Fail()
	}

	a = byte(0x00)
	b = byte(0x10)
	result = mergeNibbles(a, b)

	if result != 0x01 {
		t.Logf("result should be 0x01 but was %x", result)
		t.Fail()
	}

	a = byte(0x00)
	b = byte(0x00)
	result = mergeNibbles(a, b)

	if result != 0x00 {
		t.Logf("result should be 0x00 but was %x", result)
		t.Fail()
	}

	a = byte(0x30)
	b = byte(0x10)
	result = mergeNibbles(a, b)

	if result != 0x31 {
		t.Logf("result should be 0x31 but was %x", result)
		t.Fail()
	}
}
