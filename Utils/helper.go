package Utils

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"golang-watcher/KafkaProducer"

	"github.com/guylaor/goword"
	"github.com/ledongthuc/pdf"
)

//Logic for parsing CSV file which will be run as a goroutine
func ReadCsv(path string) {

	//fmt.Println("Reading CSV.....")
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// automatically call Close() at the end of current method
	defer file.Close()
	reader := csv.NewReader(file)
	lineCount := 0
	for {
		// read just one record, but we could ReadAll() as well
		record, err := reader.Read()
		// end-of-file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// record is an array of string so is directly printable
		fmt.Println("Record", lineCount, "is", record, "and has", len(record), "fields")
		// and we can iterate on top of that
		for i := 0; i < len(record); i++ {
			KafkaProducer.Produce(record[i])
		}
		fmt.Println()
		lineCount += 1
	}
}

//Logic for parsing text file which will be run as a goroutine
func ReadTxt(path string) {
	//fmt.Println("Reading txt......")
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// automatically call Close() at the end of current method
	defer file.Close()
	reader := bufio.NewReader(file)
	content, _ := ioutil.ReadAll(reader)

	fmt.Println(string(content))

}

//Logic for parsing docx file which will be run as a goroutine
func ReadDoc(path string) {
	//fmt.Println("Reading Docx.....")
	text, err := goword.ParseText(path)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%s ", text)
}

//Logic for parsing pdf file which will be run as a goroutine
func ReadPdf(path string) {
	//fmt.Println("Reading pdf ..........")
	content, err := readPdfUtil(path) // Read local pdf file
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}

func readPdfUtil(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}
