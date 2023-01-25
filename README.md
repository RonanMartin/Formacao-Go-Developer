# Formacao-Go-Developer

## **Go (Golang), a modern language, readable and with high capacity for multiple accesses, in addition to being extremely fast. This DIO course introduces you to your first steps in GO, the language of Google.**

## ###Conversion.go

This is a basic function:
Enter a value in Kelvin (_tempK_) and this value will be converted to Celsius (_TempC_)

## ###ProblemNUM.go

In part 1 of the challenge, the proposal was to create a code that displays all numbers, between 1 and 100, divisible by 3.
In the second part of the challenge, the proposal was to replace numbers that are multiples of 3 by "Pin" and multiples of 5 by "Pan".

## ###PingPong.go

In this challenge with channels and goroutines, the words Ping and Pong will be displayed alternately, just "go run".

## ###API-Rest.go

For this REST API I used the Gorilla MUX package. It is an agenda for entering customer data: Name, Address and Telephone.
Application where we can use POSTMAN to make requests:

- GET - http://localhost/clientes - From the function getClientes it will show all the clients in the directory.
- GET - http://localhost/clientes/{id} - From the getCliente function, it will only show the client with the corresponding ID.
- POST - http://localhost/clientes - From the createCliente function, you will insert the new Client in the Agenda, the ID will be automatically generated.
- PUT - http://localhost/clientes/{id} - From the updateCliente function, you will update a specific client.
- DELETE - http://localhost/clientes/{id} - From the function deleteCliente will delete the client of the respective ID.

## ###calc.go and calc-e-Test.go

_calc.go_:
In this challenge the objective was to simulate the basic functions of a calculator, add, subtract, multiply and divide:

- somar - All numbers entered in the function will be added;
- subtrair - From the first number inserted in the function, all others will be subtracted;
- multiplicar - All numbers will be multiplied;
- dividir - The first number will be divided by the next, and the result will be divided by the next, until the last.

_calc-e-Test.go_:
This is the unit tests of the calc.go function, where we can check if the functions bring the expected results.
