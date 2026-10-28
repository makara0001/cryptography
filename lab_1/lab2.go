package main
import(
	"fmt"
)
// func person ()int {
// 	boy := 18
// 	girl := 18
// 	if boy == 18 && girl == 18 {
// 		fmt.Printf("there are adult \n")
// 		return 1


// 	} else {
// 		fmt.Printf("there are underage\n")
// 		return 0
// 	}

// }
// func Older ()int {
// 	boy := 19
// 	girl := 18
// 	if boy > girl || girl > boy {

// 		if boy > girl {
// 			fmt.Printf(" the boy is older \n")
// 		    return 1
// 		} else {
// 			fmt.Printf("the girl is older\n")
// 			return 0
// 		}
		


// 	} else {
// 		fmt.Printf("they are the same\n")
// 		return -1
// 	}

// }

// func Gender () int {

// 	person1 := "girl"
// 	person2 := "girl"
	
// 	if person1 == "boy"|| person2 == "girl"{

// 		if person1 == "boy" {
// 			fmt.Printf("person 1 : boy\n")
		
// 		} else {
// 			fmt.Printf("person 1 : girl\n")
			
// 		} 
// 		if person2 == "girl" {
// 			fmt.Printf("person2: girl\n")
// 			return 1
// 		} else {
// 			fmt.Printf("person 2:boy")
// 			return 0
// 		}
// 	} else {
// 		fmt.Printf("undefine")
// 		return -1
// 	}
// }
// func main(){
// 	person()
// 	Older()
// 	Gender()

// }
// func Operation() (int,int){
// 	var n1 int
// 	var n2 int

// 	fmt.Printf("Enter the n1:")
// 	fmt.Scanln(&n1)

// 	fmt.Printf("Enter the n2:")
// 	fmt.Scanln(&n2)

// 	if n1 > 0 && n2 > 0 {
// 		fmt.Printf("both are positive\n")
		
// 	} else {
// 		fmt.Printf("At least one number is not positive\n")
	    
// 	}

// 	if n1 > n2 {
// 		fmt.Printf("n1 is greater than n2\n")
		 
// 	} else  n2 > n1 {
// 		fmt.Printf("n2 is greater than n1\n")
		
// 	} 

// 	if n1 != n2 {
// 		fmt.Printf("both are not equal\n")
// 		return n1,n2
// 	} else {
// 		fmt.Printf("both are equal\n ")
// 		return 1,1
// 	}


// }


// func main(){
// 	Operation()
	
// }

func Operation() (bool ,bool ,bool){
	var n1 int
	var n2 int 
	fmt.Printf("Enter the n1:")
	fmt.Scanln(&n1)
	fmt.Printf("Enter the n2:")
	fmt.Scanln(&n2)
	

	positive := n1 > 0 && n2 > 0

	greater := n1 > n2 || n2 > n1

	notEqual := !( n1==n2)

    return positive,greater,notEqual


}

func main(){
	bothpositive, onegreater, therenotEqual := Operation()
	fmt.Printf("both positive: %t\n", bothpositive)
	fmt.Printf("one is greater:%t\n",onegreater)
	fmt.Printf("not equal:%t\n",therenotEqual)

}