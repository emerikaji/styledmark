# Styledmark

Styledmark is a command-line utility for parsing markdown files into an html file. A css file for basic styling is included and automatically added to the html.

This program is written in Go and uses the [golang-commonmark](https://gitlab.com/golang-commonmark/markdown) package to parse markdown. Its developers have already made a command line utility with more options; Styledmark is intended to streamline the process for better productivity.

This README is also intended as an example file to feed to styledmark, so feel free to use it as input to check out the styling!

## header 2

### header 3

#### header 4

##### header 5

###### header 6

this is example golang code:
```go
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
```