package main

import (
    "fmt"
)

type Student struct {
    name  []string
    score []int
}

func (s *Student) Add(name string, score int) {
    s.name = append(s.name, name)
    s.score = append(s.score, score)
}

func (s *Student) Average() float64 {
    total := 0
    for _, score := range s.score {
        total += score
    }
    return float64(total) / float64(len(s.score))
}

func (s *Student) Min() (min int, name string) {
    min = s.score[0]
    index := 0
    for i, score := range s.score {
        if score < min {
            min = score
            index = i
        }
    }
    return min, s.name[index]
}

func (s *Student) Max() (max int, name string) {
    max = s.score[0]
    index := 0
    for i, score := range s.score {
        if score > max {
            max = score
            index = i
        }
    }
    return max, s.name[index]
}

func main() {
    var students Student

    for i := 0; i < 5; i++ {
        var name string
        var score int

        fmt.Printf("Masukkan nama siswa ke-%d: ", i+1)
        fmt.Scan(&name)

        fmt.Printf("Masukkan skor siswa %s: ", name)
        fmt.Scan(&score)

        students.Add(name, score)
    }

    fmt.Println("Rata-rata skor:", students.Average())

    minScore, minName := students.Min()
    fmt.Printf("Siswa dengan skor minimum: %s (%d)\n", minName, minScore)

    maxScore, maxName := students.Max()
    fmt.Printf("Siswa dengan skor maksimum: %s (%d)\n", maxName, maxScore)
}
