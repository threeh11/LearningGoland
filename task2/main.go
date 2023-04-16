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
	element    sortKey
	parentNode *int
	code       byte
}

func main() {
	resultAfterReadFile := readFile("test.txt")
	tableLetterFrequency := createTableLetterFrequency(resultAfterReadFile)
	createHaffmanThreeWithTableLetterFrequency(tableLetterFrequency)
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

func createHaffmanThreeWithTableLetterFrequency(tableLetterFrequency []sortKey) map[rune]byte {

}
