package config

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

func ParseFlag(app *kingpin.Application) (*Config, error) {
	var resp Config
	app.Flag("letters", "Up to 7 characters").Short('l').Required().StringVar(&resp.Letters)
	app.Flag("dico", "Path to a dictionary file").Short('d').Default("./assets/dico/fr.txt").StringVar(&resp.DicoPath)
	if _, err := app.Parse(os.Args[1:]); err != nil {
		return nil, err
	}
	return &resp, nil
}
