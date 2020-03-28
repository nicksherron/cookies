// source https://gist.github.com/dacort/bd6a5116224c594b14db

package main

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nicksherron/cookies/pkg"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [domain]\n", os.Args[0])
	os.Exit(2)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	domain := os.Args[1]

	output := ""
	// TODO: Output in Netscape format
	for _, cookie := range cookies.GetCookies(domain) {
		//  device_view=full;
		tmp := fmt.Sprintf("%v=%v; ", cookie.Key, cookie.DecryptedValue())
		output = output + tmp
		//fmt.Printf("%s/%s: %s\n", cookie.Domain, cookie.Key, cookie.DecryptedValue())
	}
	fmt.Println(output)
}
