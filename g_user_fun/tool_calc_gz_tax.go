package main

import (
	"fmt"
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
	if len(args) < 2 {
		fmt.Println("Please input base info")
		return
	}

	baseWage, _ := strconv.ParseFloat(args[0], 64)
	additionalWage, _ := strconv.ParseFloat(args[1], 64)
	totalPackage := baseWage + additionalWage
	var companyHouseFundPercent int
	var personalHouseFundPercent int
	var reliefAmount float64

	fmt.Printf("Input your compaay house fund percent (range : 8 - 12%s) : ", "%")
	_, err1 := fmt.Scanln(&companyHouseFundPercent)
	if err1 != nil {
		fmt.Printf("Your input err")
		return
	}
	fmt.Printf("Input your personal house fund percent (range : 8 - 12%s) : ", "%")
	_, err0 := fmt.Scanln(&personalHouseFundPercent)
	if err0 != nil {
		fmt.Printf("Your input err")
		return
	}
	fmt.Printf("Input your relief amount : ")
	_, err2 := fmt.Scanln(&reliefAmount)
	if err2 != nil {
		fmt.Printf("Your input err")
		return
	}
	fmt.Println("--------------Endowment insurance monthly (养老保险)--------------------")
	fmt.Printf("Company  amount : %.2f\n", baseWage*CompanyEndowmentInsurance)
	fmt.Printf("Personal amount : %.2f\n", baseWage*PersonalEndowmentInsurance)
	fmt.Println("-----------------------------------------------------------------------")

	fmt.Println("--------------Medical insurance monthly (医疗保险)--------------------——")
	fmt.Printf("Company  amount : %.2f\n", baseWage*CompanyMedicalInsurance)
	fmt.Printf("Personal amount : %.2f\n", baseWage*PersonalMedicalInsurance)
	fmt.Println("-----------------------------------------------------------------------")

	fmt.Println("--------------Unemployment insurance monthly (失业保险)-----------------")
	fmt.Printf("Company  amount : %.2f\n", baseWage*CompanyUnemploymentInsurance)
	fmt.Printf("Personal amount : %.2f\n", baseWage*PersonalUnemploymentInsurance)
	fmt.Println("-----------------------------------------------------------------------")

	fmt.Println("--------------Housing fund monthly (公积金)-----------------------------")
	fmt.Printf("Company  amount : %.2f\n", baseWage*float64(companyHouseFundPercent)/100)
	fmt.Printf("Personal amount : %.2f\n", baseWage*float64(personalHouseFundPercent)/100)
	fmt.Println("----------------------------------------------------------------------")

	percent, quickAmount := getTaxInfo(baseWage + additionalWage)
	tax := (totalPackage-StartTaxPrice-baseWage*PersonalEndowmentInsurance-baseWage*PersonalMedicalInsurance-baseWage*PersonalUnemploymentInsurance-baseWage*float64(personalHouseFundPercent/100)-reliefAmount)*percent - quickAmount
	fmt.Println("--------------Tax monthly (个税)---------------------------------------")
	fmt.Printf("Amount : %.2f\n", tax)
	fmt.Println("----------------------------------------------------------------------")

	fmt.Println("--------------Got-----------------------------------------------------")
	fmt.Printf("Amount : %.2f\n", totalPackage-baseWage*PersonalEndowmentInsurance-baseWage*PersonalMedicalInsurance-baseWage*PersonalUnemploymentInsurance-tax-baseWage*float64(personalHouseFundPercent/100))
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
