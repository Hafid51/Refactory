package main

import "fmt"
import "time"

func main() { 
  
    fmt.Println("Masukkan Nama Restoran: ") 
    var restoName string // set tipe variabel
    fmt.Scanln(&restoName) // Get Data input restoran

    fmt.Println("Masukkan Nama Kasir: ") 
    var cashierName string // set tipe variabel
    fmt.Scanln(&cashierName) // Get Data input kasir  

    fmt.Println("Masukkan Nama Item: ") 
    var itemName string // set tipe variabel
    fmt.Scanln(&itemName) // Get Data input makanan

    fmt.Println("Masukkan Harga: ") 
    var price int // set tipe variabel
    fmt.Scanln(&price) // Get Data input Harga  

    fmt.Println("Masukkan Nama Item: ") 
    var itemName2 string // set tipe variabel
    fmt.Scanln(&itemName2) // Get Data input makanan

    fmt.Println("Masukkan Harga: ") 
    var price2 int // set tipe variabel
    fmt.Scanln(&price2) // Get Data input Harga  
 
  
    fmt.Print("\n")
    fmt.Printf("\t \t \t Warung Makan %v", restoName)
    
    fmt.Print("\n")
    fmt.Print("\n")
    fmt.Print("\n")
    today := time.Now()
	fmt.Print("\t Tanggal : \t", today.Format("02/01/2006"))
	fmt.Print("\t", today.Format("15:04:05"))

    fmt.Print("\n")
    fmt.Print("\n")
    fmt.Printf("\t Nama Kasir \t \t \t \t %v", cashierName)
    fmt.Print("\n")
    fmt.Print("\n")
    fmt.Print("\t ==============================================")
    fmt.Print("\n")
    fmt.Print("\n")
    fmt.Printf("\t %v \b ............................... \b %d", itemName, price)
    fmt.Print("\n")
    fmt.Print("\n")
    fmt.Printf("\t %v \b ............................... \b %d", itemName2, price2)
    fmt.Print("\n")
    fmt.Print("\n")
    fmt.Print("\n")
    fmt.Print("\n")
    fmt.Print("\n")
    fmt.Print("\t Total \b ................................. \b Rp. \b")
    fmt.Print(price + price2)
    fmt.Print("\n")
    fmt.Print("\n")

}