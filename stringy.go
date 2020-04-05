package string_manipulation

import (
	"errors"
	"math"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"
)

// input is struct that holds input form user and result
type input struct {
	Input  string
	Result string
}

// StringManipulation is an interface that holds all abstract methods to manipulate strings
type StringManipulation interface {
	Between(start, end string) StringManipulation
	Boolean() bool
	CamelCase(rule ...string) string
	ContainsAll(check ...string) bool
	Delimited(delimiter string, rule ...string) StringManipulation
	Get() string
	KebabCase(rule ...string) StringManipulation
	LcFirst() string
	Lines() []string
	Pad(length int, with, padType string) string
	RemoveSpecialCharacter() string
	ReplaceFirst(search, replace string) string
	ReplaceLast(search, replace string) string
	Reverse() string
	Shuffle() string
	Surround(with string) string
	SnakeCase(rule ...string) StringManipulation
	Tease(length int, indicator string) string
	ToLower() string
	ToUpper() string
	UcFirst() string
}

// New func returns pointer to input struct
func New(val string) StringManipulation {
	return &input{Input: val}
}

// Between takes two param start and end which are string
// return value after omitting start and end part of input
func (i *input) Between(start, end string) StringManipulation {
	if start == "" && end == "" || i.Input == "" {
		return i
	}

	input := strings.ToLower(i.Input)
	lcStart := strings.ToLower(start)
	lcEnd := strings.ToLower(end)
	var startIndex, endIndex int

	if len(start) > 0 && strings.Contains(input, lcStart) {
		startIndex = len(start)
	}
	if len(end) > 0 && strings.Contains(input, lcEnd) {
		endIndex = strings.Index(input, lcEnd)
	} else if len(input) > 0 {
		endIndex = len(input)
	}
	i.Result = i.Input[startIndex:endIndex]
	return i
}

// Boolean dunc return boolean value of string value like on, off, 0, 1, yes, no
// returns boolean value of string input
func (i *input) Boolean() bool {
	input := getInput(*i)
	inputLower := strings.ToLower(input)
	off := contains(False, inputLower)
	if off {
		return false
	}
	on := contains(True, inputLower)
	if on {
		return true
	}
	panic(errors.New("invalid string value to test boolean value"))
}

// CamelCase function returns camel case value of string
// Example input: hello user
// Result : HelloUser
func (i *input) CamelCase(rule ...string) string {
	input := getInput(*i)
	// removing excess space
	wordArray := caseHelper(input, true, rule ...)
	for i, word := range wordArray {
		wordArray[i] = ucfirst(word)
	}
	return strings.Join(wordArray, "")
}

// ContainsAll function takes multiple string param and
// checks if they are present in input
func (i *input) ContainsAll(check ...string) bool {
	input := getInput(*i)
	for _, item := range check {
		if !strings.Contains(input, item) {
			return false
		}
	}
	return true
}

// Delimited function joins the string by passed delimeter
func (i *input) Delimited(delimiter string, rule ...string) StringManipulation {
	input := getInput(*i)
	if strings.TrimSpace(delimiter) == "" {
		delimiter = "."
	}
	wordArray := caseHelper(input, false, rule ...)
	i.Result = strings.Join(wordArray, delimiter)
	return i
}

func (i *input) Get() string {
	return getInput(*i)
}

func (i *input) KebabCase(rule ...string) StringManipulation {
	input := getInput(*i)
	wordArray := caseHelper(input, false, rule ...)
	i.Result = strings.Join(wordArray, "-")
	return i
}

func (i *input) LcFirst() string {
	input := getInput(*i)
	for i, v := range input {
		return string(unicode.ToUpper(v)) + input[i+1:]
	}
	return input
}

func (i *input) Lines() []string {
	input := getInput(*i)
	matchWord := regexp.MustCompile(`[\s]*[\W]\pN`)
	result := matchWord.ReplaceAllString(input, " ")
	return strings.Fields(strings.TrimSpace(result))
}

func (i *input) Pad(length int, with, padType string) string {
	input := getInput(*i)
	inputLength := len(input)
	padLength := len(with)
	if inputLength >= length {
		return input
	}
	switch padType {
	case Right:
		var count = 1 + ((length - padLength) / padLength)
		var result = input + strings.Repeat(with, count)
		return result[:length]
	case Left:
		var count = 1 + ((length - padLength) / padLength)
		var result = strings.Repeat(with, count) + input
		return result[(len(result) - length):]
	case Both:
		length := (float64(length - inputLength)) / float64(2)
		repeat := math.Ceil(length / float64(padLength))
		return strings.Repeat(with, int(repeat))[:int(math.Floor(float64(length)))] + input + strings.Repeat(with, int(repeat))[:int(math.Ceil(float64(length)))]
	default:
		return input
	}
}

func (i *input) RemoveSpecialCharacter() string {
	input := getInput(*i)
	var result strings.Builder
	for i := 0; i < len(input); i++ {
		b := input[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' {
			result.WriteByte(b)
		}
	}
	return result.String()
}

func (i *input) ReplaceFirst(search, replace string) string {
	input := getInput(*i)
	return replaceStr(input, search, replace, First)
}

func (i *input) ReplaceLast(search, replace string) string {
	input := getInput(*i)
	return replaceStr(input, search, replace, Last)
}
func (i *input) Reverse() string {
	input := getInput(*i)
	r := []rune(input)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func (i *input) Shuffle() string {
	input := getInput(*i)
	rand.Seed(time.Now().Unix())

	inRune := []rune(input)
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

func (i *input) SnakeCase(rule ...string) StringManipulation {
	input := getInput(*i)
	wordArray := caseHelper(input, false, rule ...)
	i.Result = strings.Join(wordArray, "_")
	return i
}

func (i *input) Surround(with string) string {
	input := getInput(*i)
	return with + input + with
}

func (i *input) Tease(length int, indicator string) string {
	input := getInput(*i)
	if input == "" || len(input) < length {
		return input
	}
	return input[:length] + indicator
}

func (i *input) ToLower() (result string) {
	input := getInput(*i)
	return strings.ToLower(input)
}

func (i *input) ToUpper() string {
	input := getInput(*i)
	return strings.ToUpper(input)
}

func (i *input) UcFirst() string {
	input := getInput(*i)
	return ucfirst(input)
}