package common

import (
	"testing"
)

func TestAll(t *testing.T) {
	ret := ResolvePlugins()
	for k, v := range ret {
		t.Log(k, v)

		if p, ok := v.(IPlugin); ok {
			p.DoWork()
		}
	}
}
