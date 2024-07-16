package main

import "fmt"

type Employee struct{
    ID int
    Name string
    Age int
    Sex string
}

type PEmployee struct{
    ID *int
    Name *string
    Age *int
    Sex *string
}

func convertPEmployee(employees []Employee) []PEmployee{

    var pEmployees []PEmployee
    for _,employee := range employees{
        pEmployees = append(pEmployees,PEmployee{
            ID: &employee.ID,
            Name: &employee.Name,
            Age: &employee.Age,
            Sex: &employee.Sex,
        })
    }
    return pEmployees

	    var pEmployees []PEmployee
        for _,employee := range employees{

        fmt.Println("==",i+1," loop==")
        fmt.Println("&i = ",&i)
        fmt.Println("&employee.ID = ",&employee.ID)
        fmt.Println("&employee.Name = ",&employee.Name)
        fmt.Println("&employee.Age = ",&employee.Age)
        fmt.Println("&employee.Sex = ",&employee.Sex)
        fmt.Println("&employee = ",employee)
        pEmployees = append(pEmployees,PEmployee{
            ID: &employee.ID,
            Name: &employee.Name,
            Age: &employee.Age,
            Sex: &employee.Sex,
        })
    }
    return pEmployees
}

func main(){

    employees := []Employee{
        Employee{
            ID:   1,
            Name: "山田太郎",
            Age:  23,
            Sex:  "男",
        },
        Employee{
            ID:   2,
            Name: "山田花子",
            Age:  27,
            Sex:  "女",
        },
        Employee{
            ID:   3,
            Name: "田中次郎",
            Age:  36,
            Sex:  "男",
        },
    }
    pEmployees := convertPEmployee(employees)

    fmt.Println("==PEmployees==")
    for _,pEmployee := range pEmployees{
        fmt.Println("ID = ",*pEmployee.ID)
        fmt.Println("Name = ",*pEmployee.Name)
        fmt.Println("Age = ",*pEmployee.Age)
        fmt.Println("Sex = ",*pEmployee.Sex)
        fmt.Println("--")
    }
}
