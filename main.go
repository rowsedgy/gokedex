package main

func main() {
	cfg := &Config{
		ID:        1,
		Increment: 20,
	}

	startRepl(cfg)
}
