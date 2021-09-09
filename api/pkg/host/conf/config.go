package conf

func newConfig() *Config {
	return &Config{
		App: newDefaultAPP(),
		Log: newDefaultLog(),
		MySQL:newDefaultMySQL(),
	}
}

//保存配置文件
type Config struct {
	App   *app   `toml:"app"`
	Log   *log   `toml:"log"`
	MySQL *mySQL `toml:"mysql"`
}

type app struct {
	Name string `toml:"name" env:"App_NAME"`
	Host string `toml:"host" env:"APP_HOST"`
	Port string `toml:"port" env:"APP_PORT"`
}

func newDefaultAPP() *app {
	return &app{
		Name: "demo",
		Host: "127.0.0.1",
		Port: "3306",
	}
}

type log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
}

//newDefaultLog todo
func newDefaultLog() *log {
	return &log{
		Level:   "debug",
		PathDir: "logs",
		Format:  "text",
		To:      "stdout",
	}
}

//MySQL
type mySQL struct {
	Host        string `toml:"host" env:"D_MYSQL_HOST"`
	Port        string `toml:"port" env:"D_MYSQL_PORT"`
	UserName    string `toml:"username" env:"D_MYSQL_USERNAME"`
	Password    string `toml:"password" env:"D_MYSQL_PASSWORD"`
	Database    string `toml:"database" env:"D_MYSQL_DATABASE"`
	MaxOpenConn int    `toml:"max_open_conn" env:"D_MYSQL_MAX_OPEN_CONN"`
	MaxIdleConn int    `toml:"max_idle_conn" env:"D_MYSQL_MAX_IDLE_CONN"`
	MaxLifeTime int    `toml:"max_life_time" env:"D_MYSQL_MAX_LIFE_TIME"`
	MaxIdleTime int    `toml:"max_life_time" env:"D_MYSQL_MAX_idle_TIME"`
}


func newDefaultMySQL() *mySQL{
	return &mySQL{
		Database:"test",
		Host:"127.0.0.1",
		Port:"3306",
		MaxOpenConn:200,
		MaxIdleConn: 50,
		MaxLifeTime: 1800,
		MaxIdleTime: 600,
	}
}