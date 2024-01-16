package setting

type group struct {
	Config config
	Logger log
}

var Group = new(group)

func AllInit() {
	Group.Config.Init()
	Group.Logger.Init()
}
