package main

import "fmt"

type Student struct {
    name string
    grades []int
    age int
}

func (s *Student) getAge() int {
    return s.age
}

func (s *Student) setName(n string) {
    s.name = n
}

func (s Student) getAvgGrade() float32 {
    sum := 0
    for _, v := range s.grades {
        sum += v
    }
    return float32(sum) / float32(len(s.grades))
}

func main() {
    s1 := Student{"Jack", []int{70, 80, 85, 90}, 17}
    
    s1.getAge()
    s1.setName("Dude")
    fmt.Println(s1)
    
    avg := s1.getAvgGrade()
    fmt.Println(avg)
}
