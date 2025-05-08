package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
)

func optionMenu() {

	banner := `
  ____       _   _           _               
 / ___| ___ | | | | __ _ ___| |__   ___ _ __ 
| |  _ / _ \| |_| |/ _| / __| '_ \ / _ \ '__|
| |_| | (_) |  _  | (_| \__ \ | | |  __/ |   
 \____|\___/|_| |_|\__,_|___/_| |_|\___|_|   
`

	fmt.Println(banner)
	fmt.Println("")

	var filename string
	fmt.Print("Enter file path: ")
	fmt.Scan(&filename)

  fmt.Println("")
	fmt.Println("==Options Menu==")
	fmt.Println("==================")
	fmt.Println("md5 -> 1")
	fmt.Println("sha1 -> 2")
	fmt.Println("sha256 -> 3")
	fmt.Println("sha512 -> 4")
	fmt.Println("all -> 5")
	fmt.Println("==================")
	fmt.Println("")

  fmt.Print("Choose an Option: ")
	var option int
	fmt.Scanf("%d", &option)
  fmt.Println("")

	switch option {
	case 1:
		readFile(filename, "md5")
	case 2:
		readFile(filename, "sha1")
	case 3:
		readFile(filename, "sha256")
	case 4:
		readFile(filename, "sha512")
	case 5:
		calcAllTypes(filename)
	}
}

func readFile(filename string, hashtype string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var hash hash.Hash

	switch hashtype {
	case "md5":
		hash = md5.New()
	case "sha1":
		hash = sha1.New()
	case "sha256":
		hash = sha256.New()
	case "sha512":
		hash = sha512.New()
	}

	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s -> %x %s\n", hashtype, hash.Sum(nil), filename)

}

func calcAllTypes(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

  fmt.Printf("file: %s\n", filename)
  fmt.Println("")

	defer file.Close()

	var hash hash.Hash
  
	hash = md5.New()

	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s -> %x\n", "md5", hash.Sum(nil))

	hash = sha1.New()

	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s -> %x\n", "sha1", hash.Sum(nil))

	hash = sha256.New()

	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s -> %x\n", "sha256", hash.Sum(nil))

	hash = sha512.New()

	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s -> %x\n", "sha512", hash.Sum(nil))

}

func main() {
	optionMenu()
	fmt.Println("")
}
