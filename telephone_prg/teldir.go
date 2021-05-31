package main

import "fmt"

//type teldir map[int]Entry
type Entry struct {
	name    string
	number  int
	email   string
	address string
}

func (t *teldir) add(name string, number int) {
	if (*t)[name] == 0 {
		(*t)[name] = number
	} else {
		fmt.Println("Entry Already Exists")
	}
}
func (t *teldir) update(name string, number int) {
	(*t)[name] = number
}
func (t *teldir) delete(name string) {
	delete(*t, name)
}

func (t teldir) search(name string) {
	for nm, num := range t {
		if nm == name {
			fmt.Println(nm+"'s Number is : ", num)
		}
	}
}

func (t teldir) print() {
	fmt.Println("Numbers in Directory : ")
	for name, num := range t {
		fmt.Println(name, num)
	}
}

func (t teldir) find_dup() {
	var entry_list []Entry
	for name, num := range t {
		e := new(Entry)
		e.name = name
		e.number = num
		entry_list = append(entry_list, *e)
	}
	for i, ent := range entry_list {
		for j := 0; j < i; j++ {
			if ent.number == entry_list[j].number {
				fmt.Println("Duplicate Entries Found  : \n", ent.name, ent.number, "\n", entry_list[j].name, entry_list[j].number)
			}
		}
	}
}
