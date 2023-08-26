package types

import (
	"fmt"
	"strings"
)

type Chat struct {
	id       string
	messages []Message
	people   []Person
}

type Word struct {
	word  string
	count int
}

func EmptyChat(id string) Chat {
	var messages []Message
	var people []Person
	return Chat{id, messages, people}
}

func (c *Chat) checkIfPersonExists(personName string) bool {
	for _, v := range c.people {
		if v.name == personName {
			return true
		}
	}
	return false
}

func (c *Chat) getPersonByName(name string) Person {
	for _, v := range c.people {
		if v.GetName() == name {
			return v
		}
	}

	return Person{}
}

func (c *Chat) AddRawLine(line string) {
	firstPatition := strings.Split(line, " - ") // timestap - data
	if len(firstPatition) < 2 {
		return
	}
	timestamp := strings.Split(firstPatition[0], ", ") // day - time
	if len(timestamp) < 2 {
		return
	}
	date := line[0:19]
	//time := timestamp[1]
	data := strings.Split(firstPatition[1], ": ")
	if len(data) < 2 {
		return
	}
	personName := data[0]
	text := data[1]
	var person Person

	if !c.checkIfPersonExists(personName) {
		person = Person{len(c.people) + 1, personName}
		c.people = append(c.people, person)
	} else {
		person = c.getPersonByName(personName)
	}

	message := NewMessage(person, text, date)

	c.messages = append(c.messages, message)
}

func (c *Chat) Print() {
	for _, v := range c.messages {
		fmt.Printf("%s -> %s\n", v.author.GetName(), v.message)
	}
}

// STATISTICS

func (c *Chat) GetTotalNumberTexts() int {
	return len(c.messages)
}

func (c *Chat) getTotalNumberTextsFromPerson(personName string) int {
	var count int
	for _, v := range c.messages {
		if v.author.GetName() == personName {
			count++
		}
	}
	return count
}

func (c *Chat) MostUsedWords() []Word {
	var words []Word

	for _, m := range c.messages {
		for _, w := range strings.Split(m.message, " ") {

			wr := strings.ToLower(w)

			found := false

			// search into array
			for i, v := range words {
				if v.word == wr {
					found = true
					words[i].count++
				}
			}

			if !found {
				if len(wr) > 3 && wr != "<media" && wr != "omessi>" {
					words = append(words, Word{wr, 1})
				}
			}

		}
	}

	return sortWordsByCount(words)

}

func sortWordsByCount(words []Word) []Word {
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if words[i].count < words[j].count {
				words[i], words[j] = words[j], words[i]
			}
		}
	}
	return words
}

func (c *Chat) GetTotalNumberTextsFromPeople() {
	for _, v := range c.people {
		fmt.Println(v.name, c.getTotalNumberTextsFromPerson(v.name))
	}
}

func (c *Chat) NumberOfTextsPerHour() {

}
