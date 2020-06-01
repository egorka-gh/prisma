package main

import (
	"fmt"

	"github.com/egorka-gh/sm/prisma"
)

func main() {
	srv := prisma.DefaultServer(func(m *prisma.Message) { fmt.Println(m) })

	err := srv.Run()

	fmt.Println(err)
}
