package main

import "fmt"

func main() {
	td := teldir{}

	c := 0
	for i := 0; i > -1; i++ {
		fmt.Println("Choose \n1.Add\n2.Update\n3.Search\n4.Find Duplicates\n5.Delete\n6.Print Directory\n7.Exit")
		fmt.Scanln(&c)
		if c == 1 {
			var name string
			var number int
			fmt.Println("Enter Name and number : ")
			fmt.Scanln(&name, &number)
			td.add(name, number)
		} else if c == 2 {
			var name string
			var number int
			fmt.Println("Enter Name and updated number : ")
			fmt.Scanln(&name, &number)
			td.update(name, number)
		} else if c == 3 {
			var name string
			fmt.Println("Enter Name : ")
			fmt.Scanln(&name)
			td.search(name)
		} else if c == 4 {
			td.find_dup()
		} else if c == 5 {
			var name string
			fmt.Println("Enter Name : ")
			fmt.Scanln(&name)
			td.delete(name)
		} else if c == 6 {
			td.print()
		} else {
			break
		}
	}

}

/*

Eren 9823346162
Zeke 9823346162
Mikasa 6825546164
Levi 6451184626
Jean 4627719336
Sasha 5182039448
*/
