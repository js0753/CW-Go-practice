package main

import "fmt"

func main() {
	//emp_count_map := make(map[Employee]int)
	/*
		e1 := new(Employee)
		e1.construct("John", 25, 1234, 50000, "Computer")
		//John 25 1234 50000 Computer
		e2 := new(Employee)
		e2.construct("Bruce", 27, 1678, 80000, "Mech")
		//Bruce 27 1678 80000 Mech
		e3 := new(Employee)
		e3.construct("Peter", 21, 1085, 100000, "Computer")
		//Peter 21 1085 100000 Computer
		e4 := new(Employee)
		e4.construct("Wanda", 22, 1274, 90000, "IT")
		//Wanda 22 1274 90000 IT
		e5 := new(Employee)
		e5.construct("Jean", 25, 1289, 75000, "EXTC")
		//Jean 25 1289 75000 EXTC
		e6 := new(Employee)
		e6.construct("Jean", 25, 1289, 75000, "EXTC") //duplicate
	*/
	e_max := new(Employee)

	for i := 0; i < 5; i++ {
		var name string
		var age int
		var empid int
		var salary int
		var dept string
		fmt.Println("Enter employee details : ")
		//fmt.Scanln("%s %d %d %d %s", &name, &age, &empid, &salary, &dept)
		_, err := fmt.Scanln(&name, &age, &empid, &salary, &dept)
		if err != nil {
			fmt.Println(err)
		}
		var e Employee
		e.construct(name, age, empid, salary, dept)
	}
	e6 := new(Employee)
	e6.construct("Jean", 25, 1289, 75000, "EXTC") //duplicate
	e_max.salary = 0
	for i, employee := range employee_slice {
		fmt.Println(i, employee.name, employee.empid, employee.salary)
		if employee.salary > e_max.salary {
			//fmt.Println("Here")
			e_max.copy(employee)
		}
		/*
			value, ok := emp_count_map[employee]
			fmt.Println(value)
			if ok == true {
				emp_count_map[employee] = emp_count_map[employee] + 1
			} else {
				emp_count_map[employee] = 1
			}*/
		for j := 0; j < i; j++ {
			if employee.empid == employee_slice[j].empid {
				fmt.Println("Duplicate Found for Employee ID : ", employee.empid)
			}
		}
	}
	fmt.Println("Person with Max Salary : ")
	fmt.Println(e_max.name, e_max.empid, e_max.salary)

}
