package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const apiUrl = "https://opentdb.com/api.php?amount=5&category=18&difficulty=easy&type=boolean&encode=base64"

type QuestionResponse struct {
	ResponseCode int `json:"response_code"`
	Results      []struct {
		Category         string   `json:"category"`
		Type             string   `json:"type"`
		Difficulty       string   `json:"difficulty"`
		Question         string   `json:"question"`
		CorrectAnswer    string   `json:"correct_answer"`
		IncorrectAnswers []string `json:"incorrect_answers"`
	} `json:"results"`
}

func main() {

	questions, err := GetQuestionsFromServer()
	if err != nil {
		exit("Failed to get questions from server")
	}

	scanner := bufio.NewScanner(os.Stdin)
	for i, q := range questions.Results {
		fmt.Println(fmt.Sprintf("%d. Question: %s", i+1, decodeString(q.Question)))

		correctAnswer := strings.ToLower(decodeString(q.CorrectAnswer))
		var userAnswer string

		scanner.Scan()
		text := scanner.Text()
		userAnswer = strings.ToLower(text)

		if correctAnswer == userAnswer {
			fmt.Println("The answer is correct!")
			continue
		}

		fmt.Println("The answer is wrong!")

	}
}

func GetQuestionsFromServer() (QuestionResponse, error) {

	var err error
	var questionResponse QuestionResponse

	c := http.Client{Timeout: time.Duration(1) * time.Second}

	res, err := c.Get(apiUrl)
	if err != nil {
		err = fmt.Errorf("error %s", err)
		return questionResponse, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			exit(fmt.Sprintf("Error %s", err))
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &questionResponse)
	if err != nil {
		return questionResponse, err
	}
	return questionResponse, err
}

func decodeString(s string) string {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		exit("Can not decode string")
	}
	return string(b)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
