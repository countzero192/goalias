NAME=goal

build:
	go build -o $(NAME) main.go 

clean:
	rm -f $(NAME)

install:
	cp $(NAME) /usr/local/bin
