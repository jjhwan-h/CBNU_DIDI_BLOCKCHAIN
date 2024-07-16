package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/jjhwan-h/DIDI_BLOCKCHAIN/explorer"
	"github.com/jjhwan-h/DIDI_BLOCKCHAIN/rest"
)

func usage() {
	fmt.Printf("Welcome to coin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port=: Set the PORT of the server\n")
	fmt.Printf("-mode=: Choose between 'html' and 'rest'\n")
	runtime.Goexit() //모든게 종료, defer수행 후
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}
	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest' ")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
	fmt.Println(*port, *mode)
}
