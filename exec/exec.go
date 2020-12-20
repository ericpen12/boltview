package exec

func Run(c Command) {
	if err := c.Parse(); err != nil {
		c.Error(err)
	}

	if err := c.Exec(); err != nil {
		c.Error(err)
		return
	}
	c.Ok()
}
