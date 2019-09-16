package saveload

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	_ "github.com/davecgh/go-spew/spew"
)

type SaveLoad struct {
	TableName string
	Fields    map[string]SaveLoadField
}

type SaveLoadField struct {
	Val func() string
	Set func(string)
}

func NewSaveLoad(tableName string) *SaveLoad {
	d := SaveLoad{TableName: tableName, Fields: make(map[string]SaveLoadField)}
	return &d
}

func (f *SaveLoad) Print() {
	fields := f.Fields
	for colname, slfield := range fields {
		v := slfield.Val()
		fmt.Printf("%s : %s\n", colname, v)
	}
}

func (f *SaveLoad) Save() {
	var cols []string
	var vals []string

	fields := f.Fields
	for colname, slfield := range fields {
		cols = append(cols, colname)
		vals = append(vals, slfield.Val())
	}

	fmt.Printf("INSERT INTO %s (%s) VALUES (%s)\n",
		f.TableName,
		strings.Join(cols[:], ","),
		strings.Join(vals[:], ","))
}

func (f *SaveLoad) String(name string, field *string) {
	v := func() string {
		if field == nil {
			return ""
		}
		return *field
	}

	s := func(fromSql string) {
		if field == nil {
			wtfToDoWithNilPointerError()
		} else {
			*field = fromSql
		}
	}

	m := SaveLoadField{Val: v, Set: s}
	f.Fields[name] = m
}

func (f *SaveLoad) Int(name string, field *int) {
	v := func() string {
		if field == nil {
			return "0"
		}
		s := strconv.Itoa(*field)
		return s
	}

	s := func(fromSql string) {
		if parsed, err := strconv.Atoi(fromSql); err == nil {
			if field == nil {
				wtfToDoWithNilPointerError()
			} else {
				*field = parsed
			}
		}
	}

	m := SaveLoadField{Val: v, Set: s}
	f.Fields[name] = m
}

func (f *SaveLoad) Time(name string, t *time.Time) {
	v := func() string {
		var timeStr string
		if t == nil {
			timeStr = time.Now().Format(time.RFC3339)
		} else {
			timeStr = t.Format(time.RFC3339)
		}
		return timeStr
	}

	s := func(fromSql string) {
		tFromDB, err := time.Parse(time.RFC3339, fromSql)
		if err != nil {
			if t == nil {
				wtfToDoWithNilPointerError()
			} else {
				*t = tFromDB
			}
		}
	}
	m := SaveLoadField{Val: v, Set: s}
	f.Fields[name] = m
}

// Allows the user to provide custom closures to get / set struct fields.
// The provided closures must save a pointer reference to the struct!
func (f *SaveLoad) Custom(name string, valFunc func() string, setFunc func(string)) {
	m := SaveLoadField{Val: valFunc, Set: setFunc}
	f.Fields[name] = m
}

func wtfToDoWithNilPointerError() {
	fmt.Println("not quite sure what to do here")
}
