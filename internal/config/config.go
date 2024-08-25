package config

import "flag"

type config struct {
	Addr            string
	StaticDirectory string
}

func Load() *config {
	var cfg config

	flag.StringVar(&cfg.Addr, "SNIPPETBOX_ADDR", ":4000", "Address of Snippetbox HTTP Server")
	flag.StringVar(&cfg.StaticDirectory, "SNIPPETBOX_STATIC", "./ui/static", "Path to static assets")

	return &cfg
}
