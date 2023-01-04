package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	trkveh    vehicle
	fourWheel bool
}
type sedan struct {
	sedveh vehicle
	luxury bool
}

func main() {

	v1 := vehicle{6, "violet"}

	//t1 := truck{truck.vehicle.doors:2, color:"black", fourWheel: true}
	//t1=truck{}
	trck1 := truck{
		trkveh:    vehicle{doors: 2, color: "black"},
		fourWheel: true,
	}

	sedan1 := sedan{
		sedveh: vehicle{doors: 4, color: "red"},
		luxury: true,
	}

	fmt.Println(v1, trck1, sedan1)

	fmt.Println("vehicle color", v1.color)

	fmt.Println("trck1 color: ", trck1.trkveh.color)
	fmt.Println("sedan1 doors: ", sedan1.sedveh.doors)

}
