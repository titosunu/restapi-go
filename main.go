package main

import (
	"net/http"
	"rest-api/app"
	"rest-api/controller"
	"rest-api/helper"
	"rest-api/middleware"
	"rest-api/repository"
	"rest-api/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	validate := validator.New()
	db := app.NewDB()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

// package main

// import "fmt"

// type PaymentProcessor interface {
// 	ProcessPayment(amount float64) string
// }

// type Transaction struct {
// 	Processor string
// 	Amount float64
// 	Sender string
// }

// var transactionHistory []Transaction

// type ShopeePay struct {
// 	Number string
// }

// func (shopeePay ShopeePay) ProcessPayment(amount float64) string {
// 	transaction := Transaction{
// 		Processor: "SHOPEEPAY",
// 		Sender: shopeePay.Number,
// 		Amount: amount,
// 	}
// 	transactionHistory = append(transactionHistory, transaction)
// 	return fmt.Sprintf("Processing shopee pay. payment of Rp%.2f", amount)
// }

// type Gopay struct {
// 	Email string	
// }

// func (goPay Gopay) ProcessPayment(amount float64) string {
// 	transaction := Transaction{
// 		Processor: "GOPAY",
// 		Sender: goPay.Email,
// 		Amount: amount,
// 	}
// 	transactionHistory = append(transactionHistory, transaction)
// 	return fmt.Sprintf("Processing go pay. payment of Rp%.2f", amount)
// }

// func MakePayment(paymentProcessor PaymentProcessor, amount float64)  {
// 	fmt.Println(paymentProcessor.ProcessPayment(amount))
// }

// func DisplayHistory() {
// 	fmt.Println("TRANSACTION HISTORY")
// 	for _, v := range transactionHistory {
// 		fmt.Printf("%s: Rp%.2f, FROM %s\n", v.Processor, v.Amount, v.Sender)
// 	}
// }

// func main()  {
// 	person1 := ShopeePay{Number: "1221-082136160223"}
// 	person2 := Gopay{Email: "anonim@gmail.com"}
// 	MakePayment(person1, 20000)
// 	MakePayment(person2, 100000)
// 	DisplayHistory()
// }