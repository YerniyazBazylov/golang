package employee_package3

type Employee3 struct {
    position string
    salary   float64
    address  string
}

func (e Employee3) GetPosition() string {
    return e.position
}

func (e Employee3) SetPosition(position string) {
    e.position = position
}

func (e Employee3) GetSalary() float64 {
    return e.salary
}

func (e Employee3) SetSalary(salary float64) {
    e.salary = salary
}

func (e Employee3) GetAddress() string {
    return e.address
}

func (e Employee3) SetAddress(address string) {
    e.address = address
}