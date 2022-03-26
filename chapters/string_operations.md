
```
package main

import (
	"fmt"
)

func main() {
	fmt.Println("1")
}
```
`இப்பொழுது கமாண்ட்  பிரொம்ப்டில் இருந்து  aruvadai ஃபோல்டர்க்கு சென்று  பின்வருபவற்றை பதிவு 
செய்யவும்`

```
renga@renga-Standard-PC-Q35-ICH9-2009:~/go-workspace/src/kaniyam$ go run main.go
1
```


அருமை இப்போது ஒண்றுடண் இன்னும் ஒரு ஒன்றை கூட்டி இரண்டு எண்று terminal லில் பதிவிடுவோம்

```
package main

import (
	"fmt"
)

func main() {
	fmt.Println("1"+"1")
}
```
```
renga@renga-Standard-PC-Q35-ICH9-2009:~/go-workspace/src/kaniyam$ go run main.go
11
```
11 என்பது தவறன பதில். இங்கு நாம் என்ன தவறு செய்து உள்ளேம் என்று யோசித்து பார்ப்போம். 
இங்கு நாம் இரண்டு ஒன்றுகளையிம் string தரவு வகை என்று சுட்டி காட்டியுள்ளோம். 
இங்கு மேற்க்கோள்குறி("") பயன்படுத்தி காட்டி go தொகுப்பிக்கு சுட்டுகிறோம்.
go தொகுப்பியை பொருத்தவரையில் இரண்டு string இடையில் '+' இருந்தால் 
அந்த இரண்டு string களை ஒன்றன் பின் ஒண்றக இனைக்க வேண்டும் என்று
முடிவு செய்யும்.  நாம் இங்கு தவறான தரவு வகை பயண் படுத்தியுள்ளோம்.
கணகிடுகள் செய்வதர்க்கு string தரவு வகை உகந்தல்ல