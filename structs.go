package common

import (
	"encoding/json"
	"io"
)

func Struct2Map(stc any) (m map[string]any, err error) {
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

func MustStruct2Map(stc any) (m map[string]any) {
	return Must(Struct2Map(stc))
}
