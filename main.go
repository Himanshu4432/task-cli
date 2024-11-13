package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "time"
)

// Task represents a to-do item
type Task struct {
    ID          int       `json:"id"`
    Description string    `json:"description"`
    Status      string    `json:"status"` // todo, in-progress, done
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

// File where tasks are stored
const taskFile = "tasks.json"

// loadTasks loads tasks from the JSON file
func loadTasks() ([]Task, error) {
    if _, err := os.Stat(taskFile); os.IsNotExist(err) {
        // If file doesn't exist, return empty task list
        return []Task{}, nil
    }

    data, err := ioutil.ReadFile(taskFile)
    if err != nil {
        return nil, err
    }

    var tasks []Task
    if len(data) == 0 {
        return []Task{}, nil
    }

    err = json.Unmarshal(data, &tasks)
    if err != nil {
        return nil, err
    }

    return tasks, nil
}

// saveTasks saves tasks to the JSON file
func saveTasks(tasks []Task) error {
    data, err := json.MarshalIndent(tasks, "", "  ")
    if err != nil {
        return err
    }

    return ioutil.WriteFile(taskFile, data, 0644)
}

// getNextID returns the next task ID
func getNextID(tasks []Task) int {
    maxID := 0
    for _, task := range tasks {
        if task.ID > maxID {
            maxID = task.ID
        }
    }
    return maxID + 1
}

// Add a new task
func addTask(description string) {
    tasks, err := loadTasks()
    if err != nil {
        fmt.Println("Error loading tasks:", err)
        return
    }

    newTask := Task{
        ID:          getNextID(tasks),
        Description: description,
        Status:      "todo",
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }

    tasks = append(tasks, newTask)

    err = saveTasks(tasks)
    if err != nil {
        fmt.Println("Error saving task:", err)
        return
    }

    fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
}

// List tasks with optional filter
func listTasks(filter string) {
    tasks, err := loadTasks()
    if err != nil {
        fmt.Println("Error loading tasks:", err)
        return
    }

    if filter != "" {
        var filteredTasks []Task
        for _, task := range tasks {
            if task.Status == filter {
                filteredTasks = append(filteredTasks, task)
            }
        }
        tasks = filteredTasks
    }

    if len(tasks) == 0 {
        fmt.Println("No tasks found.")
        return
    }

    for _, task := range tasks {
        fmt.Printf("ID: %d\nDescription: %s\nStatus: %s\nCreated At: %s\nUpdated At: %s\n\n",
            task.ID,
            task.Description,
            task.Status,
            task.CreatedAt.Format(time.RFC1123),
            task.UpdatedAt.Format(time.RFC1123))
    }
}

// Update a task's description
func updateTask(id int, newDescription string) {
    tasks, err := loadTasks()
    if err != nil {
        fmt.Println("Error loading tasks:", err)
        return
    }

    found := false
    for i, task := range tasks {
        if task.ID == id {
            tasks[i].Description = newDescription
            tasks[i].UpdatedAt = time.Now()
            found = true
            break
        }
    }

    if !found {
        fmt.Printf("Task with ID %d not found.\n", id)
        return
    }

    err = saveTasks(tasks)
    if err != nil {
        fmt.Println("Error saving tasks:", err)
        return
    }

    fmt.Println("Task updated successfully.")
}

// Delete a task by ID
func deleteTask(id int) {
    tasks, err := loadTasks()
    if err != nil {
        fmt.Println("Error loading tasks:", err)
        return
    }

    index := -1
    for i, task := range tasks {
        if task.ID == id {
            index = i
            break
        }
    }

    if index == -1 {
        fmt.Printf("Task with ID %d not found.\n", id)
        return
    }

    tasks = append(tasks[:index], tasks[index+1:]...)

    err = saveTasks(tasks)
    if err != nil {
        fmt.Println("Error saving tasks:", err)
        return
    }

    fmt.Println("Task deleted successfully.")
}

// Mark a task as in-progress or done
func markTask(id int, status string) {
    if status != "in-progress" && status != "done" {
        fmt.Println("Invalid status. Use 'in-progress' or 'done'.")
        return
    }

    tasks, err := loadTasks()
    if err != nil {
        fmt.Println("Error loading tasks:", err)
        return
    }

    found := false
    for i, task := range tasks {
        if task.ID == id {
            tasks[i].Status = status
            tasks[i].UpdatedAt = time.Now()
            found = true
            break
        }
    }

    if !found {
        fmt.Printf("Task with ID %d not found.\n", id)
        return
    }

    err = saveTasks(tasks)
    if err != nil {
        fmt.Println("Error saving tasks:", err)
        return
    }

    fmt.Printf("Task marked as %s successfully.\n", status)
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide a command.")
        fmt.Println("Available commands: add, list, update, delete, mark-in-progress, mark-done")
        return
    }

    command := os.Args[1]

    switch command {
    case "add":
        if len(os.Args) < 3 {
            fmt.Println("Please provide a description for the task.")
            return
        }
        description := os.Args[2]
        addTask(description)

    case "list":
        if len(os.Args) == 3 {
            filter := os.Args[2]
            if filter != "todo" && filter != "in-progress" && filter != "done" {
                fmt.Println("Invalid filter. Use 'todo', 'in-progress', or 'done'.")
                return
            }
            listTasks(filter)
        } else {
            listTasks("")
        }

    case "update":
        if len(os.Args) < 4 {
            fmt.Println("Usage: task-cli update <id> <new description>")
            return
        }
        id, err := strconv.Atoi(os.Args[2])
        if err != nil {
            fmt.Println("Invalid task ID. It should be a number.")
            return
        }
        newDescription := os.Args[3]
        updateTask(id, newDescription)

    case "delete":
        if len(os.Args) < 3 {
            fmt.Println("Usage: task-cli delete <id>")
            return
        }
        id, err := strconv.Atoi(os.Args[2])
        if err != nil {
            fmt.Println("Invalid task ID. It should be a number.")
            return
        }
        deleteTask(id)

    case "mark-in-progress":
        if len(os.Args) < 3 {
            fmt.Println("Usage: task-cli mark-in-progress <id>")
            return
        }
        id, err := strconv.Atoi(os.Args[2])
        if err != nil {
            fmt.Println("Invalid task ID. It should be a number.")
            return
        }
        markTask(id, "in-progress")

    case "mark-done":
        if len(os.Args) < 3 {
            fmt.Println("Usage: task-cli mark-done <id>")
            return
        }
        id, err := strconv.Atoi(os.Args[2])
        if err != nil {
            fmt.Println("Invalid task ID. It should be a number.")
            return
        }
        markTask(id, "done")

    default:
        fmt.Println("Unknown command:", command)
        fmt.Println("Available commands: add, list, update, delete, mark-in-progress, mark-done")
    }
}
