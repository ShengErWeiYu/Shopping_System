package main

import "Shopping_System/boot"

func main() {
	boot.DatabaseInit()
	boot.InitRouters()
}
