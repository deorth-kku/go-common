package common

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type HttpServer struct {
	*http.Server
	*http.ServeMux
}

func NewHttpServer() (server HttpServer) {
	server.ServeMux = http.NewServeMux()
	server.Server = &http.Server{Handler: server.ServeMux}
	return
}

// If not found, FileMode = 0600
func FileWithMode(str string) (file string, fm os.FileMode, err error) {
	file, m, found := strings.Cut(str, ",")
	if found {
		var t uint64
		t, err = strconv.ParseUint(m, 8, 32)
		if err != nil {
			err = fmt.Errorf("failed to parse mode %s :%s", m, err)
			return
		}
		fm = os.FileMode(t)
	} else {
		fm = 0600
	}
	return
}

func (se *HttpServer) ListenAndServe(listen string) (err error) {
	var lis net.Listener
	if filepath.IsAbs(listen) {
		f, m, err := FileWithMode(listen)
		if err != nil {
			return err
		}
		if m != 0 {
			Umask(int(^m))
		}
		addr := &net.UnixAddr{Name: f, Net: "unix"}
		lis, err = net.ListenUnix("unix", addr)
	} else {
		var addr *net.TCPAddr
		addr, err = net.ResolveTCPAddr("tcp", listen)
		if err != nil {
			return
		}
		lis, err = net.ListenTCP("tcp", addr)
	}
	if err != nil {
		return
	}
	err = se.Server.Serve(lis)
	if err == http.ErrServerClosed {
		err = nil
	}
	return
}
