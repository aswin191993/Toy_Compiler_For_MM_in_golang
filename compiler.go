package main

import "fmt"

var goval int =0
var ival int =0
var strval [5] int

func findvalue(x string){
	for j:=0; j<len(x);j++{	
		printfun(string(x[j]))
	}
	for k:=0; k<len(x);k++{	
		printfuns(string(x[k]))
	}		
		
}

func printfun(p string){
	acode:=make(map[string]string)
	acode["("]=""
	acode[")"]=""
	acode["a"]="ld"
	acode["b"]="ld"
	acode["c"]="ld"
	acode["d"]="ld"
	for k, v := range acode {
    		if(p==k){
			if(v=="ld"){
				gotostr(k,v)
			}
		}
	}
}

func printfuns(p string){
	acode:=make(map[string]string)
	acode["/"]="div"
	acode["*"]="mul"
	acode["+"]="add"
	acode["-"]="sub"
	for k, v := range acode {
    		if(p==k){
			gotosyb(v)
		}
	}
}

func gotosyb(syid string){
	if(strval[1]!=0){
		fmt.Printf("  %s r%d, r%d\n",syid,strval[0],strval[ival+1])
	}
	ival = ival+1
}

func gotostr(s,id string){
	strval[goval]=goval
	fmt.Printf("  %s r%d, %s\n",id,goval,s)
	goval=goval+1
}

func main(){	
	fmt.Println("\n\n#####COMPILER#####\n")	
	k := "(a/b)+c"
	fmt.Println("inputed code:", k,"\n\n")
	findvalue(k)
	fmt.Println("  hlt")	
}

