package utils

import (
	"encoding/json"
	"fmt"
)

func Wrap(kvs ...interface{}) map[string]interface{} {
	if kvs != nil {
		ret := map[string]interface{}{}
		for i := 0; i < len(kvs)-1; i += 2 {
			if k, ok := kvs[i].(string); ok {
				ret[k] = kvs[i+1]
			}
		}
		return ret
	}
	return nil
}

func WrapMessage(kvs ...interface{}) string {
	// be careful to pass exploded slice to func Wrap
	ret := Wrap(kvs...)
	if ret != nil {
		b, err := json.Marshal(ret)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		return string(b)
	}
	return ""
}
