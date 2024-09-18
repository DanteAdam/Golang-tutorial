package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	//1
	hobbies := [3]string{"games", "cooking", "fishing"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	//2
	fmt.Println(hobbies[1:])
	//3
	subHobbies := hobbies[:2]
	fmt.Println(subHobbies)
	//4
	subHobbies = subHobbies[1:3]
	fmt.Println(subHobbies)
	//5
	courseGoals := []string{"Learn Go!", "Learn all the basics"}
	fmt.Println(courseGoals)
	//6
	courseGoals[1] = "Learn APIs"
	courseGoals = append(courseGoals, courseGoals[1])
	fmt.Println(courseGoals)
	//7
	products := []Product{
		{"first-product",
			"a first product",
			9.99},
		{"second-product",
			"a second product",
			12.59},
	}
	fmt.Println(products)

	newProduct := Product{
		"third-product",
		"a third product",
		15.99,
	}
	products = append(products, newProduct)
	fmt.Println(products)
}
