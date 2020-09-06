/**
* Go语言词频统计，运行命令go run src/code/main.go test/words.txt
* @author unknown
* @since 2019-12-18
* 文件内容：
  hello tom glad to meet you
  yes me glad to meet you
  how are you
* 输出结果：
  Word        Frequency
  are   1
  glad  2
  hello 1
  how   1
  me    1
  meet  2
  to    2
  tom   1
  yes   1
  you   3
  Frequency → Words
  1 are, hello, how, me, tom, yes
  2 glad, meet, to
  3 you
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	pwd, _ := os.Getwd()
	filtName := "doc.txt"
	filePath := fmt.Sprintf("%s/word-frequency/%s", pwd, filtName)
	fmt.Println(filePath)
	//创建单词=>频次map
	frequencyForWord := map[string]int{} // 与:make(map[string]int)相同
	//读取文件内容
	for _, filename := range commandLineFiles([]string{filePath}) {
		//更新每个单词的频次
		updateFrequencies(filename, frequencyForWord)
	}
	//打印单词=>频次
	reportByWords(frequencyForWord)
	//反转单词=>频次 为 频次=>单词（多个）
	wordsForFrequency := invertStringIntMap(frequencyForWord)
	//打印频次=>单词（多个）
	reportByFrequency(wordsForFrequency)
}

func commandLineFiles(files []string) []string {
	/**
	 * 因为 Unix 类系统（如 Linux 或 Mac OS X 等）的命令行工具默认会自动处理通配符
	 *（也就是说，*.txt 能匹配任意后缀为 .txt 的文件，如 README.txt 和 INSTALL.txt 等），
	 * 而 Windows 平台的命令行工具（CMD）不支持通配符，所以如果用户在命令行输入 *.txt，那么程序只能接收到 *.txt。
	 * 为了保持平台之间的一致性，这里使用 commandLineFiles() 函数来实现跨平台的处理，当程序运行在 Windows 平台时，实现文件名通配功能
	 */
	if runtime.GOOS == "windows" {
		args := make([]string, 0, len(files))
		for _, name := range files {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name) // 无效模式
			} else if matches != nil {
				args = append(args, matches...)
			}
		}
		return args
	}
	return files
}

/**
 * updateFrequencies() 函数纯粹就是用来处理文件的，它打开给定的文件，并使用 defer 在函数返回时关闭文件，
 * 这里我们将文件作为一个 *bufio.Reader（使用 bufio.NewReader() 函数创建）传给 readAndUpdateFrequencies() 函数，
 * 因为这个函数是以字符串的形式一行一行地读取数据的，所以实际的工作都是在 readAndUpdateFrequencies() 函数里完成的
 */
func updateFrequencies(filename string, frequencyForWord map[string]int) {
	var file *os.File
	var err error
	if file, err = os.Open(filename); err != nil {
		log.Println("failed to open the file: ", err)
		return
	}
	defer file.Close()
	readAndUpdateFrequencies(bufio.NewReader(file), frequencyForWord)
}

/**
 * 第一部分的代码我们应该很熟悉了，用了一个无限循环来一行一行地读一个文件，
 * 当读到文件结尾或者出现错误的时候就退出循环，将错误报告给用户但并不退出程序，因为还有很多其他的文件需要去处理。
 * 任意一行都可能包括标点、数字、符号或者其他非单词字符，所以我们需要逐个单词地去读，
 * 将每一行分隔成若干个单词并使用 SplitOnNonLetters() 函数忽略掉非单词的字符，并且过滤掉字符串开头和结尾的空白。
 * 只需要记录含有两个以上（包括两个）字母的单词，可以通过使用 if 语句，如 utf8.RuneCountlnString(word) > 1 来完成。
 * 上面描述的 if 语句有一点性能损耗，因为它会分析整个单词，所以在这个程序里我们增加了一个判断条件，
 * 用来检査这个单词的字节数是否大于 utf8.UTFMax（utf8.UTFMax 是一个常量，值为 4，用来表示一个 UTF-8 字符最多需要几个字节）
 */
