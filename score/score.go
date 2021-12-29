package score

import (
	"errors"
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

func GetScore() (*Score, error) {
	data, err := os.ReadFile("./data.txt")
	score := Score{}

	if err != nil {
		log.Fatalf("An error occured while opening \"data.txt\". Err: %s", err)
		return nil, err
	}

	split := strings.Split(string(data), " ")

	if len(split) != 3 {
		log.Fatal("The data.txt has to be formed like this: X X X.")
		return nil, errors.New("The data.txt has to be formed like this: X X X.")
	}

	fillProperty(&score.Football, split[0])
	fillProperty(&score.Baseball, split[1])
	fillProperty(&score.Basketball, split[2])

	return &score, nil
}
