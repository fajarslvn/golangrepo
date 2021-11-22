/*
SETUP PACKAGES RELATIVE PATH (Terminal):
1. Make project folder on $GOPATH/src/<project-folder>.
2. cd to folder root <project-folder>.
3. Type: go mod init root/<project-folder> + Enter.
4. Import to main.go (import "root/project-folder/package")
5. If not work try type this: go env -w GO111MODULE=off
*/

package main

import (
	"crashcourse/10-packages/randoms"
	"crashcourse/10-packages/shapes"
	"fmt"
)

func main() {
	circle := shapes.Circle{Radius: 50.0}
	rectangle := shapes.Rectangle{Width: 45.0, Height: 45.0}

	fmt.Println("Circle area:", circle.Area())
	fmt.Println("Rectangle area:", rectangle.Area())
	fmt.Println("Rectangle is square:", rectangle.IsSquare())
	randoms.PrintRandom()
}