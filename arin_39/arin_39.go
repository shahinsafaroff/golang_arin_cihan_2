/*package main

import "fmt"

type myType int

func main() {

	var x myType
	x = 10

	fmt.Printf("%T, %v", x, x)

}
*/

/*package main

import "fmt"

func main() {

	x := 10
	y := &x
	*y = 20

	fmt.Println(x)

}
*/

/*package main
import "fmt"

type rectangle struct {
	a, b int
}
func (r rectangle) display() {
	fmt.Println("Width", r.a, " and ", "Length", r.b)
}
func (r rectangle) area() int {
	return r.a * r.b
}
func (r rectangle) circumference() int{
	return 2*(r.a + r.b)
}
func main(){
	r1 := rectangle{3,8}
	r1.display()
	fmt.Println("Square is: ", r1.area())
	fmt.Println("Perimeter is: ", r1.circumference())
}*/

package main

import "fmt"

type family struct {
	name string
	age int
	isMarried bool
}

func showFamily() []family {
	family1 := family {
		name: "Arin",
		age: 5,
		isMarried: false,
	}

	family2 := family{
		name: "Elis",
		age: 3,
	}

	family3 := family{
		"Gurgen",
		40,
		true,
	}

	var family4 family
	family4.name = "Gamze"
	family4.age = 40
	family4.isMarried = true

	return []family {family1, family2, family3, family4}
}

func main() {
	families := showFamily()
	for i:=0; i < len(families); i++ {
		fmt.Println(families[i])
	}
}