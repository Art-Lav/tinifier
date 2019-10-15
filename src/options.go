package main

import (
	"errors"
	"github.com/jessevdk/go-flags"
	"strings"
)

type Options struct {
	Verbose        bool     `short:"v" long:"verbose" description:"Show verbose debug information"`
	ShowVersion    bool     `short:"V" long:"version" description:"Show version and exit"`
	DisableColors  bool     `short:"C" long:"no-colors" description:"Disable color output"`
	FileExtensions []string `short:"e" long:"ext" default:"jpg,JPG,jpeg,JPEG,png,PNG" description:"Target file extensions"`
	ApiKey         string   `short:"k" long:"api-key" env:"TINYPNG_API_KEY" description:"API key <https://tinypng.com/dashboard/api>"`
	Threads        int      `short:"t" long:"threads" default:"5" description:"Threads processing count"`
	Targets        struct {
		Path []string `positional-arg-name:"files-and-directories"`
	} `positional-args:"yes" required:"true"`
}

var options Options

// Parse options using fresh parser instance.
func (o *Options) Parse() (*flags.Parser, []string, error) {
	var p = flags.NewParser(o, flags.Default)
	var s, err = p.Parse()

	return p, s, err
}

// Make options check.
func (o *Options) Check() error {
	// Check API key
	if key := strings.TrimSpace(o.ApiKey); len(key) <= 1 {
		return errors.New("tinypng.com API key is not provided")
	}

	// Check threads count
	if o.Threads <= 0 {
		return errors.New("threads count cannot be less then 1")
	}

	return nil
}
