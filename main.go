package main
import 
(
	"fmt"
	"math"
)


// Create Heap Property 
func heapify(arr *[10] int,child int,n int ,i int ) {
	largest:=i
	for j:=1;j<=child;j++{  
      if child*i+j < n && arr[child*i+j]>arr[largest]{
		largest=child*i+j
	  }
	}
	if largest!=i{
		arr[largest],arr[i]=arr[i],arr[largest]
		fmt.Println(arr[largest], arr[i])
		heapify(arr,child,n,largest)
	}
  }

// main function where which the code starts 
func main(){
    arr:=[10]int {10,21,45,454,80,9,7,85,2,63}
	var x int 
	x=1
	child:=int(math.Pow(2,(float64(x))))
	for i := (len(arr)/ child)-1; i >= 0; i--{
		heapify(&arr,child,len(arr),i)
	}  
    fmt.Println(arr)
}