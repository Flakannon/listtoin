package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Topic struct {
	Value string
}

func GetFileLines(fileName string) (topicSlice []string, err error) {
	file, e := os.Open(fileName)
	if e != nil {
		return nil, e
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		var t Topic
		t.Value = s[0]
		topicSlice = append(topicSlice, fmt.Sprintf("'%v',", t.Value))
	}

	// remove the last comma
	topicSlice[len(topicSlice)-1] = strings.TrimSuffix(topicSlice[len(topicSlice)-1], ",")

	return topicSlice, nil
}
