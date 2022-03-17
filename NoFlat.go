package noflat

type NoFlat struct {
	getterSetter GetterSetter
}

func Init(g GetterSetter) *NoFlat {
	return &NoFlat{g}
}

func (noflat *NoFlat) Get(k string) string {
	return noflat.getterSetter.Get(k)
}

func (noflat *NoFlat) Set(k string, v string) {
	noflat.getterSetter.Set(k, v)
}

func (noflat *NoFlat) Group(k string) *Group {
	return &Group{k + "_", noflat.getterSetter}
}
