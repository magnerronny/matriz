    

package main

import (
    "fmt"
    "sync"
    "time"
)


func ORDENAR(array []int, inicio int, fin int)[]int{


	for x := 0; x < fin; x++ {
		temp:=array[i]
		for y := 0; y < fin; y++ {
			if(temp<array[y]){

				temp=array[y]
				array[y]=array[x]
				array[x]=temp
			}
			
		}
	}
	array_div:=[]int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	
	var tem int
	for x := inicio; x < fin; x++ {
		tem=array[x]
		array_div[x] = tem
	}

	time.Sleep( 500*time.Millisecond)

	return array_div
}

func main() {
	array:=[]int{7,2,8,7,2,4,9,7,4,2,9,1,3,5,8,3}
	
	var wg sync.WaitGroup
		wg.Add(1)
	go func() {
		  fmt.Println("matriz 1 ",ORDENAR(array,0,4))
		  fmt.Println("matriz 2 ",ORDENAR(array,4,8))
		  fmt.Println("matriz 3 ",ORDENAR(array,8,12))
		  fmt.Println("matriz 4 ",ORDENAR(array,12,16))
		  fmt.Println("matriz general",ORDENAR(array,0,len(array)))
		wg.Done()
	}()

	wg.Wait()
 
}
