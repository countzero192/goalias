NAME=goalias
INST=/usr/local/bin

$(NAME): main.go
	go build

install: $(NAME)
	cp $(NAME) $(INST) 

uninstall:
	rm -f $(INST)/$(NAME)

clean:
	rm -f $(NAME)
