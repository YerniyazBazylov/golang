package employee_package1

type Employee1 struct {
    position string
    salary   float64
    address  string
}

func (e Employee1) GetPosition() string {
    return e.position
}

func (e Employee1) SetPosition(position string) {
    e.position = position
}

func (e Employee1) GetSalary() float64 {
    return e.salary
}

func (e Employee1) SetSalary(salary float64) {
    e.salary = salary
}

func (e Employee1) GetAddress() string {
    return e.address
}

func (e Employee1) SetAddress(address string) {
    e.address = address
}