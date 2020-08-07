package config

type App struct {
	MakeMigration bool `yaml:"MakeMigration"` //是否进行数据库迁移
	DBLog         bool `yaml:"DBLog"`
}

func (app *App) DefaultAppConfig() {
	app.MakeMigration = false
	app.DBLog = true
}
