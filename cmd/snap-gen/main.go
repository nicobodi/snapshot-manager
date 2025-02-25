package main

import (
	"log"
	"os"

	"github.com/nicobodi/snapshot-manager/internal/config"
	"github.com/nicobodi/snapshot-manager/internal/snapshot"
)

func main() {
	vol := os.Args[1]

	conf, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	err = snapshot.Snapshot(conf.SnapRoot, vol)
	if err != nil {
		log.Fatal(err)
	}
}
