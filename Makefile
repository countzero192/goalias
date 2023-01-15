NAME=goal
INST=/usr/local/bin

build: main.go
	go build -o $(NAME) main.go 

clean:
	rm -f $(NAME)

install: goal
	cp $(NAME) $(INST) 

uninstall:
	rm -f $(INST)/$(NAME)
