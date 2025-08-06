package must

func Get[V any](v V, e error) V {
	if e != nil {
		panic(e)
	}
	return v
}

func Have[V any](v V, found bool) V {
	if !found {
		panic("does not have")
	}
	return v
}
