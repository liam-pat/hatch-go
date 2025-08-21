package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	// PB means Public, PV means Private
	PB_RETIREMENT_PERCENT float64 = 0.14
	PV_RETIREMENT_PERCENT float64 = 0.08

	PB_MEDICAL_PERCENT float64 = 0.07
	PV_MEDICAL_PERCENT float64 = 0.02

	PB_UNEMPLOYMENT_PERCENT float64 = 0.0048
	PV_UNEMPLOYMENT_PERCENT float64 = 0.002

	StartTaxPrice float64 = 5000
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("Usage: go run z_gz_tax.go <base salary> <bonus>")
	}

	base, _ := strconv.ParseFloat(args[0], 64)
	bonus, _ := strconv.ParseFloat(args[1], 64)

	var PBHouseFundPercent int
	var PVHouseFundPercent int
	var deduction int

	fmt.Printf("PB 公积金(8%% - 12%%):")
	_, _ = fmt.Scanln(&PBHouseFundPercent)
	fmt.Printf("PV 公积金(8%% - 12%%):")
	_, _ = fmt.Scanln(&PVHouseFundPercent)
	fmt.Printf("deduction: ")
	_, _ = fmt.Scanln(&deduction)

	fmt.Println("-----------------------------------------------------------------------")
	fmt.Printf("PB 公积金(%d%%), PV 公积金(%d%%), 减免(%d)\n", PBHouseFundPercent, PVHouseFundPercent, deduction)

	fmt.Println("-----------------------------------------------------------------------")
	fmt.Printf("%-40s: %.2f\n", "Retirement (养老金)", base*PB_RETIREMENT_PERCENT)
	fmt.Printf("%-40s: %.2f\n", "Medical (医疗金)", base*PB_MEDICAL_PERCENT)
	fmt.Printf("%-40s: %.2f\n", "Unemployment (失业金)", base*PB_UNEMPLOYMENT_PERCENT)
	fmt.Printf("%-40s: %.2f\n", "Housing fund (公积金)", base*float64(PBHouseFundPercent)/100)
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Printf("%-40s: %.2f\n", "Retirement (养老金)", base*PV_RETIREMENT_PERCENT)
	fmt.Printf("%-40s: %.2f\n", "Medical (医疗金)", base*PV_MEDICAL_PERCENT)
	fmt.Printf("%-40s: %.2f\n", "Unemployment (失业金)", base*PV_UNEMPLOYMENT_PERCENT)
	fmt.Printf("%-40s: %.2f\n", "Housing fund (公积金)", base*float64(PVHouseFundPercent)/100)
	fmt.Println("-----------------------------------------------------------------------")
	tax, gain := calcTax(base, base*PV_RETIREMENT_PERCENT, base*PV_MEDICAL_PERCENT, base*PV_UNEMPLOYMENT_PERCENT, base*float64(PVHouseFundPercent)/100, float64(deduction))
	fmt.Printf("%-40s: %.2f\n", "Tax    (个税金)", tax)
	fmt.Printf("%-40s: %.2f\n", "Result (应到手)", gain+bonus)
	fmt.Println("-----------------------------------------------------------------------")
}

func getTaxInfo(monthly float64) (float64, float64) {
	if monthly <= 5000 {
		return 0.0, 0
	}
	if 5000 < monthly && monthly <= 8000 {
		return 0.03, 0
	}
	if 8000 < monthly && monthly <= 17000 {
		return 0.10, 210
	}
	if 17000 < monthly && monthly <= 30000 {
		return 0.20, 1410
	}
	if 30000 < monthly && monthly <= 40000 {
		return 0.25, 2660
	}
	if 40000 < monthly && monthly <= 60000 {
		return 0.30, 4410
	}
	if 60000 < monthly && monthly <= 85000 {
		return 0.35, 7160
	}
	return 0.45, 15160
}

func calcTax(monthly float64, endowmentInsurance float64, medicalInsurance float64, unemploymentInsurance float64, houseFund float64, deduction float64) (float64, float64) {
	common := monthly - StartTaxPrice - endowmentInsurance - medicalInsurance - unemploymentInsurance - deduction

	percent, quickAmount := getTaxInfo(common)
	tax := common*percent - quickAmount

	gain := monthly - endowmentInsurance - medicalInsurance - unemploymentInsurance - houseFund - tax

	return tax, gain
}
