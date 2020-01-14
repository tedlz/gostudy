package files

// Values *
type Values map[string][]string

// Get *
func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

// Add *
func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}
