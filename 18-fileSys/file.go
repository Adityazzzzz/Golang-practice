package main

import (
	"bufio"
	"fmt"
	"os"
)
// -------------------------------------------------------------------------------------------------
func fileInfoDemo() {
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("Name:", info.Name())
	fmt.Println("Is Directory:", info.IsDir())
	fmt.Println("Size:", info.Size())
	fmt.Println("Permissions:", info.Mode())
	fmt.Println("Last Modified:", info.ModTime())
}

// -------------------------------------------------------------------------------------------------
func readUsingBuffer() {
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 12)
	n, err := file.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bytes read:", n)
	fmt.Println("Content:", string(buf[:n]))
}

// -------------------------------------------------------------------------------------------------
func readWholeFile() {
	data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

// -------------------------------------------------------------------------------------------------
func readDirectory() {
	dir, err := os.Open(".")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	files, err := dir.ReadDir(-1)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		fmt.Println(f.Name(), f.IsDir())
	}
}

// -------------------------------------------------------------------------------------------------
func writeFile() {
	f, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("Hi Go\n")
	f.WriteString("Nice language\n")

	data := []byte("Hello Golang\n")
	f.Write(data)
}

// -------------------------------------------------------------------------------------------------
func copyFile() {
	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create("example_copy.txt")
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		writer.WriteByte(b)
	}

	writer.Flush()
	fmt.Println("File copied successfully")
}

// -------------------------------------------------------------------------------------------------
func deleteFile() {
	err := os.Remove("example_copy.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File deleted successfully")
}

// -------------------------------------------------------------------------------------------------
func main() {
	fileInfoDemo()
	readUsingBuffer()
	readWholeFile()
	readDirectory()
	writeFile()
	copyFile()
	deleteFile()
}
