package cstructs

import (
	"encoding/json"
	"io"

	"github.com/deorth-kku/go-common/args"
)

func ToMap(stc any) (m map[string]any, err error) {
	rd, wt := io.Pipe()
	defer rd.Close()
	enc := json.NewEncoder(wt)
	dec := json.NewDecoder(rd)
	var encerr error
	go func() {
		encerr = enc.Encode(stc)
		wt.Close()
	}()
	err = dec.Decode(&m)
	if encerr != nil {
		return nil, encerr
	}
	return
}

func MustToMap(stc any) (m map[string]any) {
	return args.Must(ToMap(stc))
}
