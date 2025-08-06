package desc

type String string

func (me String) String() string {
	return string(me)
}

func Str(s string) String {
	return String(s)
}
