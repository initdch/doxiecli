package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"time"

	"github.com/umahmood/doxiego"
)

func main() {
	doxieGo, err := doxiego.Hello()
	if err != nil {

		fmt.Println(err)
		//...
	}

	fmt.Println("Doxie name", doxieGo.Name)
	fmt.Println(doxieGo.HasPassword)

	// get a list of scanned items on the scanner
	items, err := doxieGo.Scans()
	if err != nil {
		//...
	}

	for _, s := range items {
		fmt.Println("name:", s.Name, "size:", s.Size, "modified:", s.Modified)
	}

	// download a scan

	for _, s := range items {
		img, err := doxieGo.Scan(s.Name)
		if err != nil {
			//...
		}

		// Somewhere in the same package
		fileName := GetFilenameDate(s.Name)
		file, err := os.Create("out/" + fileName)
		if err != nil {
			// Handle error
		}
		defer file.Close()

		jpeg.Encode(file, img, nil)
	}

	//...

}

func GetFilenameDate(filename string) string {
	// Use layout string for time format.
	const layout = "01-02-2006"
	// Place now in the string.
	t := time.Now()
	return t.Format(layout) + "_" + filename
}
