package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

type config struct {
	filePath    string
	profilePath string
	file        *os.File
	profile     *os.File
}

var cfg config

func init() {
	flag.StringVar(&cfg.filePath, "o", "", "output image path")
	flag.StringVar(&cfg.profilePath, "prof", "", "write cpu profile to file")
}

func initConfig() {
	if cfg.profilePath != "" {
		var err error
		cfg.profile, err = os.Create(cfg.profilePath)
		if err != nil {
			log.Panic(err)
		}
		pprof.StartCPUProfile(cfg.profile)
		defer pprof.StopCPUProfile()
	}
	if cfg.filePath == "" {
		log.Panic(fmt.Errorf("-o not set"))
	}
	var err error
	cfg.file, err = os.OpenFile(cfg.filePath,
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0755)
	if err != nil {
		log.Panic(err)
	}
}

func cleanUp() {
	handleFileClose(cfg.file)
	handleFileClose(cfg.profile)
}

func handleFileClose(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Panic(err)
	}
}
