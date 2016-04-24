//usr/bin/env go run $0 $@ ; exit

package main

import (
	"bytes"
	"github.com/ajstarks/svgo"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	var output bytes.Buffer

	canvas := svg.New(&output)
	canvas.Start(250, 250)
	canvas.Circle(125, 125, 100, "fill:#28262C;stroke:#5BA642;stroke-width:5px")
	canvas.Text(125, 135, "xn0px90!", "text-anchor:middle;font-size:36px;fill:#EB5633")
	canvas.End()

	js.Global.Get("document").Get("body").Set("innerHTML", output.String())
}
