package lib

import (
	"bufio"
	"fmt"
	"os"
	"whatsappminer/lib/types"
)

func ElaborateChat(path string) types.Chat {
	readFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	chat := types.EmptyChat(path)
	for _, v := range fileLines {
		chat.AddRawLine(v)
	}

	fmt.Println(chat.GetTotalNumberTexts())
	chat.GetTotalNumberTextsFromPeople()
	fmt.Println(chat.MostUsedWords()[:10])

	return chat

	//chat.Print()
}
