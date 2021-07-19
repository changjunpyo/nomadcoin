package main

import (
	"github.com/changjunpyo/nomadcoin/cli"
	"github.com/changjunpyo/nomadcoin/db"
)

func main() {
	defer db.DB().Close()

	cli.Start()
}
