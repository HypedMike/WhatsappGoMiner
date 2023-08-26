package types

type Message struct {
	author  Person
	message string
	time    string
}

func NewMessage(author Person, message string, time string) Message{
	return Message{author, message, time}
}