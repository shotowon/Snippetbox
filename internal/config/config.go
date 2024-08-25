package config

import "flag"

type config struct {
	Env             string
	Addr            string
	StaticDirectory string
	LogFile         string
}

func Load() *config {
	var cfg config

	flag.StringVar(&cfg.Env, "SNIPPETBOX_ENV", EnvLocal, "Environment in which Snippetbox is running")
	flag.StringVar(&cfg.Addr, "SNIPPETBOX_ADDR", ":4000", "Address of Snippetbox HTTP Server")
	flag.StringVar(&cfg.StaticDirectory, "SNIPPETBOX_STATIC", "./ui/static", "Path to static assets")
	flag.StringVar(&cfg.LogFile, "SNIPPETBOX_LOGFILE", "snippetbox.log", "File to write logs")

	flag.Parse()
	return &cfg
}
