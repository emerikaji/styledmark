build:
	go build

install:
	go install

clean:
	rm styledmark.css
	rm out.html

run:
	./styledmark -i README.md