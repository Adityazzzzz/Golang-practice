package main
import (
	"fmt"
	"time"
)
type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer
}

// method to change any value which is defined in struct
func (o *order) changeStatus(status string){
	o.status = status
}

func main(){
	myorder := order{
		id:"1",
		amount: 50.00,
		status: "received",
	}
	myorder.createdAt = time.Now()

	myorder.changeStatus("confirmed")//imp

	
	fmt.Println(myorder)
	fmt.Println(myorder.amount)
	fmt.Println(myorder.id)
}