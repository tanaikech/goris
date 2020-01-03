// Package main (goris.go) :
// This file is included all commands and options.
package main

import (
	"fmt"
	"os"

	"github.com/tanaikech/goris/ris"
	"github.com/urfave/cli"
)

const (
	appname = "goris"
)

// dispres : Display results
func dispres(r []string, c int) {
	if len(r) < c {
		c = len(r)
	}
	for i := 0; i < c; i++ {
		fmt.Printf("%s\n", r[i])
	}
}

// handler : Handler of goris
func handler(c *cli.Context) error {
	if len(c.String("fromurl")) == 0 && len(c.String("fromfile")) == 0 {
		fmt.Fprintf(os.Stderr, "Error: No parameters. You can see help by '$ %s -h'\n", appname)
		os.Exit(1)
	}
	var results []string
	if len(c.String("fromurl")) > 0 && len(c.String("fromfile")) == 0 {
		results = ris.DefImg(c.Bool("webpages")).ImgFromURL(c.String("fromurl"))
	}
	if len(c.String("fromurl")) == 0 && len(c.String("fromfile")) > 0 {
		results = ris.DefImg(c.Bool("webpages")).ImgFromFile(c.String("fromfile"))
	}
	if c.Bool("download") && (len(c.String("fromurl")) > 0 || len(c.String("fromfile")) > 0) && !c.Bool("webpages") {
		ris.Download(results, c.Int("number"))
	}
	dispres(results, c.Int("number"))
	return nil
}

// main : Main of goris
func main() {
	app := cli.NewApp()
	app.Name = appname
	app.Authors = []*cli.Author{
		{Name: "tanaike [ https://github.com/tanaikech/goris ] ", Email: "tanaike@hotmail.com"},
	}
	app.Usage = "Search for images with Google Reverse Image Search."
	app.Version = "1.1.1"
	app.Commands = []*cli.Command{
		{
			Name:        "search",
			Aliases:     []string{"s"},
			Usage:       "[ " + appname + " s -u URL ] or [ " + appname + " -f file ]",
			Description: "Do search images.",
			Action:      handler,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "fromurl, u",
					Aliases: []string{"u"},
					Usage:   "Reverse Image Search from an URL.",
				},
				&cli.StringFlag{
					Name:    "fromfile, f",
					Aliases: []string{"f"},
					Usage:   "Reverse Image Search from an image file.",
				},
				&cli.IntFlag{
					Name:    "number, n",
					Aliases: []string{"n"},
					Usage:   "Number of retrieved image URLs. ( 1 - 100 )",
					Value:   10,
				},
				&cli.BoolFlag{
					Name:    "download, d",
					Aliases: []string{"d"},
					Usage:   "Download images from retrieved URLs.",
				},
				&cli.BoolFlag{
					Name:    "webpages, w",
					Aliases: []string{"w"},
					Usage:   "This is boolean. Retrieve web pages with matching images on Google top page. When this is not used, images are retrieved.",
				},
			},
		},
	}
	app.Run(os.Args)
}
