package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"os"
	"sort"
)

type person struct {
	First string //need to first letter capital to turn into JSON
	Last  string
	Age   int
}

type byAge []person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type byName []person

func (a byName) Len() int           { return len(a) }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool { return a[i].First < a[j].First }

func main() {
	p1 := person{
		First: "Bimo",
		Last:  "Pangestu",
		Age:   21,
	}

	p2 := person{
		First: "Natasha",
		Last:  "OK",
		Age:   15,
	}

	people := []person{p1, p2}

	fmt.Println(people)
	x, err := json.Marshal(people) //marshaling data
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(x))
	retPeople := []person{}
	err = json.Unmarshal(x, &retPeople)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("retPeople =", retPeople)
	fmt.Println("LOL")
	fmt.Fprintln(os.Stdout, "LOL")
	io.WriteString(os.Stdout, "LOL")

	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	fmt.Println(xi)

	sort.Strings(xs)
	fmt.Println(xs)

	p3 := person{"James", "Bond", 32}
	p4 := person{"Moneypenny", "miss", 27}
	p5 := person{"Q", "Dr", 64}
	p6 := person{"M", "Dr", 56}

	peoples := []person{p3, p4, p5, p6}

	//sort Custom
	fmt.Println(peoples)
	sort.Sort(byAge(peoples))
	fmt.Println(peoples)
	sort.Sort(byName(peoples))
	fmt.Println(peoples)

	//bcrypt
	s := `password`

	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("------------")
	fmt.Println(s)
	fmt.Println("------------")
	fmt.Println(bs)

	loginPass := `password1`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPass))
	if err != nil {
		fmt.Println("Wrong Password")
	} else {
		fmt.Println("Success LOGIN")
	}
}
