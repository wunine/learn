package main

import (
	"fmt"
	"io"
	"text/template"

	"github.com/go-gomail/gomail"
)

type test struct {
	Column1 string
	Column2 string
	Column3 int
	Column4 float32
}

type Table struct {
	THead   []string
	TBody   []test
	Caption string
}

var head = []string{"column1", "column2", "column3", "column4"}

func GenTable() Table {
	var body []test
	for i := 0; i < 10; i++ {
		f := "value%d-%d"
		row := test{
			fmt.Sprintf(f, i, 1),
			fmt.Sprintf(f, i, 2),
			3,
			4.0,
		}
		body = append(body, row)
	}
	return Table{head, body, "this is caption"}
}

func ExecuteTpl(w io.Writer, data interface{}, filenames ...string) error {
	tmpl, err := template.ParseFiles(filenames...)
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}

func SendMail(d *gomail.Dialer, from string, to []string, data interface{}, tplnames ...string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", "test")
	m.AddAlternativeWriter("text/html", func(w io.Writer) error {
		return ExecuteTpl(w, data, tplnames...)
	})
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func main() {
	from := "1497298196@qq.com"
	to := []string{"1497298196@qq.com"}
	host := "smtp.qq.com"
	username, password := "1497298196@qq.com", "vjemifuptlvvieeg"
	tplnames := []string{"table.tpl"}
	table := GenTable()
	dailer := gomail.NewDialer(host, 25, username, password)
	err := SendMail(dailer, from, to, table, tplnames...)
	if err != nil {
		panic(err)
	}
}
