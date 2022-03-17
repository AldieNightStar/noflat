package noflat

import "strings"

type Group struct {
	prefix       string
	getterSetter GetterSetter
}

func (g *Group) Get(k string) string {
	return g.getterSetter.Get(g.prefix + k)
}

func (g *Group) Set(k string, v string) {
	g.addKeyIfNotExists(k)
	g.setInner(k, v)
}

func (g *Group) setInner(k string, v string) {
	g.getterSetter.Set(g.prefix+k, v)
}

func (g *Group) Exists(k string) bool {
	s := g.Get("#keys")
	if s == "" {
		return false
	}
	for _, el := range strings.Split(s, ";") {
		if el == k {
			return true
		}
	}
	return false
}

func (g *Group) addKeyIfNotExists(k string) {
	if !g.Exists(k) {
		keys := strings.Split(g.Get("#keys"), ";")
		if len(keys) == 1 && keys[0] == "" {
			// When keys list len is 1 and has empty value
			keys = []string{k}
		} else {
			keys = append(keys, k)
		}
		keysStr := strings.Join(keys, ";")
		g.setInner("#keys", keysStr)
	}
}

func (g *Group) SubGroup(k string) *Group {
	return &Group{k + "_", g}
}

func (g *Group) Keys() []string {
	arr := make([]string, 0, 8)
	for _, key := range strings.Split(g.Get("#keys"), ";") {
		if strings.Contains(key, "#") {
			continue
		}
		arr = append(arr, key)
	}
	return arr
}
