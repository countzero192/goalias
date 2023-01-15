NAME=goal
INST=/usr/local/bin

build:
	go build -o $(NAME) main.go 

clean:
	rm -f $(NAME)

install:
	cp $(NAME) $(INST) 

uninstall:
	rm -f $(INST)/$(NAME)
