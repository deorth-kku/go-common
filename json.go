package common

import (
	"io"
	"os"

	"github.com/deorth-kku/iterjson"
)

func IterJsonFprint[K comparable, V any](f io.Writer, data any) (err error) {
	enc := iterjson.NewEncoder[K, V](f)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	err = enc.Encode(data)
	return
}

func IterJsonPrint[K comparable, V any](data any) error {
	return IterJsonFprint[K, V](os.Stdout, data)
}

func JsonFprint(f io.Writer, data any) error {
	return IterJsonFprint[string, map[string]any](f, data)
}

func JsonPrint(data any) error {
	return JsonFprint(os.Stdout, data)
}
