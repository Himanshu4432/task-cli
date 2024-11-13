# Task Tracker CLI

![Task Tracker](https://img.shields.io/badge/Go-1.20-blue.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Building the Application](#building-the-application)
- [Usage](#usage)
  - [Adding a Task](#adding-a-task)
  - [Listing Tasks](#listing-tasks)
  - [Updating a Task](#updating-a-task)
  - [Deleting a Task](#deleting-a-task)
  - [Marking a Task as In Progress or Done](#marking-a-task-as-in-progress-or-done)
- [Configuration](#configuration)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [License](#license)

---

## Introduction

**Task Tracker CLI** is a simple and efficient command-line application built in Go for managing your to-do list. It allows you to add, update, delete, and track the status of your tasks seamlessly from the terminal. Whether you're a developer, student, or anyone looking to organize tasks, this CLI tool is designed to help you stay productive.

---

## Features

- **Add Tasks**: Quickly add new tasks with descriptions.
- **Update Tasks**: Modify existing task descriptions.
- **Delete Tasks**: Remove tasks that are no longer needed.
- **Mark Status**: Change task status to `todo`, `in-progress`, or `done`.
- **List Tasks**: View all tasks or filter them based on their status.
- **Persistent Storage**: Tasks are stored in a `tasks.json` file, ensuring data persistence between sessions.

---

## Prerequisites

Before you begin, ensure you have met the following requirements:

- **Go Installed**: Version 1.20 or later. You can download it from [Go's official website](https://golang.org/dl/).
- **Git Bash or PowerShell**: For running commands on Windows.
- **A Code Editor**: Such as VSCode, GoLand, or any editor of your choice.

---

## Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/task-cli.git
   cd task-cli
   ```

2. **Alternatively, Download the Source Code**

   You can download the source code as a ZIP file from the repository and extract it to your desired location.

---

## Building the Application

### 1. **Navigate to the Project Directory**

   Open your terminal (Git Bash or PowerShell) and navigate to the `task-cli` directory:

   ```bash
   cd "C:\Users\YourName\task-cli"   # For PowerShell
   # or
   cd ~/task-cli                      # For Git Bash
   ```

### 2. **Build the Executable**

   Use the `go build` command to compile the application:

   ```bash
   go build -o task-cli.exe
   ```

   **Explanation:**
   - `go build`: Compiles the Go source files.
   - `-o task-cli.exe`: Specifies the output file name as `task-cli.exe`.

### 3. **Verify the Build**

   Ensure that the `task-cli.exe` file has been created in your project directory:

   ```bash
   ls
   ```

   **Expected Output:**

   ```
   main.go
   task-cli.exe
   tasks.json   # This file will be created when you add your first task
   ```

---

## Usage

After successfully building the application, you can start using the `task-cli` to manage your tasks.

### **Running the Executable**

- **In Git Bash:**

  ```bash
  ./task-cli.exe [command] [arguments]
  ```

- **In PowerShell:**

  ```powershell
  .\task-cli.exe [command] [arguments]
  ```

### **Adding an Alias (Optional for Git Bash Users)**

To simplify command execution in Git Bash, you can create an alias:

1. **Open the `.bashrc` File:**

   ```bash
   nano ~/.bashrc
   ```

2. **Add the Alias:**

   ```bash
   alias task-cli='./task-cli.exe'
   ```

3. **Save and Exit:**

   - Press `Ctrl + O` to save.
   - Press `Enter` to confirm.
   - Press `Ctrl + X` to exit.

4. **Reload the Configuration:**

   ```bash
   source ~/.bashrc
   ```

5. **Use the Alias:**

   ```bash
   task-cli add "Buy groceries"
   ```

### **Commands and Their Descriptions**

#### **1. Add a New Task**

**Command:**

```bash
task-cli add "Task Description"
```

**Example:**

```bash
./task-cli.exe add "Buy groceries"    # Git Bash
.\task-cli.exe add "Buy groceries"     # PowerShell
```

**Output:**

```
Task added successfully (ID: 1)
```

#### **2. List All Tasks**

**Command:**

```bash
task-cli list
```

**Example:**

```bash
./task-cli.exe list    # Git Bash
.\task-cli.exe list     # PowerShell
```

**Output:**

```
ID: 1
Description: Buy groceries
Status: todo
Created At: Wed, 13 Nov 2024 10:00:00 UTC
Updated At: Wed, 13 Nov 2024 10:00:00 UTC
```

#### **3. List Tasks by Status**

**Commands:**

- **List Done Tasks:**

  ```bash
  task-cli list done
  ```

- **List To-Do Tasks:**

  ```bash
  task-cli list todo
  ```

- **List In-Progress Tasks:**

  ```bash
  task-cli list in-progress
  ```

**Example:**

```bash
./task-cli.exe list done             # Git Bash
.\task-cli.exe list in-progress      # PowerShell
```

#### **4. Update a Task**

**Command:**

```bash
task-cli update <id> "New Description"
```

**Example:**

```bash
./task-cli.exe update 1 "Buy groceries and cook dinner"    # Git Bash
.\task-cli.exe update 1 "Buy groceries and cook dinner"     # PowerShell
```

**Output:**

```
Task updated successfully.
```

#### **5. Delete a Task**

**Command:**

```bash
task-cli delete <id>
```

**Example:**

```bash
./task-cli.exe delete 1    # Git Bash
.\task-cli.exe delete 1     # PowerShell
```

**Output:**

```
Task deleted successfully.
```

#### **6. Mark a Task as In Progress**

**Command:**

```bash
task-cli mark-in-progress <id>
```

**Example:**

```bash
./task-cli.exe mark-in-progress 1    # Git Bash
.\task-cli.exe mark-in-progress 1     # PowerShell
```

**Output:**

```
Task marked as in-progress successfully.
```

#### **7. Mark a Task as Done**

**Command:**

```bash
task-cli mark-done <id>
```

**Example:**

```bash
./task-cli.exe mark-done 1    # Git Bash
.\task-cli.exe mark-done 1     # PowerShell
```

**Output:**

```
Task marked as done successfully.
```

---

## Configuration

### **tasks.json**

- **Location:** The `tasks.json` file is created in the project directory (`task-cli`) and is used to store all tasks.
- **Structure:**

  ```json
  [
    {
      "id": 1,
      "description": "Buy groceries",
      "status": "todo",
      "created_at": "2024-11-13T10:00:00Z",
      "updated_at": "2024-11-13T10:00:00Z"
    },
    {
      "id": 2,
      "description": "Complete the Go CLI project",
      "status": "in-progress",
      "created_at": "2024-11-14T09:30:00Z",
      "updated_at": "2024-11-14T09:45:00Z"
    }
  ]
  ```

- **Note:** Avoid manually editing this file to prevent data corruption. Use the CLI commands to manage tasks.

---

## Troubleshooting

### **1. Executable Not Recognized**

**Error:**

```
bash: task-cli: command not found
```

**Solution:**

- **Ensure You're Using the Correct Syntax:**

  - **Git Bash:**

    ```bash
    ./task-cli.exe add "Buy groceries"
    ```

  - **PowerShell:**

    ```powershell
    .\task-cli.exe add "Buy groceries"
    ```

- **Check If the Executable Exists:**

  ```bash
  ls
  ```

  Ensure `task-cli.exe` is listed.

### **2. Build Errors Due to Unused Imports**

**Error:**

```
.\main.go:6:5: "fmt" imported and not used
.\main.go:9:5: "strconv" imported and not used
```

**Solution:**

- **Use All Imported Packages:**

  Ensure that all imported packages in `main.go` are utilized in your code. For example, use `fmt` for printing and `strconv` for string conversions.

- **Remove Unused Imports:**

  If certain packages are not needed, remove them from the import list.

### **3. Permission Issues**

**Symptoms:**

- Unable to execute the binary.
- Errors related to file permissions.

**Solution:**

- **Run as Administrator:**

  Try running Git Bash or PowerShell as an administrator.

- **Check Antivirus Settings:**

  Ensure that your antivirus software is not blocking `task-cli.exe`.

### **4. JSON File Issues**

**Symptoms:**

- Tasks not being saved or loaded correctly.
- Corrupted `tasks.json` file.

**Solution:**

- **Delete `tasks.json`:**

  If the file is corrupted, delete it. The application will recreate it when you add a new task.

- **Ensure Read/Write Permissions:**

  Verify that you have the necessary permissions to read from and write to the `tasks.json` file.

---

## Contributing

Contributions are welcome! If you have suggestions for improvements or encounter any issues, feel free to open an [issue](https://github.com/yourusername/task-cli/issues) or submit a [pull request](https://github.com/yourusername/task-cli/pulls).

### **Steps to Contribute:**

1. **Fork the Repository**
2. **Clone Your Fork:**

   ```bash
   git clone https://github.com/Himanshukr4432/task-cli.git
   cd task-cli
   ```

3. **Create a New Branch:**

   ```bash
   git checkout -b feature/YourFeatureName
   ```

4. **Make Your Changes**

5. **Commit Your Changes:**

   ```bash
   git commit -m "Add new feature"
   ```

6. **Push to Your Fork:**

   ```bash
   git push origin feature/YourFeatureName
   ```

7. **Open a Pull Request**

---

## License

This project is licensed under the [MIT License](LICENSE).

---

## Acknowledgements

- Built with ❤️ using Go.
- Inspired by the need for simple and effective task management tools.

---

## Contact

For any queries or feedback, feel free to reach out:

- **Email:** your.email@example.com
- **GitHub:** [himanshukr4432](https://github.com/himanshukr4432)

---

**Happy Task Tracking! 🚀**