package config

import "flag"

type config struct {
	Env             string
	Addr            string
	StaticDirectory string
}

func Load() *config {
	var cfg config

	flag.StringVar(&cfg.Env, "SNIPPETBOX_ENV", "LOCAL", "Environment in which Snippetbox is running")
	flag.StringVar(&cfg.Addr, "SNIPPETBOX_ADDR", ":4000", "Address of Snippetbox HTTP Server")
	flag.StringVar(&cfg.StaticDirectory, "SNIPPETBOX_STATIC", "./ui/static", "Path to static assets")

	flag.Parse()
	return &cfg
}
