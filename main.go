package main

import (
	"github.com/jjhwan-h/CBNU_DIDI_BLOCKCHAIN/cli"
	"github.com/jjhwan-h/CBNU_DIDI_BLOCKCHAIN/db"
)

func main() {
	defer db.Close()

	cli.Start()

}
