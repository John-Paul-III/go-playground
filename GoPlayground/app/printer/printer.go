package printer


type printer struct {
	value string
}

func (p printer) Value() string {
	return p.value
}

func NewPrinter(value string) printer {
	if value != "Hi" {
		return printer{value}
	}
	return printer{""}
}
