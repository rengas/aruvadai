### var கேய்வோர்ட்  vs சுருக்கெழுத்து குறியீடு
நமது நிரலில் பின்வருமாரு ஒரு சிரிய திருத்தம் செய்வோம்.

```
package main

import (
	"fmt"
)

func main() {
    var v = "Kaniyam"
	fmt.Println(v)
	v = "Foundation"
	fmt.Println(v)
	v = "Kaniyam Foundation"
	fmt.Println(v)
}
```
```
renga@renga-Standard-PC-Q35-ICH9-2009:~/go-workspace/src/kaniyam$ go run main.go
Kaniyam  
Foundation
Kaniyam Foundation
```

இந்த நிரலுக்கும் நாம் சுருக்கெழுத்து குறியீடு பயன்படுத்தி எழுதிய நிரலுக்கும் உள்ள வேறுபாடு main 
ஃபங்சன்லில் உள்ள முதல் வரியில். இங்கு `:=` பதிலாக  `=` முதல் வரியில் பயன்படுத்தி உள்ளோம்
̀var கேய்வோர்ட் பயன்படுத்தி `v` ஒரு வேறியபிள் என்று தொகுப்பிக்கு குறுப்பிடுகிறொம். 
நாம் இரண்டு நாம் எவ்வகையான தரவு பயன் படுத்துகிறோம் என்று தொகுப்பிக்கு சுட்டிகாட்டவில்லை
இங்கும் go தொகுப்பியிடம் நாம் எவ்வகையான தரவு பயன் படுத்த வேண்டும்
என்பதை முடிவும் செயும் பொருப்பை கொடுக்கிறோம்.

## Scoping
நமது நிரலில் ஒரு சிரிய திருத்தம் செய்து var கேய்வோர்ட்க்கும் சுருக்கெழுத்து குறியீடுக்கும் 
இடையில் உள்ள வேறுபாட்டை அறிவோம். 
package main

```import (
	"fmt"
)

func main() {
    var v1 string = "Kaniyam Foundation"
	fmt.Println(v1)

	var v2 = "Kaniyam Foundation"
    fmt.Println(v2)

	var v3 string
	v3 = "Kaniyam Foundation"
}
```

இங்கு தொகுப்பிக்கு `v1` ` ஒரு string என்று தொகுப்பிக்கு குறவிட்டோம். 

