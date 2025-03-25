package reader

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"sync"
	"time"
)

// CSVReader 는 입력된 csv 형태 문자열을 Token 의 형태로 조회할 수 있도록 한다.
// 문자열 입력 --> 버퍼에 보관 --> Token 형태로 조회
type CSVReader struct {
	lineScanner *bufio.Scanner
	inputBuffer *bytes.Buffer
	readBuffer  *bytes.Buffer
	mu          sync.Mutex
}

func NewCSVReader() *CSVReader {
	ib := bytes.NewBuffer(make([]byte, 0, 1024))
	return &CSVReader{
		inputBuffer: ib,
		lineScanner: bufio.NewScanner(ib),
		readBuffer:  bytes.NewBuffer(make([]byte, 0, 1024)),
	}
}

func (c *CSVReader) InputStream(input io.Reader) {
	c.mu.Lock()
	defer c.mu.Unlock()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		c.inputBuffer.WriteString(scanner.Text() + "\n")
	}
}

func (c *CSVReader) LoadNextLine() bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.readBuffer.Reset()
	if c.lineScanner.Scan() {
		c.readBuffer.Write(c.lineScanner.Bytes())
		return true
	}
	return false
}

func (c *CSVReader) read(key any) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	idx, ok := key.(int)
	if !ok {
		return "", false
	}

	data := c.readBuffer.Bytes()
	start, cnt := 0, 0

	for i, b := range data {
		if b == ',' {
			if cnt == idx {
				return string(data[start:i]), true
			}
			start = i + 1
			cnt++
		}
	}

	if cnt == idx {
		return string(data[start:]), true
	}

	return "", false
}

func (c *CSVReader) StringGetter(idx any) func() (string, bool) {
	return func() (string, bool) {
		return c.read(idx)
	}
}

func (c *CSVReader) IntGetter(idx any) func() (int, bool) {
	return func() (int, bool) {
		str, ok := c.read(idx)
		if !ok {
			return 0, false
		}
		val, err := strconv.Atoi(str)
		return val, err == nil
	}
}

func (c *CSVReader) TimeGetter(idx any, layout string) func() (time.Time, bool) {
	return func() (time.Time, bool) {
		str, ok := c.read(idx)
		if !ok {
			return time.Time{}, false
		}
		t, err := time.Parse(layout, str)
		return t, err == nil
	}
}
