package main

// CGO_CFLAGS=-I/usr/local/Cellar/sane-backends/1.0.24_1/include/ CGO_LDFLAGS=-L/usr/local/Cellar/sane-backends/1.0.24_1/lib/ go get github.com/tburke/sane
// # Fujitsu S1300i
// firmware /usr/local/Cellar/sane-backends/1.0.24_1/share/sane/epjitsu/1300i_0D12.nal
// usb 0x04c5 0x128d
// https://sane-project.gitlab.io/html/
// image fmt https://sane-project.gitlab.io/html/doc008.html
// http://cpansearch.perl.org/src/RATCLIFFE/Sane-0.05/examples/scanadf.pl

import (
	"fmt"
	"github.com/tburke/sane"
	"io"
	"os"
	// "image/png"
	"golang.org/x/image/tiff"
)

func main() {
	err := sane.Init()
	if err != nil {
		fmt.Printf("Init failed. $s\n", err)
		return
	}
	devs, err := sane.Devices()
	if err != nil {
		fmt.Printf("No devices. %s\n", err)
		return
	}
	var name string
	for _, d := range devs {
		fmt.Printf("%#v\n", d)
		name = d.Name
	}
	c, err := sane.Open(name)
	if err != nil {
		fmt.Printf("No scanner. %s\n", err)
		return
	}
	inf, err := c.SetOption("source", "ADF Duplex")
	fmt.Printf("Option: %+v, %v\n", inf, err)
	inf, err = c.SetOption("mode", "Gray")
	fmt.Printf("Option: %+v, %v\n",inf, err)
	option, _ := c.GetOption("page-loaded")
	if option.(bool) {
		r := c.NewReader()
		page := 0
		for r.Next() == nil {
			page += 1
			f, ferr := os.Create(fmt.Sprintf("scan_%d.pnm", page))
			if ferr != nil {
				fmt.Println(ferr)
				return
			}
			n, err := io.Copy(f, r)
			fmt.Printf("Bytes copied: %d\nError: %s\n", n, err)
			f.Close()
		}
	}
	if false {
		var img *sane.Image
		page := 0
		for err == nil {
			img, err = c.ReadImage()
			if err == nil {
				page += 1
				fmt.Printf("Scanned a page.\n%#v\n", img)
				f, ferr := os.Create(fmt.Sprintf("tib%d.pnm", page))
				if ferr != nil {
					fmt.Println(ferr)
					return
				}
				f.Write(img.PNM())
				f.Close()
				f, ferr = os.Create(fmt.Sprintf("tib%d.tiff", page))
				if ferr != nil {
					fmt.Println(ferr)
					return
				}
				tiff.Encode(f, *img, nil)
				f.Close()
			}
		}
		fmt.Printf("%+v\n", err)
	}
	c.Close()
	sane.Exit()
}
