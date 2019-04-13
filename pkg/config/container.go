package config

type container struct {
	imageName string
	envList   []envVal
	portList  []int
}

func newContainer(img string) *container {
	return &container{imageName: img}
}

func (c *container) ports(vals []int) {
	c.portList = append(c.portList, vals...)
}

func (c *container) env(name string, def ...string) {
	e := envVal{name: name}
	if len(def) > 0 {
		e.def = def[0]
	}
	c.envList = append(c.envList, e)
}

func (c *container) secretenv(name string) {
	c.envList = append(c.envList, envVal{
		name:   name,
		secret: true,
	})
}
