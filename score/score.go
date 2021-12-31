package score

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Score struct {
	Football   int
	Baseball   int
	Basketball int
}

func fillProperty(property *int, value string) {
	tmp, err := strconv.Atoi(value)
	*property = tmp

	if err != nil {
		log.Fatalf("Could not convert the file content to int. Err: %s", err)
	}
}

func getLastLineWithSeek(filepath string) (string, error) {
	fileHandle, err := os.Open(filepath)

	if err != nil {
		log.Fatalf("An error occured while opening \"data.txt\". Err: %s", err)
		return "", err
	}

	defer fileHandle.Close()

	line := ""
	var cursor int64 = 0
	stat, _ := fileHandle.Stat()
	filesize := stat.Size()
	for {
		cursor -= 1
		fileHandle.Seek(cursor, io.SeekEnd)

		char := make([]byte, 1)
		fileHandle.Read(char)

		if cursor != -1 && (char[0] == 10 || char[0] == 13) { // stop if we find a line
			break
		}

		line = fmt.Sprintf("%s%s", string(char), line) // there is more efficient way

		if cursor == -filesize { // stop if we are at the begining
			break
		}
	}

	return line, nil
}

func GetScore() (*Score, error) {
	line, err := getLastLineWithSeek("./data.txt")
	score := Score{}

	if err != nil {
		return nil, err
	}

	split := strings.Split(string(line), " ")

	if len(split) != 3 {
		log.Fatal("The data.txt has to be formed like this: X X X.")
		return nil, errors.New("The data.txt has to be formed like this: X X X.")
	}

	fillProperty(&score.Football, split[0])
	fillProperty(&score.Baseball, split[1])
	fillProperty(&score.Basketball, split[2])

	return &score, nil
}
