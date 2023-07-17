package main

type ParkingSystem struct {
	CarType map[int]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		CarType: map[int]int{1: big, 2: medium, 3: small},
	}
}

func (p *ParkingSystem) AddCar(carType int) bool {
	if p.CarType[carType] > 0 {
		p.CarType[carType] -= 1
		return true
	}

	return false
}

func main() {

	var parkingSystem ParkingSystem
	parkingSystem.AddCar(1)

}
