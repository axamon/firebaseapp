package main

import "flag"

var auth = flag.String("auth", "auth.json", "File autenticazione  di google")

func main() {

	flag.Parse()

}
