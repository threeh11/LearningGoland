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
	letter string
	count  int
	parentNode,
	leftChild,
	rightChild *elementThree
	code      byte
	isParent  bool
	isChild   bool
	isChecked bool
}

func main() {
	resultAfterReadFile := readFile("test2.txt")
	tableLetterFrequency := createTableLetterFrequency(resultAfterReadFile)
	resultThreeHuffman := createHuffmanThreeWithTableLetterFrequency(tableLetterFrequency)
	getAlphabetForEncodingFile(resultThreeHuffman[0], tableLetterFrequency)
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

func createHuffmanThreeWithTableLetterFrequency(tableLetterFrequency []sortKey) []elementThree {
	var huffmanThree []elementThree

	for _, element := range tableLetterFrequency {
		huffmanThree = append(huffmanThree, elementThree{
			letter:     string(element.letter),
			count:      element.count,
			parentNode: nil,
			leftChild:  nil,
			rightChild: nil,
			code:       1,
			isChild:    true,
		})
	}

	for len(huffmanThree) > 1 {
		var lIndex, rIndex int
		leftChild, lIndex, rightChild, rIndex := getTwoLastMin(huffmanThree)
		parentNode := getNewElement(&leftChild, &rightChild)
		leftChild.parentNode = &parentNode
		rightChild.parentNode = &parentNode
		leftChild.code = 1
		rightChild.code = 0

		if len(huffmanThree) >= 3 {
			huffmanThree = append(huffmanThree[:lIndex], huffmanThree[lIndex+1:]...)
			if rIndex == 0 {
				rIndex = 1
			}
			huffmanThree = append(huffmanThree[:rIndex-1], huffmanThree[rIndex:]...)
		} else {
			huffmanThree = huffmanThree[:len(huffmanThree)-1]
			huffmanThree = huffmanThree[:len(huffmanThree)-1]
		}

		huffmanThree = append(huffmanThree, parentNode)
	}

	return huffmanThree
}

func getTwoLastMin(values []elementThree) (elementThree, int, elementThree, int) {
	min1, min2 := values[0], values[1]
	min1Index := 0
	min2Index := 1
	if min2.count < min1.count {
		min1, min2 = min2, min1
		min1Index, min2Index = min2Index, min1Index
	}

	for i := 2; i < len(values); i++ {
		if values[i].count < min1.count {
			min2 = min1
			min1 = values[i]
			min1Index = i
		} else if values[i].count < min2.count {
			min2 = values[i]
			min2Index = i
		}
	}

	return min1, min1Index, min2, min2Index
}

func getNewElement(leftElement *elementThree, rightElement *elementThree) elementThree {
	return elementThree{
		letter:     leftElement.letter + rightElement.letter,
		count:      leftElement.count + rightElement.count,
		parentNode: nil,
		leftChild:  leftElement,
		rightChild: rightElement,
		code:       1,
		isParent:   true,
		isChecked:  false,
	}
}

func getAlphabetForEncodingFile(topThree elementThree, tableLetterFrequency []sortKey) {
	resultAlphabet := make(map[string]string)
	for _, element := range tableLetterFrequency {
		fmt.Print(element)
		//resultAlphabet[string(element.letter)] = findLetterInThree(topThree, string(element.letter))
	}
	fmt.Print(resultAlphabet)
}

func findLetterInThree(currentTarget elementThree, letter string) string {
	if currentTarget.isChild && currentTarget.letter == letter {
		return getCodeByPosition(currentTarget)
	} else {
		if !currentTarget.leftChild.isChecked {
			currentTarget = *currentTarget.leftChild
		} else if !currentTarget.rightChild.isChecked {
			currentTarget = *currentTarget.rightChild
		} else {
			currentTarget = *currentTarget.parentNode
		}
		findLetterInThree(currentTarget, letter)
	}
	return ""
}

func getCodeByPosition(currentElement elementThree) string {
	resultCode := string(currentElement.code)
	for currentElement.parentNode != nil {
		resultCode += string(currentElement.code)
		currentElement = *currentElement.parentNode
	}
	return resultCode
}
