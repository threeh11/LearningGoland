package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

// Создаем структуру для сортировки
type sortKey struct {
	letter rune
	count  int
}

// Создаем элемент дерева
type elementThree struct {
	element string
	count   int
	parentNode,
	leftChild,
	rightChild *elementThree
	code byte
}

func main() {
	resultAfterReadFile := readFile("test.txt")
	tableLetterFrequency := createTableLetterFrequency(resultAfterReadFile)
	createHuffmanThreeWithTableLetterFrequency(tableLetterFrequency)
}

func readFile(fileName string) string {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
	}
	defer file.Close()

	fileStat, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("Ошибка получения информации о файлe:", err)
	}

	dataInFile := make([]byte, fileStat.Size())
	resultString := ""

	for {
		n, err := file.Read(dataInFile)
		if err == io.EOF {
			break
		}
		resultString += string(dataInFile[:n])
	}

	return resultString
}

func createTableLetterFrequency(inputString string) []sortKey {
	alphabetRuneInText := createAlphabetForString(inputString)
	var tableLetterFrequency = make(map[rune]int, len(alphabetRuneInText))
	for i := 0; i < len(alphabetRuneInText); i++ {
		//Тут заполняем нашу таблицу где ключ runa а значение количество раз которых она встречается в тексте
		tableLetterFrequency[alphabetRuneInText[i]] = getLetterFrequency(alphabetRuneInText[i], inputString)
	}
	sortedTableLetterFrequency := sortMapByAscendingLetterFrequency(tableLetterFrequency)

	return sortedTableLetterFrequency
}

func createAlphabetForString(text string) []rune {
	// Сначала он проходит по всему тексту и смотрит какие rune вообще есть в этом тексте
	runeInText := make(map[rune]bool)
	for _, char := range text {
		runeInText[char] = true
	}
	// А потом он просто из отображения создает массив rune по этому тексту
	var arrayRuneInText []rune
	for char := range runeInText {
		arrayRuneInText = append(arrayRuneInText, char)
	}

	return arrayRuneInText
}

func getLetterFrequency(letter rune, inputString string) int {
	//Тут мы просто проходим по тексту и смотрим сколько раз наша rune встречается по тексту
	count := 0
	for _, char := range inputString {
		if char == letter {
			count++
		}
	}

	return count
}

func sortMapByAscendingLetterFrequency(inputMap map[rune]int) []sortKey {
	var arraySortKey []sortKey

	//Тут мы просто заполняем срез, данными из нашего Отображения, и получаем срез sortKey с нашими значениями
	for letterMap, countMap := range inputMap {
		arraySortKey = append(arraySortKey, sortKey{
			letter: letterMap,
			count:  countMap,
		})
	}

	//Тут обычная сортировка по убыванию
	sort.Slice(arraySortKey, func(i, j int) bool {
		return arraySortKey[i].count > arraySortKey[j].count
	})

	return arraySortKey
}

func createHuffmanThreeWithTableLetterFrequency(tableLetterFrequency []sortKey) {
	var huffmanThree []elementThree

	for _, element := range tableLetterFrequency {
		huffmanThree = append(huffmanThree, elementThree{
			element:    string(element.letter),
			count:      element.count,
			parentNode: nil,
			leftChild:  nil,
			rightChild: nil,
			code:       1,
		})
	}

	for len(huffmanThree) >= 2 {
		leftChild, rightChild := getTwoLastMin(huffmanThree)
		parentNode := getNewElement(&leftChild, &rightChild)
		leftChild.parentNode = &parentNode
		rightChild.parentNode = &parentNode
		leftChild.code = 1
		rightChild.code = 0

		huffmanThree = huffmanThree[:len(huffmanThree)-1]
		huffmanThree = huffmanThree[:len(huffmanThree)-1]

		huffmanThree = append(huffmanThree, parentNode)

		sort.Slice(huffmanThree, func(i, j int) bool {
			return huffmanThree[i].count > huffmanThree[j].count
		})
	}
	fmt.Println(huffmanThree[0].rightChild, huffmanThree[0].leftChild)
}

func getTwoLastMin(values []elementThree) (elementThree, elementThree) {
	min1, min2 := values[0], values[1]
	if min2.count < min1.count {
		min1, min2 = min2, min1
	}

	for i := 2; i < len(values); i++ {
		if values[i].count < min1.count {
			min2 = min1
			min1 = values[i]
		} else if values[i].count < min2.count {
			min2 = values[i]
		}
	}

	return min1, min2
}

func getNewElement(leftElement *elementThree, rightElement *elementThree) elementThree {
	return elementThree{
		element:    leftElement.element + rightElement.element,
		count:      leftElement.count + rightElement.count,
		parentNode: nil,
		leftChild:  leftElement,
		rightChild: rightElement,
		code:       1,
	}
}