func readAndUpdateFrequencies(reader *bufio.Reader,
	frequencyForWord map[string]int) {
	for {
		line, err := reader.ReadString('\n')
		for _, word := range SplitOnNonLetters(strings.TrimSpace(line)) {
			if len(word) > utf8.UTFMax ||
				utf8.RuneCountInString(word) > 1 {
				frequencyForWord[strings.ToLower(word)] += 1
			}
		}
		if err != nil {
			if err != io.EOF {
				log.Println("failed to finish reading the file: ", err)
			}
			break
		}
	}
}

/**
 * 用来在非单词字符上对一个字符串进行切分，首先我们为 strings.FieldsFunc() 函数创建一个匿名函数 notALetter，
 * 如果传入的是字符那就返回 false，否则返回 true，
 * 然后返回调用函数 strings.FieldsFunc() 的结果，调用的时候将给定的字符串和 notALetter 作为它的参数
 */
func SplitOnNonLetters(s string) []string {
	notALetter := func(char rune) bool { return !unicode.IsLetter(char) }
	return strings.FieldsFunc(s, notALetter)
}

/**
 * 首先创建一个空的映射，用来保存反转的结果，但是我们并不知道它到底要保存多少个项，
 * 因此我们假设它和原来的映射容量一样大，然后简单地遍历原来的映射，将它的值作为键保存到反转的映射里，并将键增加到对应的值里去，
 * 新的映射的值就是一个字符串切片，即使原来的映射有多个键对应同一个值，也不会丢掉任何数据
 */
func invertStringIntMap(intForString map[string]int) map[int][]string {
	stringsForInt := make(map[int][]string, len(intForString))
	for key, value := range intForString {
		stringsForInt[value] = append(stringsForInt[value], key)
	}
	return stringsForInt
}

func reportByWords(frequencyForWord map[string]int) {
	words := make([]string, 0, len(frequencyForWord))
	wordWidth, frequencyWidth := 0, 0
	for word, frequency := range frequencyForWord {
		words = append(words, word)
		if width := utf8.RuneCountInString(word); width > wordWidth {
			wordWidth = width
		}
		if width := len(fmt.Sprint(frequency)); width > frequencyWidth {
			frequencyWidth = width
		}
	}
	sort.Strings(words)
	/**
	 * 经过排序之后我们打印两列标题，第一个是 "Word"，为了能让 Frequency 最后一个字符 y 右对齐，
	 * 需要在 "Word" 后打印一些空格，通过 %*s 可以实现的打印固定长度的空白，也可以使用 %s 来打印 strings.Repeat(" ", gap) 返回的字符串
	 */
	gap := wordWidth + frequencyWidth - len("Word") - len("Frequency")
	fmt.Printf("Word %*s%s\n", gap, " ", "Frequency")
	for _, word := range words {
		fmt.Printf("%-*s %*d\n", wordWidth, word, frequencyWidth,
			frequencyForWord[word])
	}
}

/**
 * 首先创建一个切片用来保存频率，并按照频率升序排列，然后再计算需要容纳的最大长度并以此作为第一列的宽度，之后输出报告的标题，
 * 最后，遍历输出所有的频率并按照字母升序输出对应的单词，如果一个频率有超过两个对应的单词则单词之间使用逗号分隔开
 */
func reportByFrequency(wordsForFrequency map[int][]string) {
	frequencies := make([]int, 0, len(wordsForFrequency))
	for frequency := range wordsForFrequency {
		frequencies = append(frequencies, frequency)
	}
	sort.Ints(frequencies)
	width := len(fmt.Sprint(frequencies[len(frequencies)-1]))
	fmt.Println("Frequency → Words")
	for _, frequency := range frequencies {
		words := wordsForFrequency[frequency]
		sort.Strings(words)
		fmt.Printf("%*d %s\n", width, frequency, strings.Join(words, ", "))
	}
}
