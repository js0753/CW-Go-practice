package main

import "fmt"

func main() {
	td := teldir{}

	c := 0
	for i := 0; i > -1; i++ {
		fmt.Println("Choose \n1.Add\n2.Update\n3.Search\n4.Find Duplicates\n5.Delete\n6.Print Directory\n7.Load Directory from File\n8.Exit")
		fmt.Scanln(&c)
		if c == 1 {
			for k := 0; k > -1; k++ {
				var name string
				var number int
				var email string
				var address string
				fmt.Println("Enter Name, number ,email and address : ")

				_, err := fmt.Scanln(&name, &number, &email, &address)
				if err != nil {
					//fmt.Println(err)
					break
				}
				td.add(name, number, email, address)
			}
		} else if c == 2 {
			var name string
			var number int
			fmt.Println("Enter Name and updated number,email,address : ")
			fmt.Scanln(&name, &number)
			td.update(name, number)
		} else if c == 3 {
			var name string
			fmt.Println("Enter Name : ")
			fmt.Scanln(&name)
			td.searchByName(name)
		} else if c == 4 {
			td.find_dup()
		} else if c == 5 {
			var name string
			fmt.Println("Enter Name : ")
			fmt.Scanln(&name)
			td.delete(name)
		} else if c == 6 {
			td.print()
		} else if c == 7 {
			td = TelDirFromFile("teldir")
		} else {
			break
		}
	}

}

/*

Eren 9823346162 eren@gmail.com Mumbai
Zeke 9823346162 zeke@gmail.com Pune
Mikasa 6825546164 mikasa@gmail.com Delhi
Levi 6451184626 levi@gmail.com Chennai
Jean 4627719336 jean@gmail.com Kolkata
Sasha 5182039448 sasha@gmail.com Mumbai
-1
*/
