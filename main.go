/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package main

import (
	"github.com/joho/godotenv"
	"github.com/spolivin/jobtracker/cmd"
)

func main() {
	_ = godotenv.Load()
	cmd.Execute()
}
