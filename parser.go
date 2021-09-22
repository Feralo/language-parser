package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"unicode"
)

const INT_LIT = 10
const IDENT = 11
const ASSIGN_OP = 20
const ADD_OP = 21
const SUB_OP = 22
const MULT_OP = 23
const DIV_OP = 24
const LEFT_PAREN = 25
const RIGHT_PAREN = 26

var charClass int
var nextChar rune
var lexeme [100]rune
var lexLen int
var token int
var nextToken int
var in_fp *bufio.Scanner

var isAlpha = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

const LETTER = 0
const DIGIT = 1
const UNKNOWN = 99
const EOF = 98

func lookup(ch rune) int {
	switch ch {
	case '(':
		addChar()
		nextToken = LEFT_PAREN
		break
	case ')':
		addChar()
		nextToken = RIGHT_PAREN
		break
	case '+':
		addChar()
		nextToken = ADD_OP
		break
	case '-':
		addChar()
		nextToken = SUB_OP
		break
	case '*':
		addChar()
		nextToken = MULT_OP
		break
	case '/':
		addChar()
		nextToken = DIV_OP
		break
	default:
		addChar()
		nextToken = EOF
		break
	}
	return nextToken
}

func addChar() {
	lexeme[lexLen] = nextChar
	lexLen++
	lexeme[lexLen] = 0
}

func getChar() {
	in_fp.Scan()
	nextString := in_fp.Text()
	if len(nextString) > 0 {
		nextChar = []rune(nextString)[0]
		if (65 <= nextChar && nextChar <= 90) || (97 <= nextChar && nextChar <= 122) {
			charClass = LETTER
		} else if unicode.IsDigit(nextChar) {
			charClass = DIGIT
		} else {
			charClass = UNKNOWN
		}
	} else {
		charClass = EOF
	}

}

func getNonBlank() {
	for unicode.IsSpace(nextChar) {
		if charClass != EOF {
			getChar()
		} else {
			lex()
		}
	}
}

func lex() int {
	lexLen = 0
	if charClass == EOF {
		fmt.Println("Next token is 98, Next lexeme is EOF")
		nextToken = EOF
	} else {
		switch charClass {
		case LETTER:
			addChar()
			getChar()
			for charClass == LETTER || charClass == DIGIT {
				addChar()
				getChar()
			}
			nextToken = IDENT
			break
		case DIGIT:
			addChar()
			getChar()
			for charClass == DIGIT {
				addChar()
				getChar()
			}
			nextToken = INT_LIT
			break
		case UNKNOWN:
			lookup(nextChar)
			getChar()
			break
		}
		fmt.Print("Next token is: ")
		fmt.Print(nextToken)
		fmt.Print(", Next lexeme is: ")
		for i := 0; i < lexLen; i++ {
			fmt.Print(string(lexeme[i]))
		}
		fmt.Println()

	}
	return nextToken
}

func main() {
	filename := "test.txt"

	filebuffer, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputdata := string(filebuffer)
	in_fp = bufio.NewScanner(strings.NewReader(inputdata))
	in_fp.Split(bufio.ScanRunes)

	for charClass != EOF {
		lex()
	}

}
