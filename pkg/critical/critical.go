package critical

type Critical struct {
}

func (c *Critical) DoNothing() {
	println("Maybe something")
}
