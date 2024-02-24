package main

import "os"

func main() {
	if len(os.Args) < 2 {
		println("help + run default")
		return
	}

	switch os.Args[1] {
	case "h", "help":
		println("help")
	case "v", "ver", "version":
		println("version : 0.1.0")
	case "build", "generate":
		println("generate static webpages")
	case "serve":
		println("generate static webpages, then serve via localhost")
	case "local", "localhost":
		println("use localhost or 127.0.0.1 or 0.0.0.0")
	default:
		println("help + run default")
	}
}
