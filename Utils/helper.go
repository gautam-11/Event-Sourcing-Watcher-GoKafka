package Utils

import (
	"Event-Sourcing-Watcher-GoKafka/KafkaProducer"
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/guylaor/goword"
	"github.com/ledongthuc/pdf"
	"github.com/tealeg/xlsx"
)

var csvmutex sync.Mutex
var txtmutex sync.Mutex
var xlsxmutex sync.Mutex
var docmutex sync.Mutex

//ReadCsv : Logic for parsing CSV file which will be run as a goroutine
func ReadCsv(path string) {

	csvmutex.Lock()

	defer csvmutex.Unlock()
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
			KafkaProducer.Produce(record[i], "csv")
		}
		lineCount += 1
	}
}

//ReadTxt : Logic for parsing text file which will be run as a goroutine
func ReadTxt(path string) {

	txtmutex.Lock()

	defer txtmutex.Unlock()

	//fmt.Println("Reading txt......")
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// automatically call Close() at the end of current method
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		KafkaProducer.Produce(string(scanner.Text()), "txt")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

//ReadDoc : Logic for parsing docx file which will be run as a goroutine
func ReadDoc(path string) {

	docmutex.Lock()

	defer docmutex.Unlock()
	//fmt.Println("Reading Docx.....")
	text, err := goword.ParseText(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s ", text)
}

//ReadPdf : Logic for parsing pdf file which will be run as a goroutine
func ReadPdf(path string) {

	//fmt.Println("Reading pdf ..........")
	content, err := readPdfUtil(path) // Read local pdf file
	if err != nil {
		fmt.Println(err)
		return
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

//ReadXlsx : Logic for parsing xlsx file which will be run as a goroutine
func ReadXlsx(path string) {

	xlsxmutex.Lock()

	defer xlsxmutex.Unlock()

	xlFile, err := xlsx.OpenFile(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			var buffer bytes.Buffer
			for _, cell := range row.Cells {
				buffer.WriteString(cell.String())
				buffer.WriteString(" ")
			}
			KafkaProducer.Produce(buffer.String(), "xlsx")

			buffer.Reset()

		}
	}
}
