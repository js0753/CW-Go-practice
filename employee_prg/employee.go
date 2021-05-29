package main

type Employee struct {
	name   string
	age    int
	empid  int
	salary int
	dept   string
}

var employee_slice []Employee

func (e Employee) construct(name string, age int, empid int, salary int, dept string) {
	e.name = name
	e.age = age
	e.empid = empid
	e.salary = salary
	e.dept = dept
	employee_slice = append(employee_slice, e)
}
