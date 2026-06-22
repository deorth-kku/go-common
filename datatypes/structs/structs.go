package cstructs

import (
	"encoding/json/v2"
	"io"

	"github.com/deorth-kku/go-common/args"
)

func ToMap(stc any) (m map[string]any, err error) {
	rd, wt := io.Pipe()
	defer rd.Close()
	var encerr error
	go func() {
		encerr = json.MarshalWrite(wt, stc)
		wt.Close()
	}()
	err = json.UnmarshalRead(rd, &m)
	if encerr != nil {
		return nil, encerr
	}
	return
}

func MustToMap(stc any) (m map[string]any) {
	return args.Must(ToMap(stc))
}
