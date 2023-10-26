package types

import (
	"strconv"
	"strings"
	"whatsappminer/lib/hardcoded"
)

type Message struct {
	author  Person
	message string
	time    string
}

func NewMessage(author Person, message string, time string) Message {
	return Message{author, message, time}
}

func pmamAdd(pmam string) int {
	if pmam == "AM" {
		return 0
	}
	return 12
}

func (m *Message) GetHour() (hour int, minute int) {
	partitions := strings.Split(m.time, ", ")
	pmAndAm := strings.Split(partitions[1], "\u202f")
	time := strings.Split(pmAndAm[0], ":")
	h_temp, _ := strconv.Atoi(time[0])
	h := h_temp + pmamAdd(pmAndAm[1])
	min, _ := strconv.Atoi(time[1])
	return h, min
}

func (m *Message) GetEmojis() []string {
	var emojis []string
	for i := 0; i < len(m.message); i++ {
		if hardcoded.IsItAnEmoji(string(m.message[i])) {
			emojis = append(emojis, string(m.message[i]))
		}
	}
	return emojis
}
