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

func (s *Student) getAvgGrade() float32 {
    sum := 0
    
    for _, v := range s.grades {
        sum += v
    }
    return float32(sum) / float32(len(s.grades))
}

func (s *Student) getMaxGrade() int {
    curMax := 0
    
    for _, v := range s.grades {
        if curMax < v {
            curMax = v
        }
    }
    return curMax
}

func main() {
    s1 := Student{"Jack", []int{70, 80, 85, 90, 98}, 21}
    s2 := Student{"Jack", []int{75, 83, 75, 92, 100, 80}, 19}
    
    s1.getAge()
    s1.setName("Dude")
    
    s2.getAge()
    s2.setName("Joe")
    
    fmt.Println("First Student:", s1, "Second Student:", s2)
    
    avg1 := s1.getAvgGrade()
    avg2 := s2.getAvgGrade()
    
    fmt.Println("First Student:", avg1, "Second Student:", avg2)
    
    max1 := s1.getMaxGrade()
    max2 := s2.getMaxGrade()
    fmt.Println("First Student:", max1, "Second Student:", max2)
}
