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
	markdownFlag = flag.String("input", "", "path to input file")
	flag.StringVar(markdownFlag, "i", *markdownFlag, "alias for -input")
	nameFlag = flag.String("output", "out", "name of output file")
	flag.StringVar(nameFlag, "o", *nameFlag, "alias for -output")
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
