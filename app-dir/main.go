package main

import toolkit "github.com/kaungmyathan22/golang-toolkit"

func main() {
	var tools toolkit.Tools
	err := tools.CreateDirIfNotExists("./test-dir")
	if err != nil {
		panic(err)
	}
}
