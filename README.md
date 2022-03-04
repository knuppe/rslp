# RSLP

A [Go (golang)](http://golang.org) implementation of the Stemming Algorithm for the Portuguese Language for natural language processing.

This package uses the algorithm described in [this article](http://doi.ieeecomputersociety.org/10.1109/SPIRE.2001.10024) by Viviane Moreira Orengo and Christian Huyck.

![Schema](https://raw.githubusercontent.com/knuppe/rslp/main/steps.png)


## Usage


Here is a minimal Go program that uses this package in order
to stem a single word **or** a sentence.

```go
package main
import (
	"fmt"
	"github.com/knuppe/rslp"
)

func main(){
    // single word
	stemmed := rslp.Stem("cantárei")
	fmt.Println(stemmed) // Prints "cant"

    // sentence
    stemmed = rslp.StemSentence("Que você compartilhe livremente, nunca recebendo mais do que você dá.")
    fmt.Println(stemmed) // Prints "que voc compartilh livremente, nunc receb mais do que voc da."
}
```


## License (MIT)

Copyright (c) 2022 Gustavo Knuppe

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
