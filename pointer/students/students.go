package main

import "fmt"

type Student struct {
    name  string
    score []int
}

func (s Student) AddScore(score int) Student {
    s.score = append(s.score, score)
    return s
}

func (s Student) AverageScore() float64 {
    var total int
    for _, score := range s.score {
        total += score
    }
    return float64(total) / float64(len(s.score))
}

func (s Student) MinScore() (string, int) {
    minScore := s.score[0]
    minName := s.name
    for _, score := range s.score {
        if score < minScore {
            minScore = score
            minName = s.name
        }
    }
    return minName, minScore
}

func (s Student) MaxScore() (string, int) {
    maxScore := s.score[0]
    maxName := s.name
    for _, score := range s.score {
        if score > maxScore {
            maxScore = score
            maxName = s.name
        }
    }
    return maxName, maxScore
}

func main() {
    var students []Student
fmt.Println("-------------------------")
    for i := 0; i < 5; i++ {
        fmt.Printf("Masukkan %d Nama Siswa: ", i+1)
        var name string
        fmt.Scan(&name)

        student := Student{name: name, score: []int{}}
        students = append(students, student)

        fmt.Printf("Input %d Nilai Siswa: ", i+1)
        var score int
        fmt.Scan(&score)

        students[i] = students[i].AddScore(score)
    }

    var totalScore int
    for _, student := range students {
        totalScore += student.score[0]
    }
    avgScore := float64(totalScore) / float64(len(students))

    var minName string
    var minScore int
    var maxName string
    var maxScore int
    for _, student := range students {
        name, score := student.MinScore()
        if score < minScore || minScore == 0 {
            minScore = score
            minName = name
        }
        name, score = student.MaxScore()
        if score > maxScore || maxScore == 0 {
            maxScore = score
            maxName = name
        }
    }

    fmt.Printf("Skor Rata-rata: %.0f\n", avgScore)
    fmt.Printf("Nilai Min Siswa: %s (%d)\n", minName, minScore)
    fmt.Printf("Nilai Maksimal Siswa: %s (%d)\n", maxName,maxScore)
fmt.Println("-------------------------")
}
