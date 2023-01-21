package main

import (
	"fmt"
)

func main() {

    const image_width int = 256
    const image_height int = 256

	fmt.Println("P3")
	fmt.Println(image_width, image_height)
	fmt.Println(255)

	var ir, ig, ib int
	var r, g, b float64

	for j:= image_height -1; j >= 0 ; j-- {
		for i := 0 ; i < image_width ; i++ {
			
			r = float64(i)/float64(image_height-1)
			g = float64(j)/float64(image_width-1)
			b = 0.25
			
			ir = int(255.999 * r)
			ig = int(255.999 * g)
			ib = int(255.999 * b)

			fmt.Println(ir,ig,ib)

}
	}
	
}
