package stdin

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Init() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		fmt.Println(string((scanner.Bytes())))

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error on reading %p", err)
	}
}
