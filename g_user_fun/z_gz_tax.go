package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	CompanyEndowmentInsurance  float64 = 0.14
	PersonalEndowmentInsurance float64 = 0.08

	CompanyMedicalInsurance  float64 = 0.07
	PersonalMedicalInsurance float64 = 0.02

	CompanyUnemploymentInsurance  float64 = 0.0048
	PersonalUnemploymentInsurance float64 = 0.002

	StartTaxPrice float64 = 5000
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("Usage: go run z_gz_tax.go <base salary> <bonus>")
	}

	base, _ := strconv.ParseFloat(args[0], 64)
	bonus, _ := strconv.ParseFloat(args[1], 64)

	total := base + bonus
	// PB means Public, PV means Private
	var PBHouseFundPercent int
	var PVHouseFundPercent int
	var reliefAmount int

	fmt.Printf("company provident housing fund (range : 8 - 12):")
	_, _ = fmt.Scanln(&PBHouseFundPercent)
	fmt.Printf("personal house fund percent (range : 8 - 12):")
	_, _ = fmt.Scanln(&PVHouseFundPercent)

	fmt.Printf("relief : ")
	_, _ = fmt.Scanln(&reliefAmount)
	fmt.Printf("your info, PB %d,PV %d, relief %d\n", PBHouseFundPercent, PVHouseFundPercent, reliefAmount)

	fmt.Println("--------------Endowment insurance (养老保险)--------------------")
	fmt.Printf("Company  amount : %.2f\n", base*CompanyEndowmentInsurance)
	fmt.Printf("Personal amount : %.2f\n", base*PersonalEndowmentInsurance)
	fmt.Println("-----------------------------------------------------------------------")

	fmt.Println("--------------Medical insurance (医疗保险)--------------------——")
	fmt.Printf("Company  amount : %.2f\n", base*CompanyMedicalInsurance)
	fmt.Printf("Personal amount : %.2f\n", base*PersonalMedicalInsurance)
	fmt.Println("-----------------------------------------------------------------------")

	fmt.Println("--------------Unemployment insurance (失业保险)-----------------")
	fmt.Printf("Company  amount : %.2f\n", base*CompanyUnemploymentInsurance)
	fmt.Printf("Personal amount : %.2f\n", base*PersonalUnemploymentInsurance)
	fmt.Println("-----------------------------------------------------------------------")

	fmt.Println("--------------Housing fund (公积金)-----------------------------")
	fmt.Printf("Company  amount : %.2f\n", base*float64(PBHouseFundPercent)/100)
	fmt.Printf("Personal amount : %.2f\n", base*float64(PVHouseFundPercent)/100)
	fmt.Println("----------------------------------------------------------------------")

	percent, quickAmount := getTaxInfo(base + total)
	tax := (total-StartTaxPrice-base*PersonalEndowmentInsurance-base*PersonalMedicalInsurance-base*PersonalUnemploymentInsurance-base*float64(PVHouseFundPercent/100)-float64(reliefAmount))*percent - quickAmount
	fmt.Println("--------------Tax (个税)---------------------------------------")
	fmt.Printf("Amount : %.2f\n", tax)
	fmt.Println("----------------------------------------------------------------------")

	fmt.Println("--------------Got-----------------------------------------------------")
	fmt.Printf("Amount : %.2f\n", total-base*PersonalEndowmentInsurance-base*PersonalMedicalInsurance-base*PersonalUnemploymentInsurance-tax-base*float64(PVHouseFundPercent/100))
	fmt.Println("----------------------------------------------------------------------")
}

func getTaxInfo(total float64) (float64, float64) {
	if total <= 3000 {
		return 0.03, 0
	}
	if 3000 < total && total <= 12000 {
		return 0.1, 210
	}
	if 12000 < total && total <= 25000 {
		return 0.2, 1410
	}
	if 25000 < total && total <= 35000 {
		return 0.25, 2660
	}
	if 35000 < total && total <= 55000 {
		return 0.3, 4410
	}
	if 55000 < total && total <= 80000 {
		return 0.35, 7160
	}
	return 0.45, 15160
}
