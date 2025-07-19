##add some stuff here for building and cleanup 
clean:
	rm -f /tmp/bookapi

tidy:
	go mod tidy
run:
	@echo "Cleaning project first.."
	$(MAKE) clean
	go build -o /tmp/bookapi
	go run main.go