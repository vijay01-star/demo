package main
import ("fmt")
//create struct
type Person struct {
  name string
  age int
  job string
  salary int
}
func add(a,b int)int{
 return a + b
}
<<<<<<< HEAD
// func sub(a,b int)int{
//  re1:=a-b
//   return re1
// } 
=======
func sub(a,b int)int{
 re1:=a-b
  return re1
}
>>>>>>> 713925a7283734c7653e62981e3375f4df6e82b1

func main() {

  fmt.Println("Addition: ", add(10, 5),"\n","Subtraction: ", sub(10, 5))
 // fmt.Println("Subtraction: ", sub(10, 5))
  var pers1 Person//struct variable
  var pers2 Person//sruct variable

  // Pers1 specification
  pers1.name = "vijay"
  pers1.age = 22
  pers1.job = "learner"
  pers1.salary = 28000

  // Pers2 specification
  pers2.name = "kumar"
  pers2.age = 22
  pers2.job = "Trainee"
  pers2.salary = 22500

  // Access and print Pers1 info
  fmt.Println("person1 Name: ", pers1.name,"\n","person2 Name: ",pers2.name)
  fmt.Println("person1 Age: ", pers1.age,"\n","person2 age: ",pers2.age)
  fmt.Println("person1 Job: ", pers1.job,"\n","person2 Job: ",pers2.job)
  fmt.Println("person1 Salary: ", pers1.salary,"\n","person2 Salary: ",pers2.salary)

  // Access and print Pers2 info
  // fmt.Println("Name: ", pers2.name)
  // fmt.Println("Age: ", pers2.age)
  // fmt.Println("Job: ", pers2.job)
  // fmt.Println("Salary: ", pers2.salary)
 
}

     

