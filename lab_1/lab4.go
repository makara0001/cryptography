package main
import("fmt")

func Add(a ,b int) (int){
	return a + b
}

func Sub(a ,b int) (int){
	return a - b
}
func Mul(a ,b int) ( int){
	return a * b
}
func Div(a , b float64) (float64 ){
	if b== 0 {
		fmt.Printf("Infinte\n")
		return 0
	}
	return a / b 
}
func Mod(a , b int) (int){
	if b == 0 {
		fmt.Printf("Infinte\n")
		return 0
	}
	return a % b
}

func main  (){
	for {
		
		fmt.Printf("==== mini calculator =====\n")
		fmt.Printf("1).Add 2).Sub 3).Mul 4).Div 5).Mod 6).Exit\n")
		fmt.Printf("Choose: ")
	
		var choice int 
		fmt.Scanln(&choice)

		if choice == 6{
			fmt.Printf("Goodbye")
			break;
		}

		if choice != 1 && choice != 2 && choice != 3 && choice != 4 && choice !=5 {
			fmt.Printf("Choice Again\n")
			continue 
		}

	
		var num1, num2 int
		fmt.Printf("Enter the num1:")
		fmt.Scanln(&num1)
		fmt.Printf("Enter the num2:")
		fmt.Scanln(&num2)
	
		switch choice{
		case 1 :
			result := Add(num1,num2)
			fmt.Printf("The Result: %d + %d = %d\n",num1,num2,result)
		case 2 : 
			result := Sub(num1,num2)
			fmt.Printf("The Result: %d - %d = %d\n",num1,num2,result)
	    case 3 :
			result := Mul(num1,num2)
			fmt.Printf("The Result: %d * %d = %d\n",num1,num2,result)
		case 4 :
			result := Div(float64(num1),float64(num2))
			fmt.Printf("The Result: %d / %d = %.2f\n",num1,num2,result)

		case 5 :
			result := Mod(num1,num2)
			fmt.Printf("The Result: %d %% %d = %d\n",num1,num2,result)

	
		} 
	}



}