package motorbike

// Operate பயனிடலாரின் இன்புடை உள் வாங்கி, மோட்டார்சைக்கிள் உபயோக படுத்த உதவுகிறது  
 func Operate(accelerate int,stop bool, right float64, left float64)(bool success, err error) {
 	action := resolveAction(accelerate int,break bool, right float64, left float64)
	    if action =="" {
		   return false, errors.New("unable to resolve")
     }
     return process(action),nil
 }

 func resolveAction(accelerate int,stop bool, right float64, left float64) (string){
 	// inputs should be validated and return corresponding result
 }

 func process(action string ) (bool,error){
 	// do required calculations,
 }