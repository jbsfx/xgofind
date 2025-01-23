package util

import "fmt"

func PrintBanner() {
	banner := `
	
XGOFIND - http/https interceptor
==========================================
`
	fmt.Printf("%v\nVersion: %v | %v\n\n", banner, Version, Author)
}
