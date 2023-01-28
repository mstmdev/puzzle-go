package _map

type GoMap map[any]any

func (m GoMap) Len() int {
	return len(m)
}

func (m GoMap) Set(k, v any) {
	m[k] = v
}

func (m GoMap) Delete(k any) {
	delete(m, k)
}

func (m GoMap) Get(k any) any {
	return m[k]
}

func (m GoMap) GetOk(k any) (v any, ok bool) {
	v, ok = m[k]
	return v, ok
}

func (m GoMap) AllData() []Data {
	var allData []Data
	for k, v := range m {
		allData = append(allData, newData(k, v))
	}
	return allData
}
