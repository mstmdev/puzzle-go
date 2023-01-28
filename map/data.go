package _map

type Data struct {
	K any
	V any
}

func newData(k, v any) Data {
	return Data{K: k, V: v}
}
