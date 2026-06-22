package chttp

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	ctest "github.com/deorth-kku/go-common/test"
)

func TestHttp(t *testing.T) {
	server := NewServer()
	dir := t.TempDir()
	time.AfterFunc(1*time.Second, func() {
		server.Close()
	})
	err := server.ListenAndServe(filepath.Join(dir, "123.sock") + ",0666")
	ctest.NoError(t, err)
}

func TestAbstract(t *testing.T) {
	server := NewServer()
	time.AfterFunc(1*time.Second, func() {
		server.Close()
	})
	err := server.ListenAndServe("@123.sock")
	ctest.NoError(t, err)
}

func TestParseMode(t *testing.T) {
	str, m, _, err := FileWithMode("te\\,st,0666")
	if err != nil {
		t.Error(err)
		return
	}
	if m != 0666 {
		t.Error("wrong")
	}
	fmt.Println(str)
}

func TestCheckDirPerm(t *testing.T) {
	dirname := filepath.Join(t.TempDir(), "123")
	err := CheckDirWritePermission(dirname)
	if err == nil {
		t.Error("not exist but no error")
	} else {
		fmt.Println("expected no exist error:", err)
	}
	err = os.Mkdir(dirname, 0000)
	if err != nil {
		t.Error(err)
		return
	}
	defer os.Remove(dirname)
	err = CheckDirWritePermission(dirname)
	if err == nil {
		t.Error("not permissoned but no error")
	} else {
		fmt.Println("expected no permission error:", err)
	}
	err = os.Chmod(dirname, 0755)
	if err != nil {
		t.Error(err)
		return
	}
	err = CheckDirWritePermission(dirname)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckFilePerm(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "123", "123")
	err := CheckFileWritePermission(filename)
	if err == nil {
		t.Error("parent not exist but no error")
	} else {
		t.Log("expected parent not exist error:", err)
	}

	dirname := filepath.Dir(filename)
	err = os.Mkdir(dirname, 0000)
	if err != nil {
		t.Error(err)
		return
	}
	err = CheckFileWritePermission(filename)
	if err == nil {
		t.Error("parent no permisson but no error")
	} else {
		t.Log("expected parent no permisson error:", err)
	}
	ctest.NoError(t, os.Chmod(dirname, 0755))

	f, err := os.Create(filename)
	if err != nil {
		t.Error(err)
		return
	}
	f.Close()

	err = CheckFileWritePermission(filename)
	if err != nil {
		t.Error(err)
		return
	}

	err = os.Chmod(filename, 0000)
	if err != nil {
		t.Error(err)
		return
	}
	err = CheckFileWritePermission(filename)
	if err == nil {
		t.Error("no permisson but no error")
	} else {
		t.Log("expected no permisson error:", err)
	}
}
