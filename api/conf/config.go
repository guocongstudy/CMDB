package conf

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func newConfig() *Config {
	return &Config{
		APP:   newDefaultAPP(),
		Log:   newDefaultLog(),
		Mysql: newDefaultMysql(),
	}
}

type Config struct {
	APP   *app   `toml:"app"`
	Log   *log   `toml:"log"`
	Mysql *mysql `toml:"mysql"`
}

type app struct {
	Name string `toml:"name" env:"APP_NAME"`
	Host string `toml:"host" env:"APP_HOST"`
	Port string `toml:"port" env:"APP_PORT"`
}

func (a *app) Addr() string {
	return fmt.Sprintf("%s:%s", a.Host, a.Port)
}

func newDefaultAPP() *app {
	return &app{
		Name: "demo",
		Host: "127.0.0.1",
		Port: "8080",
	}
}

type log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
}

func newDefaultLog() *log {
	return &log{
		Level:   "debug",
		PathDir: "../../logs/app.log",
		Format:  "text",
		To:      "stdout",
	}
}

type mysql struct {
	Host        string `toml:"host" env:"D_MYSQL_HOST"`
	Port        string `toml:"port" env:"D_MYSQL_PORT"`
	Username    string `toml:"username" env:"D_MYSQL_USERNAME"`
	Password    string `toml:"password" env:"D_MYSQL_PASSWORD"`
	Database    string `toml:"database" env:"D_MYSQL_DATABASE"`
	MaxOpenConn int    `toml:"max_open_conn" env:"D_MYSQL_MAX_OPEN_CONN"`
	MaxIdleConn int    `toml:"max_idle_conn" env:"D_MYSQL_MAX_IDLE_CONN"`
	MaxLifeTime int    `toml:"max_life_time" env:"D_MYSQL_MAX_LIFE_TIME"`
	MaxIdleTime int    `toml:"max_idle_time" env:"D_MYSQL_MAX_IDLE_TIME"`

	lock sync.Mutex
}

var (
	db *sql.DB
)

func (m *mysql) GetDB() (*sql.DB, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if db == nil {
		conn, err := m.getDBConn()
		if err != nil {
			return nil, err
		}

		db = conn
	}
	return db, nil
}

func (m *mysql) getDBConn() (*sql.DB, error) {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true",
		m.Username, m.Password, m.Host, m.Port, m.Database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error, %s", dsn, err)
	}

	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql <%s> error, %s", dsn, err.Error())
	}

	return db, nil
}

func newDefaultMysql() *mysql {
	return &mysql{
		Host:        "127.0.0.1",
		Port:        "3306",
		Username:    "go_course",
		Database:    "go_course",
		MaxOpenConn: 200,
		MaxIdleConn: 50,
		MaxLifeTime: 1000,
		MaxIdleTime: 600,
	}
}
