package main

import "os"

func main() {
	if len(os.Args) < 2 {
		println("No arguments!")
		PrintHelp()
		return
	}

	switch os.Args[1] {
	case "h", "-h", "help", "--help":
		PrintHelp()
	case "v", "-v", "ver", "--ver", "version", "--version":
		println("version : 0.1.0")
	case "build", "--build", "generate", "--generate":
		println("generate static webpages")
	case "serve", "--serve":
		println("generate static webpages, then serve via localhost")
	case "local", "--local", "localhost", "--localhost":
		println("Let's open http://localhost:8080")
		LocalServe()
	default:
		println("Unknown argument!")
		PrintHelp()
	}
}

func PrintHelp() {
	println(`
Usage:
	kmt <arg>
Use one of these commands:

	b | -b | build | --build | generate | --generate
		create the static files

	v | -v | ver | --ver | version | --version
		show kmt version

	h | -h | help | --help
		show this help

	serve | --serve
		generate the static files, serve them on localhost/server

	local | --local | localhost | --localhost
		use localhost or 127.0.0.1 or 0.0.0.0
			`)
}
