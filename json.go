package pongo2filterjson

import (
	"encoding/json"
	"github.com/flosch/pongo2"
)

func init() {
	RegisterPongo()
}

func RegisterPongo() {
	pongo2.RegisterFilter("json", FilterOutputJson)
}

func FilterOutputJson(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	out, _ := json.Marshal(in.Interface())
	return pongo2.AsValue(string(out)), nil
}
