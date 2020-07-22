package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	os.Setenv("TEST", "eyyoo")
	err := godotenv.Load()
	if err != nil {
		panic("Error loading env")
	}
	fmt.Println(os.Getenv("SECRET_KEY"))

}
func main() {
	fmt.Println(os.Getenv("TEST"))
	fmt.Println(os.Getenv("TESTENV"))
	fmt.Println(os.Getenv("S3_BUCKET"))
	fmt.Println(os.Getenv("SECRET_KEY"))
}
