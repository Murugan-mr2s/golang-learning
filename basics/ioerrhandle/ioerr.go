package main

import ( "fmt"
		  "os"	
		  "log"
		  "errors"	
)


func listDir(dpath string ) [] string {

	entries, err :=os.ReadDir(dpath)
	if err!=nil {
		log.Fatal(err)
	}
	
	names := [] string {}
	for _,file := range entries {
		names = append(names, file.Name()) 
	}
	return names
}


func printFiles( names [] string ) {
	for _,name := range names {
		fmt.Println(name)
	}
}


func readData(name string) {

	bytes,err :=os.ReadFile(name)
	if err !=nil {
		log.Fatal("File Error:" , err)
	}
	os.Stdout.WriteString( string(bytes) )
}

func writeData(name string,data string) {
	err:= os.WriteFile(name, []byte(data) , 0666)
	if err !=nil {
		log.Fatal("file error:", err)
	}
}

func writeFile(name string, data string) {

	file ,err :=os.Create(name)
	if err!=nil {
		log.Fatal("file error :", err)
	}

	nb,err :=file.Write(  []byte(data) )
	if err!=nil {
		log.Fatal("file readd error :", err)
	}
	os.Stdout.WriteString( string(nb) )
	defer file.Close()

}


func readFile(name string) {
	
	file, err :=os.Open(name)
	if err!=nil {
		log.Fatal("file open error :" , err)
	}
	defer file.Close()
	b:=make ([]byte, 3)
	var n int =0

	for ; err==nil; {
		n,err=file.Read(b)
		os.Stdout.WriteString(string(b[:n]))
	} 
}

func divide(a,b int) (int, error) {

	if b==0 {
		return 0,errors.New("It rise divide by zero")
	}
	return  a/b,nil
} 

func divide2(a,b float64) (float64, error) {
	
		if b==0 {
			return 0,fmt.Errorf("It rise divide by zero %f/%f",a,b)
		}
		return  a/b,nil
	} 


func DzeroDetector(a,b float64) float64  {

	defer func() {
		if err:=recover(); err !=nil {
			fmt.Println("catch panic:", err)
		} 
	}()

	if b==0 {
		panic (" rise divide by zero")
	}
	return a/b 
}

type DivisionError struct {
	divident float64
	divisor  float64
}

func (e *DivisionError) Error() string {
	return fmt.Sprintf("Division  not allowed %f / %f" ,e.divident,e.divisor)
}

func divide3(a,b float64) (float64,error)  {
	if b==0 {
		return 0,&DivisionError{a,b}	
	} 
	return a/b , nil	
}

func main() {

/*	
	fmt.Println("IO Error Handling")
	names :=listDir("E:\\golearn")
	printFiles(names)

	//writeData("data\\test2.txt","hello world 123")
	//readData("test.txt")

	writeFile("data\\test2.txt","hello world 123")
	readFile("data\\test2.txt")
*/
	/*
	_,er:=divide(1000,200)
	if er!=nil {
		fmt.Println("Err:",er)
	} 

	_,er2:=divide(1000,0)
	if er2!=nil {
		fmt.Println("Err:",er2)
	} 
	*/

	/*	
	_,er2:=divide2(1000,0)
	if er2!=nil {
		fmt.Println("Err:",er2)
	}
	*/

	/*
	DzeroDetector(1000,0)
	fmt.Println("panic restored!")
   */

   _, err :=divide3(2000,0)
   if(err !=nil) {
		fmt.Println(err)
   }

   fmt.Println("finished")
}	