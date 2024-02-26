package setting

type group struct {
	Config    config
	Logger    log
	Dao       mDao
	EmailMark emailMark
	Worker    worker
	Maker     maker
	Oss       oss
	Auto      auto
}

var Group = new(group)

func AllInit() {
	Group.Config.Init()
	Group.Logger.Init()
	Group.Dao.Init()
	Group.EmailMark.Init()
	Group.Worker.Init()
	Group.Maker.Init()
	Group.Oss.Init()
	Group.Auto.Init()
}
