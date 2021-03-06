package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type byAge []user

func (b byAge) Len() int           { return len(b) }
func (b byAge) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byAge) Less(i, j int) bool { return b[i].Age < b[j].Age }

type byLast []user

func (b byLast) Len() int           { return len(b) }
func (b byLast) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byLast) Less(i, j int) bool { return b[i].Last < b[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	sort.Sort(byAge(users))
	for _, v := range users {
		fmt.Println(v.Age)
	}
	fmt.Println("---------------")
	sort.Sort(byLast(users))
	for _, v := range users {
		fmt.Println(v.Last)
	}
}
