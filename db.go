package nanodb

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"
)

type DB struct {
	path   string
	opened bool
	sync.Mutex
	file *os.File
}

func Open(path string) (*DB, error) {
	db := &DB{opened: true}
	db.path = path
	f, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	db.file = f

	return db, nil
}

func (db *DB) Put(key, value string) error {
	db.Lock()
	defer db.Unlock()
	check := db.Get(key)
	if check != "" {
		// fmt.Println("Reached key already found")
		err := db.delete(key)
		if err != nil {
			log.Fatal(err)
		}

	}
	key = url.QueryEscape(key)
	value = url.QueryEscape(value)
	s := fmt.Sprintf("%s:%s\n", key, value)
	_, err := db.file.Write([]byte(s))
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) Get(key string) string {
	db.file.Seek(0, 0)
	var value string
	sc := bufio.NewScanner(db.file)
	for sc.Scan() {
		text := sc.Text()
		// fmt.Println(text)
		splitText := strings.Split(text, ":")
		keyText, _ := url.QueryUnescape(splitText[0])
		if keyText == key {
			value, _ = url.QueryUnescape(splitText[1])
			break
		}
	}
	return value
}
func (db *DB) delete(key string) error {
	db.file.Seek(0, 0)

	var bs []byte
	buf := bytes.NewBuffer(bs)

	sc := bufio.NewScanner(db.file)
	for sc.Scan() {
		text := sc.Text()
		// fmt.Println(text)
		splitText := strings.Split(text, ":")
		keyText, _ := url.QueryUnescape(splitText[0])
		if keyText != key {
			_, err := buf.Write(sc.Bytes())
			if err != nil {
				return err
			}
			_, err = buf.WriteString("\n")
			if err != nil {
				return err
			}
		}
	}

	err := os.WriteFile(db.path, buf.Bytes(), 0666)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) Delete(key string) error {
	db.Lock()
	defer db.Unlock()
	db.file.Seek(0, 0)

	var bs []byte
	buf := bytes.NewBuffer(bs)

	sc := bufio.NewScanner(db.file)
	for sc.Scan() {
		text := sc.Text()
		// fmt.Println(text)
		splitText := strings.Split(text, ":")
		keyText, _ := url.QueryUnescape(splitText[0])
		if keyText != key {
			_, err := buf.Write(sc.Bytes())
			if err != nil {
				return err
			}
			_, err = buf.WriteString("\n")
			if err != nil {
				return err
			}
		}
	}

	err := os.WriteFile(db.path, buf.Bytes(), 0666)
	if err != nil {
		return err
	}
	return nil
}
