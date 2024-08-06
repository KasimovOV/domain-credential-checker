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
