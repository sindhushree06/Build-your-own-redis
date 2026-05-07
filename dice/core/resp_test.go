package core

import (
	"fmt"
	"testing"
)

func TestSimpleStringDecode(t *testing.T) {
	cases := map[string]string{
		"+OW\r\n": "OW",
	}
	for k, v := range cases {
		value, _ := Decode([]byte(k))
		if v != value {
			t.Fail()
		}
	}
}

func TestErr(t *testing.T) {
	cases := map[string]string{
		"-Error message\r\n": "Error message",
	}
	for k, v := range cases {
		value, _ := Decode([]byte(k))
		if v != value {
			t.Fail()
		}
	}
}

func TestBulkStringDecode(t *testing.T) {
	cases := map[string]string{
		"$5\r\nOWOWO\r\n": "OWOWO",
		"$0\r\n\r\n":      "",
	}
	for k, v := range cases {
		value, _ := Decode([]byte(k))
		if v != value {
			t.Fail()
		}
	}
}

func TestArrayDecode(t *testing.T) {
	cases := map[string][]string{
		"*2\r\n$5\r\nOWOWO\r\n$0\r\n\r\n":                  {"OWOWO", ""},
		"*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n ": {"SET", "key", "value"},
	}
	fmt.Println(cases["key"])
	for k, v := range cases {
		value, _ := Decode([]byte(k))
		array := value.([]interface{})
		if len(array) != len(v) {
			t.Fail()
		}
		for i := range array {
			if fmt.Sprintf("%v", v[i]) != fmt.Sprintf("%v", array[i]) {
				t.Fail()
			}
		}
	}
}
