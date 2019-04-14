package config

type root struct {
	appList []*app
}

func (r *root) app(name string) *app {
	ret := &app{
		name: name,
	}
	r.appList = append(r.appList, ret)
	return ret
}

type app struct {
	name          string
	locationList  []string
	portList      []int
	containerList []*container
}

func (a *app) locations(locs []string) {
	a.locationList = locs
}

func (a *app) container(image string) *container {
	envelope := newContainer(image)
	a.containerList = append(a.containerList, envelope)
	return envelope
}
