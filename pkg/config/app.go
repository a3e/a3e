package config

func app(name string) *appContainer {
	return &appContainer{
		name: name,
	}
}

type appContainer struct {
	name          string
	locationList  []string
	portList      []int
	containerList []*container
}

func (a *appContainer) locations(locs []string) {
	a.locationList = locs
}

func (a *appContainer) container(image string) *container {
	envelope := newContainer(image)
	a.containerList = append(a.containerList, envelope)
	return envelope
}
