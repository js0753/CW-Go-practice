package main

import "fmt"

func main() {
	//emp_count_map := make(map[Employee]int)
	for i := 0; i < 5; i++ {
		var name string
		var age int
		var empid int
		var salary int
		var dept string
		fmt.Println("Enter employee details : ")
		_, err := fmt.Scanln(&name, &age, &empid, &salary, &dept)
		if err != nil {
			fmt.Println(err)
		}
		var e Employee
		e.construct(name, age, empid, salary, dept)
	}
	e_max := new(Employee)
	e_max.salary = 0
	for i, employee := range employee_slice {
		fmt.Println(i, employee.name, employee.empid, employee.salary)
		if employee.salary > e_max.salary {
			e_max.name = employee.name
			e_max.empid = employee.empid
			e_max.salary = employee.salary
		}

		e6 := new(Employee)
		e6.construct("Jean", 25, 1289, 75000, "EXTC") //duplicate
		for j := 0; j < i; j++ {
			if employee.empid == employee_slice[j].empid {
				fmt.Println("Duplicate Found for Employee ID : ", employee.empid)
			}
		}
	}
	fmt.Println("Person with Max Salary : ")
	fmt.Println(e_max.name, e_max.empid, e_max.salary)

}

/*
John 25 1234 50000 Computer
Bruce 27 1678 80000 Mech
Peter 21 1085 100000 Computer
Wanda 22 1274 90000 IT
Jean 25 1289 75000 EXTC
*/
