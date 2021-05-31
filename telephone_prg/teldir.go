package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//type teldir map[int]Entry
type Entry struct {
	name    string
	number  int
	email   string
	address string
}

type teldir []Entry

func TelDirFromFile(filename string) teldir {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// opt1 - log err and return call to newDeck()
		// opt2 - log err and quit program
		fmt.Println("ERROR:", err)
		os.Exit(1)

	}
	var temp_td teldir
	s := strings.Split(string(bs), "\n")
	for _, row := range s {
		values := strings.Split(row, ",")
		if len(values) > 1 {
			var e Entry
			//fmt.Println(values[2])
			e.name = values[0]
			e.number, _ = strconv.Atoi(values[1])
			e.email = values[2]
			e.address = values[3]
			temp_td = append(temp_td, e)
		}
	}
	//temp_es.print()
	return temp_td
}

func (t *teldir) add(name string, number int, email string, address string) {
	var e Entry
	e.name = name
	e.number = number
	e.email = email
	e.address = address
	*t = append(*t, e)
	e.addToFile("teldir")
}

func (t teldir) saveToFile(filename string) error {
	main_string := ""
	for _, row := range t {
		row_string := row.name + "," + fmt.Sprint(row.number) + "," + row.email + "," + row.address
		// conversion from int to string yields a string of one rune, not a string of digits (did you mean fmt.Sprint(x)?)
		main_string = main_string + row_string + "\n"
	}
	return os.WriteFile(filename, []byte(main_string), 0666)
}

func (e Entry) addToFile(filename string) error {
	entry_string := e.name + "," + fmt.Sprint(e.number) + "," + e.email + "," + e.address
	f, err := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //  0644 == rw-r-r //os.O_WRONLY is to state the file mode, wether read,write, or both, in this case write
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	_, write_err := f.WriteString(entry_string + "\n")

	return write_err
}

func (t *teldir) update(name string, number int) {
	for i, entry := range *t {
		if entry.name == name {
			(*t)[i].number = number
		}
	}
	(*t).saveToFile("teldir")
}
func (t *teldir) delete(name string) {
	//delete(*t, name)
	for i, entry := range *t {
		if entry.name == name {
			*t = append((*t)[:i], (*t)[i+1:]...) // ... lets you pass multiple arguments to a function from a slice. Here append function is called recursively for each element in (*t)[i+1:]
			break
		}
	}
	(*t).saveToFile("teldir")
}

func (t teldir) searchByName(name string) {
	for _, entry := range t {
		if entry.name == name {
			fmt.Println(entry.name+"'s Number is : ", entry.number)
		}
	}
}

func (t teldir) print() {
	fmt.Println("Numbers in Directory : ")
	for _, entry := range t {
		fmt.Println(entry.name, entry.number, entry.email, entry.address)
	}
}

func (t teldir) find_dup() {
	for i, ent := range t {
		for j := 0; j < i; j++ {
			if ent.number == t[j].number {
				fmt.Println("Duplicate Entries Found  : \n", ent.name, ent.number, "\n", t[j].name, t[j].number)
			}
		}
	}
}
