package config

import "flag"

type Config struct {
	Env             string
	Addr            string
	StaticDirectory string
	LogFile         string
	MainDSN         string
}

func Load() *Config {
	var cfg Config

	flag.StringVar(&cfg.Env, "SNIPPETBOX_ENV", EnvLocal, "Environment in which Snippetbox is running")
	flag.StringVar(&cfg.Addr, "SNIPPETBOX_ADDR", ":4000", "Address of Snippetbox HTTP Server")
	flag.StringVar(&cfg.StaticDirectory, "SNIPPETBOX_STATIC", "./ui/static", "Path to static assets")
	flag.StringVar(&cfg.LogFile, "SNIPPETBOX_LOGFILE", "snippetbox.log", "File to write logs")
	flag.StringVar(&cfg.MainDSN, "SNIPPETBOX_MAIN_DSN", "postgres://snippetbox:snippetbox@localhost:5050/snippetbox", "link to main DSN of snippetbox")

	flag.Parse()
	return &cfg
}
