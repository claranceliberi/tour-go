package vin

import "fmt"


type VIN string

//constructor to validate the VIN

func NewVin(vin string) (VIN, error){

	if len(vin) != 17 {
		return "", fmt.Errorf("invalid VIN %s: more or less than 17 characters", vin)
	}

	return VIN(vin), nil
}

// returns the manufacturer of the VIN
func (v VIN) Manufacturer() string {
	return string(v[0:3])
}

type EUVIN VIN

func NewEUVIN(vin string) (EUVIN, error){
	
	v, err := NewVin(vin)

	if err != nil {
		return "", err
	}

	return EUVIN(v), nil
}

func (v EUVIN) Manufacturer() string {

	manufacture := VIN(v).Manufacturer();

	if manufacture[2] == '9' {
		manufacture += string(v[11:14])
	}

	return manufacture

}

