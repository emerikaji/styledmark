build:
	go build

install:
	go install

clean:
	rm styledmark.css
	rm out.html

test:
	go build
	./styledmark -i README.md