package main

import (
	"errors"
	"flag"
	"os"

	"gitlab.com/golang-commonmark/markdown"
)

const (
	css = `h1, h2, h3, h4, h5, h6 {
	color: black;
}

a {
	font-style: normal;
}
`
	head = `<html lang="fr">
<head>
<meta charset="UTF-8">
<link href="styledmark.css" rel="stylesheet">
<title>Out</title>
</head>
<body>
`
	tail = `</body>
</html>`
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
	err = os.WriteFile(*nameFlag+".html", []byte(head+md.RenderToString(b)+tail), 0666)
	if err != nil {
		panic(err)
	}
	_, err = os.Stat("styledmark.css")
	if os.IsNotExist(err) {
		err = os.WriteFile("styledmark.css", []byte(css), 0666)
		if err != nil {
			panic(err)
		}
	}
}
