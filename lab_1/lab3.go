
package main
import ("fmt") 

func Bitwise() (int,int,int,int){
	var x int
	var y int
	fmt.Printf("Enter the x:")
	fmt.Scanln(&x)
	fmt.Printf("Enter the y:")
	fmt.Scanln(&y)



	myXor := x ^ y
	myNot := ^x ^ y
	myOR := x | y
	myAND := x & y



	return myXor, myNot,myOR,myAND


	
}

func main (){
	XoR,Not,OR,And:= Bitwise()
	fmt.Printf("the result of myXoR:%d\n", XoR)
	fmt.Printf("the result of myNot:%d\n", Not)
	fmt.Printf("the result of myOR:%d\n", OR)
	fmt.Printf("the result of myAND:%d\n", And)

}