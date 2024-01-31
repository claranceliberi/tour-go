package vin_test

import (
	"testing"
	"vin/vin"
)



const (
	validVIN = "W09000051T2123456"
	invalidVIN = "W0"
)

const euSmallVIN = "W09000051T2123456"


func TestVIN_New(t *testing.T){

	_, err := vin.NewVin(validVIN)

	if err != nil {
		t.Errorf("Creating a valid VIN should not return an error: %s", err.Error())
	}

	_, err = vin.NewVin(invalidVIN)

	if err == nil {
		t.Errorf("Creating an invalid VIN should return an error: %s", err.Error())
	}

}

func TestVin_Manufacturer(t *testing.T){

	vin, _ := vin.NewVin(validVIN)

	manufacture := vin.Manufacturer()

	if(manufacture != "W09123"){
		t.Errorf("Expected W09123, got %s", manufacture)
	}

}

func TestVIN_EU_SmallManufacturer(t *testing.T) {

	testVIN, _ := vin.NewEUVIN(euSmallVIN)
	manufacturer := testVIN.Manufacturer()
	if manufacturer != "W09123" {
	  t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
	}
  }

  // this fails with an error
func TestVIN_EU_SmallManufacturer_Polymorphism(t *testing.T) {

	var testVINs[] vin.VIN
	testVIN, _ := vin.NewEUVIN(euSmallVIN)
	// having to cast testVIN already hints something is odd
	testVINs = append(testVINs, vin.VIN(testVIN))
  
	for _, vin := range testVINs {
	  manufacturer := vin.Manufacturer()
	  if manufacturer != "W09123" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
	  }
	}
  }