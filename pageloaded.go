package main

import (
	"fmt"
	"github.com/tjgq/sane"
)

func main() {
	err := sane.Init()
	if err != nil {
		fmt.Printf("Init failed. $s\n",err)
		return
	}
	c, err := sane.Open("epjitsu:libusb:001:007")
	if err != nil {
		fmt.Printf("No scanner. $s\n",err)
		return
	}
	inf, err := c.SetOption("source", "ADF Duplex")
	fmt.Printf("Option: %+v, %v\n",inf, err)
	option, _ := c.GetOption("page-loaded")
	if option.(bool) {
		for err == nil {
			_, err = c.ReadImage()
			fmt.Printf("Scaned a page.\n")
		}
		fmt.Printf("%+v\n", err)
	}
	c.Close()
}

