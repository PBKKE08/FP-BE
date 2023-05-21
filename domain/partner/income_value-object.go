package partner

type Income struct {
	value string
}

func (i Income) GetValue() string {
	return i.value
}
