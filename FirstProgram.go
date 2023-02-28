package main

import (
	"bufio"
	"fmt"
	"os"
)

func validate(title, description string, price, discount, categoryId, manofacturerId int, image string) bool {
	if len(title) != 0 && len(title) <= 100 {
		if len(description) != 0 && len(description) <= 1000 {
			if price >= 20 && price <= 50000 {
				if discount >= 0 && discount <= 95 {
					if categoryID >= 1 && categoryID <= 10 {
						if manufacturerID >= 1 && manufacturerID <= 50 {
							if len(image) <= 300 {
								return true
							}
						}
					}
				}
			}
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter product creation mode (single/multiple): ")
	scanner.Scan()
	productVolume := scanner.Text()
	if productVolume == "single" {
		fmt.Println("Title:")
		scanner.Scan()
		title := scanner.Text()
		fmt.Println()

		fmt.Println("Description:")
		scanner.Scan()
		description := scanner.Text()
		fmt.Println()

		fmt.Println("Price:")
		scanner.Scan()
		price := scanner.Text()
		fmt.Println()

		fmt.Println("Discount:")
		scanner.Scan()
		discount := scanner.Text()
		fmt.Println()

		fmt.Println("Category ID:")
		scanner.Scan()
		categoryId := scanner.Text()
		fmt.Println()

		fmt.Println("Manofacturer ID:")
		scanner.Scan()
		manofacturerId := scanner.Text()
		fmt.Println()

		fmt.Println("Image:")
		scanner.Scan()
		image := scanner.Text()
		fmt.Println()
	}
}
