package main
import "fmt"

func main(){
	var users map[string]string
	users = make(map[string]string)

	users["mohsin"] = "blockchain"
	users["puneet"] = "angular"
	users["karthik"] = "frontend"
	users["gagan"] = "node sdk"
	
	for user := range users {
		fmt.Printf("%s %s\n",user,users[user])
	}

	delete(users,"mohsin")
	fmt.Printf("user deleted")
}