package main

import(
	"fmt"
	"github.com/schollz/wifiscan"
)

func
main() {
	wifis, err := wifiscan.Scan()
	if err != nil {
	}
	for _, w := range wifis {
		fmt.Println(w.SSID, w.RSSI)
	}
}
