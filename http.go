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

const esc = "\x00"

// If not found, FileMode = 0600
func FileWithMode(str string) (file string, fm os.FileMode, found bool, err error) {
	str = strings.Replace(str, "\\,", esc, -1)
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
		fm = 0000
	}
	file = strings.Replace(file, esc, ",", -1)
	return
}

func ParseListen(listen string) (net.Listener, error) {
	if strings.HasPrefix(listen, "@") {
		// Abstract unix socket (Linux only)
		return net.ListenUnix("unix", &net.UnixAddr{Name: "\x00" + listen[1:], Net: "unix"})
	} else if filepath.IsAbs(listen) {
		f, m, found, err := FileWithMode(listen)
		if err != nil {
			return nil, err
		}
		addr := &net.UnixAddr{Name: f, Net: "unix"}
		lis, err := net.ListenUnix("unix", addr)
		if err != nil {
			return nil, err
		}
		if found {
			err = os.Chmod(f, m)
		}
		return lis, err
	} else {
		var addr *net.TCPAddr
		addr, err := net.ResolveTCPAddr("tcp", listen)
		if err != nil {
			return nil, err
		}
		return net.ListenTCP("tcp", addr)
	}
}

func (se *HttpServer) ListenAndServe(listen string) error {
	lis, err := ParseListen(listen)
	if err != nil {
		return err
	}
	err = se.Server.Serve(lis)
	if err == http.ErrServerClosed {
		return nil
	}
	return err
}
