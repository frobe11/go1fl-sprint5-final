package actioninfo

import (
	"fmt"
	"log"
)

// DataParser описывает общий сценарий обработки активности:
// разбор строки с данными и формирование строки с информацией.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info обрабатывает набор строк с данными об активностях:
// парсит каждую строку и выводит информацию о ней.
func Info(dataset []string, dp DataParser) {
	for _, datastring := range dataset {
		if err := dp.Parse(datastring); err != nil {
			log.Println(err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(info)
	}
}
