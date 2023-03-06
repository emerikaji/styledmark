package main

import (
	"errors"
	"flag"
	"os"

	"gitlab.com/golang-commonmark/markdown"
)

const (
	css = `@import url('https://fonts.googleapis.com/css2?family=Merriweather:ital,wght@0,300;0,400;0,700;0,900;1,300;1,400;1,700;1,900&display=swap');
body {
	font-family: 'Merriweather', serif;
	font-size: 16px;
	margin: 32px 64px 32px 64px;
}

h1,h2,h3,h4,h5,h6 {
	font-weight: 900;
}

h1 {
	font-size: 64px;
	text-align: center;
}

h2 {
	font-size: 48px;
}

h3 {
	font-size: 32px;
}

h4 {
	font-size: 24px;
}

h5 {
	font-size: 16px;
}

h6 {
	font-size: 12px;
}

a {
	text-decoration: none;
	color:darkslateblue;
}

a:hover {
	color:slateblue;
}

pre {
	border: 2px;
	border-style: solid;
	border-color: black;
	word-wrap: normal;
	white-space: pre-wrap;
	tab-size: 4;
}	
`
	head = `<html lang="fr">
<head>
<meta charset="UTF-8">
<link rel="stylesheet"
      href="//cdn.jsdelivr.net/gh/highlightjs/cdn-release@11.7.0/build/styles/obsidian.min.css">
<script src="//cdn.jsdelivr.net/gh/highlightjs/cdn-release@11.7.0/build/highlight.min.js"></script>
<script>hljs.highlightAll();</script>
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
