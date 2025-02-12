```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	ID   int
	Name string
	Age  int
}

var Students []Student

func main() {
	ReadDataFromFile("students.txt")
	fmt.Println("Loaded students:")
	PrintStudents()

	fmt.Println("Adding new student...")
	AddStudent(Student{ID: 4, Name: "John", Age: 22})
	PrintStudents()

	fmt.Println("Updating student...")
	UpdateStudent(4, Student{ID: 4, Name: "John", Age: 23})
	PrintStudents()

	fmt.Println("Deleting student...")
	DeleteStudent(4)
	PrintStudents()
}

func ReadDataFromFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		parts := strings.Split(string(line), ",")
		id, _ := strconv.Atoi(parts[0])
		age, _ := strconv.Atoi(parts[2])

		Students = append(Students, Student{ID: id, Name: parts[1], Age: age})
	}
}

func PrintStudents() {
	for _, student := range Students {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", student.ID, student.Name, student.Age)
	}
}

func AddStudent(student Student) {
	Students = append(Students, student)
}

func UpdateStudent(id int, updatedStudent Student) {
	for i, student := range Students {
		if student.ID == id {
			Students[i] = updatedStudent
			return
		}
	}
}

func DeleteStudent(id int) {
	for i, student := range Students {
		if student.ID == id {
			Students = append(Students[:i], Students[i+1:]...)
			return
		}
	}
}
```
Цей код виконує базову обробку даних студентів, які зберігаються в текстовому файлі. Він зчитує дані з файлу, додає нового студента, оновлює інформацію про студента та видаляє студента. Код підтримує 150 рядків коду, якщо ви додасте більше функцій для обробки даних або складніших операцій з даними.