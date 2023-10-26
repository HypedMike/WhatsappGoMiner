package lib

import (
	"bufio"
	"fmt"
	"os"
	"whatsappminer/lib/types"
)

func ElaborateChat(path string) types.Chat {
	// read file
	readFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	// get lines
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	// init chat and add lines
	chat := types.EmptyChat(path)
	for _, v := range fileLines {
		chat.AddRawLine(v)
	}

	fmt.Printf("Total number of texts: %d\n", chat.GetTotalNumberTexts())
	chat.GetTotalNumberTextsFromPeople()
	fmt.Print(chat.MostUsedWords()[:10])

	hours := chat.NumberOfTextsPerHour(true)

	for j, h := range hours {
		fmt.Print(j + 1)
		for i := 0; i < h; i++ {
			fmt.Print("|")
		}
		fmt.Print("\n")
	}

	return chat

	//chat.Print()
}
