package main

import "fmt"

func main() {
	//emp_count_map := make(map[Employee]int)
	c := 0
	for i := 0; i > -1; i++ {
		fmt.Println("Choose \n1.New Employee DB\n2.Load Employee DB\n3.Save Employee DB\n4.Find Duplicates\n5.Find Max Salary\n6.Print Employee DB\n7.Exit")
		fmt.Scanln(&c)
		if c == 1 {
			for i := 0; i > -1; i++ {
				var name string
				var age int
				var empid int
				var salary int
				var dept string
				fmt.Println("Enter employee details : ")
				_, err := fmt.Scanln(&empid, &name, &age, &salary, &dept)
				if err != nil {
					fmt.Println(err)
					break
				}
				if name == "-1" {
					break
				}
				var e Employee
				e.construct(empid, name, age, salary, dept)
				employee_slice = append(employee_slice, e)
			}
		} else if c == 2 {
			fmt.Println("Enter File Name : ")
			var filename string
			fmt.Scanln(&filename)
			employee_slice = EmployeeSliceFromFile(filename)
		} else if c == 3 {
			var filename string
			fmt.Println("Enter File Name : ")
			fmt.Scanln(&filename)
			employee_slice.saveToFile(filename)
		} else if c == 4 {
			//e6 := new(Employee) // .\main.go:41:27: cannot use e6 (type *Employee) as type Employee in append
			//var e6 Employee                               // new returns a pointer of type Employee while var creates a variable of type Employee
			//e6.construct(1289, "Jean", 25, 75000, "EXTC") //duplicate
			//employee_slice = append(employee_slice, e6)
			findDuplicate(employee_slice)
		} else if c == 5 {
			e_max := findMax(employee_slice)
			fmt.Println("Person with Max Salary : ")
			fmt.Println(e_max.name, e_max.empid, e_max.salary)
		} else if c == 6 {
			employee_slice.print()
		} else {
			break
		}
	}

}

/*
1234 John 25 50000 Computer
1678 Bruce 27 80000 Mech
1085 Peter 21 100000 Computer
1274 Wanda 22 90000 IT
1289 Jean 25 75000 EXTC
1289 Jean 25 75000 EXTC
-1
*/
