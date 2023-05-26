package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func fileStringWrite(str []string) error {
	fileSrc, errOpen := os.Create("wikipython.txt")
	defer fileSrc.Close()
	if errOpen != nil {
		return errOpen
	}

	writer := bufio.NewWriter(fileSrc)
	defer writer.Flush()

	for i := 0; i < len(str); i++ {
		_, errWrt := fmt.Fprintln(writer, str[i])
		if errWrt != nil {
			return errWrt
		}
	}
	return nil
}

func displayConsole(occurence int, listLine []int) {
	fmt.Println("=== Summary ===")
	fmt.Println("Number of occurence of Go : ", occurence)
	fmt.Println("Number of lines where 'go' is present : ", len(listLine))
	fmt.Println("Lines : ", listLine)
	fmt.Println("=== End of Summary ===")
}

// Parcourt la chaine en remplacant 'python' à chaque occurence de Go
func stringTransform(stringSrc []string) ([]string, int, []int) {
	occurence := 0
	listLine := make([]int, 0)
	for i := 0; i < len(stringSrc); i++ {
		if strings.Contains(stringSrc[i], "go") || strings.Contains(stringSrc[i], "Go") {
			occurence += strings.Count(stringSrc[i], "Go")
			occurence += strings.Count(stringSrc[i], "go")
			stringSrc[i] = strings.Replace(stringSrc[i], "go", "python", -1)
			stringSrc[i] = strings.Replace(stringSrc[i], "Go", "Python", -1)
			listLine = append(listLine, i)
		}
	}
	return stringSrc, occurence, listLine
}

// Récupère les données du fichier .txt et le met dans un tableau
func fileStringRec() ([]string, error) {
	fileSrc, errSrc := os.Open("wikigo.txt")
	defer fileSrc.Close()
	if errSrc != nil {
		return nil, errSrc
	}

	lineIndex := 1
	scan := bufio.NewScanner(fileSrc)
	fmt.Println(scan.Text())
	stringTab := make([]string, 0)
	for ; scan.Scan(); lineIndex++ {
		stringTab = append(stringTab, scan.Text())
	}

	return stringTab, nil
}

func main() {
	input, error := fileStringRec()
	if error != nil {
		fmt.Println(error)
		return
	}
	input, occurence, listLine := stringTransform(input)
	displayConsole(occurence, listLine)
	fileStringWrite(input)
}
