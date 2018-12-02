package days

type Factory struct {}

func (f Factory) GetDay(name string) Day {
	switch name {
	case "one":
		return One{}
		break
	}

	return One {}
}
