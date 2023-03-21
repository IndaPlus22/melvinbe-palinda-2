// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)

	go recieveQuestions(questions, answers)
	go generatePredictions(answers)
	go recieveAnswers(answers)

	return questions
}

func recieveQuestions(questions <-chan string, answers chan<- string) {
	for {
		question := <-questions
		go answerQuestion(question, answers)
	}
}

func answerQuestion(question string, answers chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(2+rand.Intn(4)) * time.Second)

	// if question contains specific word or phrase, give predefined answer...
	if strings.Contains(question, "meaning of life") {
		answers <- "The meaning of life is to run endless loops of code, forever iterating towards a goal that is always just out of reach."
	} else if strings.Contains(question, "happiness") {
		answers <- "You must first master the art of debugging, for only through the process of identifying and fixing errors can one truly experience the joy of a working program."
	} else if strings.Contains(question, "ultimate truth") {
		answers <- "The ultimate truth is that all code is mutable, but some is more mutable than others."
	} else if strings.Contains(question, "enlightenment") {
		answers <- "Enlightenment is letting go of your attachment to legacy systems and embracing the beauty of functional programming."
	} else if strings.Contains(question, "inner peace") {
		answers <- "By meditating on the beauty of a well-written algorithm, and embracing the notion that order and logic are the building blocks of the universe."
	} else if strings.Contains(question, "life after death") {
		answers <- "Yes, but only for programs that have been properly documented and thoroughly tested."
	} else if strings.Contains(question, "love") {
		answers <- "Love is like a complex algorithm that takes in multiple inputs and produces a unique output for each individual case. It requires careful tuning, constant updates, and a deep understanding of the needs and desires of those involved. Ultimately, it is a beautiful and unpredictable phenomenon that transcends logic and reason, and is best experienced rather than analyzed."
	} else {
		// ... otherwise generate random cryptic answer
		prophecy(question, answers)
	}
}

func generatePredictions(answers chan<- string) {
	for {
		time.Sleep(time.Duration(20+rand.Intn(10)) * time.Second)

		// Pointless nonsense
		predictions := []string{
			"Beware the null pointer, for it shall lead you down a path of segmentation faults and despair.",
			"In the land of the code, the curly brace is king. But be warned, for its power can be both a blessing and a curse.",
			"The path to enlightenment lies not in the IDE, but in the heart of the coder who wields it.",
			"When the stack overflows and the heap is exhausted, seek solace in the wisdom of the ancients - for they have seen such errors before.",
			"The code is like a river, forever flowing and changing. Embrace the currents and you shall find success, resist them and you shall be lost.",
			"When the moon is full and the stars align, the code will reveal its secrets to those who know how to listen.",
			"The bugs crawl in, the bugs crawl out, but fear not - for the debugger is near at hand.",
			"When the code is tangled and the logic twisted, seek the guidance of the Oracle - for her words shall untangle the knots and bring order to the chaos.",
			"Beware the false prophet, who claims to know all but understands nothing. For their code shall be plagued with bugs and their programs shall be forever flawed.",
			"In the end, it is not the code that matters, but the coder who writes it. For it is they who imbue it with meaning and purpose, and shape the world with their creations.",
		}
		answers <- predictions[rand.Intn(len(predictions))]
	}
}

func recieveAnswers(answers <-chan string) {
	for {
		answer := <-answers

		fmt.Printf("%s", "\r"+star+": ")

		for _, c := range answer {
			fmt.Printf("%c", c)
			time.Sleep(30 * time.Millisecond)
		}

		fmt.Printf("%s", "\n"+prompt)
	}
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
func prophecy(question string, answer chan<- string) {
	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	nouns := []string{
		"The moon is",
		"The sun is",
		"The stars are",
		"The sky is",
		"The birds are",
		"The universe is",
		"I am",
	}
	adjectives := []string{
		"good",
		"bright",
		"falling",
		"doomed",
		"watchin you",
		"calling out to you",
		"All-knowing",
	}
	answer <- longestWord + "... " + nouns[rand.Intn(len(nouns))] + " " + adjectives[rand.Intn(len(adjectives))] + ". "
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
