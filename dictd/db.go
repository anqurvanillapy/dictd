package dictd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Response message to the client
type Response = string

// Response message for status
const (
	StatusOK   = "ok"
	StatusFail = "bad"
)

var db = make(map[string]string)

var cmdInfo = map[byte]int{
	'+': 2,
	'=': 1,
	'-': 1,
}

// Cmd is the database command
type Cmd interface {
	String() string
	Execute() Response
}

// ParseMsg parses the request and gives out the DB command
func ParseMsg(msg string) Cmd {
	if msg == "" {
		return nil
	}

	cmd := msg[0]
	argStr := msg[1:]

	if argStr == "" {
		switch cmd {
		case '*':
			return &All{}
		case '!':
			return &Clr{}
		case '?':
			return &Siz{}
		}
	}

	args := strings.Split(argStr, "\r\n")

	argc, ok := cmdInfo[cmd]
	if !ok || len(args) != argc {
		return nil
	}

	switch cmd {
	case '+':
		return &Set{args[0], args[1]}
	case '=':
		return &Get{args[0]}
	case '-':
		return &Del{args[0]}
	}

	return nil
}

// Set the key-value pair
type Set struct {
	Key string
	Val string
}

func (c *Set) String() string {
	return fmt.Sprintf("Set(key=%q, val=%q)", c.Key, c.Val)
}

// Execute Set command
func (c *Set) Execute() Response {
	db[c.Key] = c.Val
	return StatusOK
}

// Get the value by key
type Get struct {
	Key string
}

func (c *Get) String() string {
	return fmt.Sprintf("Get(key=%q)", c.Key)
}

// Execute Get command
func (c *Get) Execute() Response {
	if v, ok := db[c.Key]; ok {
		return fmt.Sprintf("%q", v)
	}
	return StatusFail
}

// Del to delete value by key
type Del struct {
	Key string
}

func (c *Del) String() string {
	return fmt.Sprintf("Del(key=%q)", c.Key)
}

// Execute Del command
func (c *Del) Execute() Response {
	delete(db, c.Key)
	return StatusOK
}

// All to dump database
type All struct{}

func (All) String() string {
	return "All"
}

// Execute All command
func (All) Execute() Response {
	res, err := json.Marshal(db)
	if err != nil {
		log.Println(err)
	}
	return string(res)
}

// Clr to clear database
type Clr struct{}

func (Clr) String() string {
	return "Clr"
}

// Execute Clr command
func (Clr) Execute() Response {
	db = make(map[string]string)
	return StatusOK
}

// Siz to get number of pairs
type Siz struct{}

func (Siz) String() string {
	return "Siz"
}

// Execute Siz command
func (Siz) Execute() Response {
	return strconv.Itoa(len(db))
}
