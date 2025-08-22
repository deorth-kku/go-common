package common

import (
	"os"
	"testing"
)

func SkipLongTest(t *testing.T) {
	if _, ok := t.Deadline(); ok {
		t.SkipNow()
	}
}

func CreateTestConf(t *testing.T, data string) string {
	f, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatal(err)
		return ""
	}
	_, err = f.WriteString(data)
	if err != nil {
		t.Fatal(err)
		return ""
	}
	err = f.Close()
	if err != nil {
		t.Fatal(err)
		return ""
	}
	return f.Name()
}
