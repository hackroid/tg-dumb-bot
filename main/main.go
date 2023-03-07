package main

func main() {
	initBot()
	addHandler("help", help)
	addHandler("choice", choice)
	addHandler("status", status)
	addHandler("default", ddefault)
	startHandler()
}
