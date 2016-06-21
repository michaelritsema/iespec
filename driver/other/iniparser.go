package main

import (
	"fmt"
	"iespec/stuff"
	"os"
)

func main() {
	dictFile, _ := os.Create("ziften-dictionary.go")
	dict := stuff.LoadINI("ziften_iespec.ini")

	header := "package iespec\n import (\n \"github.com/michaelritsema/ipfix\"\n)\n var MyFields = [...]ipfix.DictionaryEntry{"
	dictFile.WriteString(header)
	structTemplate := "\nipfix.DictionaryEntry{Name: \"%s\", EnterpriseID: %d, FieldID: %d, Type: ipfix.FieldTypes[\"%s\"],},"

	for k, v := range dict.Field {
		dictFile.WriteString(fmt.Sprintf(structTemplate, k, v.Enterprise, v.ID, v.Type))

		//fmt.Printf("%s", entries[i])

		//fmt.Printf("key: %s ID: %d EnterpriseID: %d Type: %s\n", k, v.ID, v.Enterprise, v.Type)
	}

	footer := "}"
	dictFile.WriteString(footer)

}
