package main

import "flag"

type Config struct {
	Add    string
	Remove int
	Toggle int
}

func ParseFlags() Config {
	var config Config

	flag.StringVar(&config.Add, "add", "", "Add a new todo item")
	flag.IntVar(&config.Remove, "remove", -1, "Remove the todo item at given index")
	flag.IntVar(&config.Remove, "toggle", -1, "Toggle todo item completion state at given index")

	flag.Parse()

	return config
}
