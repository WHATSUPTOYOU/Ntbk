```go
package main

import (
	"fmt"
	"image"
	"os"
	"path/filepath"

	"github.com/vitali-fedulov/images3"
)

var files []string

func walkFunc(path string, info os.FileInfo, err error) error {
	files = append(files, path)
	return nil
}

func main() {
	Path := "C:\\Users\\chenhaoer\\Desktop\\imageTest"
	filepath.Walk(Path, walkFunc)
	f, err := os.Open(files[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(files)
	img1, _, _ := image.Decode(f)
	image1 := images3.Icon(img1, "")
	for i, name := range files {
		if i <= 1 {
			continue
		}
		f, err := os.Open(name)
		if err != nil {
			panic(err)
		}
		img2, _, _ := image.Decode(f)
		image2 := images3.Icon(img2, "")
		res := images3.Similar(image1, image2)
		th1, th2, th3 := images3.EucMetric(image1, image2)
		fmt.Printf("Image %d, PropMetric %f\n", i, images3.PropMetric(image1, image2))
		fmt.Printf("Image %d, EucMetric th1 %f, th2 %f, th3 %f\n", i, th1, th2, th3)
		fmt.Printf("Image %d, Similarity %v\n", i, res)
		image1 = image2
	}

	// f, _ = os.Open("C:\\Users\\chenhaoer\\Desktop\\imageTest\\1691392429.png")
	// img1, _, _ = image.Decode(f)
	// icon1 := images3.Icon(img1, "")
	// f2, _ := os.Open("C:\\Users\\chenhaoer\\Desktop\\imageTest\\1691392430.png")
	// img2, _, _ := image.Decode(f2)
	// icon2 := images3.Icon(img2, "")
	// fmt.Printf("Similarity %v\n", images3.Similar(icon1, icon2))
}```
