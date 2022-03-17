package noflat

type GetterSetter interface {
	Get(string) string
	Set(string, string)
}
