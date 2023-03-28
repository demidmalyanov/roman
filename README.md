## Roman numbers for Go

## Install

    $ go get github.com/demidmalyanov/roman

## Use

See it in action:

```go
package yours

import (
  "github.com/demidmalyanov/roman"
)

func compareRomans(r1, r2 string) {
	//convert
	val := romanToInt(r1)
	val2 := romanToInt(r2)

	if(val==val2){
		...
	}
    
    //compare
    if(compareTwoRomans(r1,r2)<1){
        ...
    }
    
    // using  viniculum system 
    val := vinculumRomanToInt(r1)
	val2 := vinculumRomanToInt(r2)
    
    
}

```
