package main
import "fmt"
//Deferred function
var isConnected bool = false
func main() {
	fmt.Printf("Connection Open: %v\n", isConnected)
	databaseProcessing()
	fmt.Printf("Connection Open: %v\n", isConnected)
}
func databaseProcessing() {
	connect()
	fmt.Println("Deferring disconnect!")
	defer disconnect() //This one will be deferred until the end.
	fmt.Printf("Connection Open: %v\n", isConnected)
	fmt.Println("Doing Something")
}
func connect(){
	isConnected = true
	fmt.Println("Connected to Database!")
}
func disconnect(){
	isConnected = false
	fmt.Println("Disconnected from Database!")
}