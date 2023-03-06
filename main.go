package main

import (
	"errors"
	"flag"
	"os"

	"gitlab.com/golang-commonmark/markdown"
)

const (
	head = "<html lang=\"fr\">\n<head>\n<meta charset=\"UTF-8\">\n<link href=\"style.css\" rel=\"stylesheet\">\n<title>Out</title>\n</head>\n<body>\n"
	tail = "</body>\n</html>\n"
)

var (
	md           *markdown.Markdown
	nameFlag     *string
	markdownFlag *string
)

func init() {
	md = markdown.New(markdown.XHTMLOutput(true), markdown.HTML(true))
	markdownFlag = flag.String("in", "", "Path to input file")
	nameFlag = flag.String("out", "out", "Name of output file")
	flag.Parse()
}

func main() {
	if *markdownFlag == "" {
		panic(errors.New("no input file given"))
	}
	b, err := os.ReadFile(*markdownFlag)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(*nameFlag+".html", []byte(head+md.RenderToString(b)+tail), 0644)
	if err != nil {
		panic(err)
	}
}
