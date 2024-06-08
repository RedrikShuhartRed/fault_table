package text

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/RedrikShuhartRed/fault_table/models"
)

func CreateFault() (models.Fault, error) {
	var fault models.Fault
	reader := bufio.NewReader(os.Stdin)
	for {

		var err error

		fmt.Println("Введите номер турбины")
		fault.Turbine, err = reader.ReadString('\n')
		if err != nil {
			log.Printf("Ошибка ввода номера турбины: %s", err)
			return fault, err
		}

		fmt.Println("Введите дату ошибки в формате гггг.мм.дд")
		fault.Date, err = reader.ReadString('\n')
		if err != nil {
			log.Printf("Ошибка ввода даты: %s", err)
			return fault, err
		}

		fmt.Println("Введите код ошибки")
		fault.Code, err = reader.ReadString('\n')
		if err != nil {
			log.Printf("Ошибка ввода кода аварии: %s", err)
			return fault, err
		}

		fmt.Println("Введите описание ошибки")
		fault.Description, err = reader.ReadString('\n')
		if err != nil {
			log.Printf("Ошибка ввода описания аварии: %s", err)
			return fault, err
		}

		fault.Turbine = strings.TrimRight(fault.Turbine, "\r\n")
		fault.Date = strings.TrimRight(fault.Date, "\r\n")
		fault.Code = strings.TrimRight(fault.Code, "\r\n")
		fault.Description = strings.TrimRight(fault.Description, "\r\n")

		if fault.Turbine == "" || fault.Date == "" || fault.Code == "" || fault.Description == "" {
			fmt.Print("Неоходимо заполнить все поля\n")
			continue
		}
		break
	}

	return fault, nil
}

func EnterMain() (string, error) {
	var enter string
	fmt.Println(`Выберите пункт: 
1 - Добавить аварию в список, 
2 - Получить весь список аварий, 
3 - Получить список аварий одной турбины, 
4 - Получить список по определенной аварии.`)
	_, err := fmt.Scanln(&enter)
	if err != nil {
		log.Printf("Ошибка выбора пункта: %s", err)
		return "", err
	}
	return enter, nil
}

func EnterGetAll() (string, error) {
	var enter string
	for {
		fmt.Println(`
1 - Выгрузить по возрастанию даты аварии,
2 - Выгрузить по убыванию даты аварии,
3 - Выгрузить за определенный период,
4 - Сгруппировать выгрузку по номеру турбины,
5 - Сгруппировать выгрузку по коду аварии.`)
		_, err := fmt.Scanln(&enter)
		if err != nil {
			log.Printf("Ошибка выбора пункта: %s", err)
			return "", err
		}
		if enter != "1" && enter != "2" && enter != "3" && enter != "4" && enter != "5" {
			fmt.Println("\nВведите цифры от 1 до 5")
			continue
		}
		break
	}
	return enter, nil
}

func EnterGetByTurbine() (string, error) {
	fmt.Println("Введите номер турбины")
	reader := bufio.NewReader(os.Stdin)
	turbine, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Ошибка ввода номера турбины: %s", err)
		return "", err
	}
	turbine = strings.TrimRight(turbine, "\r\n")
	return turbine, nil
}

func EnterGetByFault() (string, error) {
	fmt.Println("Введите код аварии")
	reader := bufio.NewReader(os.Stdin)
	code, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Ошибка ввода кода аварии: %s", err)
		return "", err
	}
	code = strings.TrimRight(code, "\r\n")
	return code, nil
}

func GetBetweenDate() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите начало периода в формате гггг.мм.дд")

	begin, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Ошибка ввода начала периода: %s", err)
		return "", "", err
	}
	begin = strings.TrimRight(begin, "\r\n")

	fmt.Println("Введите окончание периода периода в формате гггг.мм.дд")
	end, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Ошибка ввода окончания периода: %s", err)
		return "", "", err
	}
	end = strings.TrimRight(end, "\r\n")
	return begin, end, nil
}
