# domain-credential-checker

A Go application for authenticating users against a network share. This tool allows you to test multiple user credentials from a CSV file and generate detailed logs for both successful and unsuccessful attempts.

## Features

- **Batch Processing**: Test a list of user credentials from a CSV file.
- **Detailed Logging**: Separate logs for successful and failed authentication attempts.
- **Easy Setup**: Simple to compile and run with just a few commands.

## Running the Program

Follow these steps to get started:

1. **Save the Source Code**  
   Save the provided source code into a file named `main.go`.

2. **Open Terminal**  
   Open your terminal and navigate to the directory containing `main.go`.

3. **Compile the Program**  
   Compile the code into an executable using the following command:

   ```sh
   go build -o network_connect main.go

4. **Run the Program**
Execute the compiled program with:

   ```sh
   ./network_connect

5. **Prepare Your CSV File**
Ensure your CSV file contains user logins in the correct format. The program will process these logins using the password you provide.

**Usage**

* **CSV Format**: The CSV file should contain one user login per line in the following format: username@domain.
* **Password**: The password you provide will be used for all login attempts.






# RUSSIAN


**domain-credential-checker**
Приложение на Go для аутентификации пользователей в сетевой папке. Этот инструмент позволяет протестировать несколько учетных записей из файла CSV и создавать подробные логи успешных и неудачных попыток подключения.

**Возможности**
* Пакетная обработка: Проверьте список учетных данных из файла CSV.
* Подробное логирование: Отдельные логи для успешных и неудачных попыток аутентификации.
* Простота настройки: Легко компилируется и запускается с помощью нескольких команд.

**Запуск программы**
Следуйте этим шагам, чтобы начать работу:

1. Сохраните исходный код в файл с именем main.go.

2. Откройте терминал и перейдите в каталог, содержащий main.go.

3. Скомпилируйте код в исполняемый файл с помощью следующей команды:

   ```sh
   go build -o network_connect main.go
   
4. Запустите программу
Выполните скомпилированную программу командой:

   ```sh
   ./network_connect
5. Подготовьте ваш CSV-файл
Убедитесь, что ваш CSV-файл содержит логины пользователей в правильном формате. Программа будет обрабатывать эти логины с использованием указанного вами пароля.

**Использование**
Формат CSV: Файл CSV должен содержать один логин пользователя на строку в следующем формате: username@domain.
Пароль: Пароль, который вы укажете, будет использоваться для всех попыток входа.
