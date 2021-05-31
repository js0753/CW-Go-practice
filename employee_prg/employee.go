package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Employee struct {
	empid  int
	name   string
	age    int
	salary int
	dept   string
}

type EmployeeSlice []Employee

var employee_slice EmployeeSlice

func EmployeeSliceFromFile(filename string) EmployeeSlice {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// opt1 - log err and return call to newDeck()
		// opt2 - log err and quit program
		fmt.Println("ERROR:", err)
		os.Exit(1)

	}
	var temp_es EmployeeSlice
	s := strings.Split(string(bs), "\n")
	for _, row := range s {
		values := strings.Split(row, ",")
		if len(values) > 1 {
			var e Employee
			//fmt.Println(values[2])
			empid, _ := strconv.Atoi(values[0])
			age, _ := strconv.Atoi(values[2])
			salary, _ := strconv.Atoi(values[3])
			e.construct(empid, values[1], age, salary, values[4])

			temp_es = append(temp_es, e)
		}
	}
	//temp_es.print()
	return temp_es
}

func (es EmployeeSlice) saveToFile(filename string) error {
	main_string := ""
	for _, row := range es {
		row_string := fmt.Sprint(row.empid) + "," + row.name + "," + fmt.Sprint(row.age) + "," + fmt.Sprint(row.salary) + "," + row.dept
		// conversion from int to string yields a string of one rune, not a string of digits (did you mean fmt.Sprint(x)?)
		main_string = main_string + row_string + "\n"
	}
	return os.WriteFile(filename, []byte(main_string), 0666)
}

func (e *Employee) construct(empid int, name string, age int, salary int, dept string) {
	e.name = name
	e.age = age
	e.empid = empid
	e.salary = salary
	e.dept = dept

}

func findMax(employee_slice []Employee) *Employee {
	e_max := new(Employee)
	e_max.salary = 0
	for i, employee := range employee_slice {
		fmt.Println(i, employee.name, employee.empid, employee.salary)
		if employee.salary > e_max.salary {
			e_max.name = employee.name
			e_max.empid = employee.empid
			e_max.salary = employee.salary
		}
	}
	return e_max
}

func findDuplicate(employee_slice []Employee) {
	for i, employee := range employee_slice {
		for j := 0; j < i; j++ {
			if employee.empid == employee_slice[j].empid {
				fmt.Println("Duplicate Found for Employee ID : ", employee.empid)
			}
		}
	}
}

func (es EmployeeSlice) print() {
	fmt.Println("\nEmployee Databse : ")
	for _, e := range es {
		fmt.Println(e.empid, e.name, e.age, e.salary, e.dept)
	}
	fmt.Println("")
}
