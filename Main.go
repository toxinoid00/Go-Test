package main
import "fmt"
import "strconv"
import "strings"

type Node struct {
	data 			byte
	children 	[]Node
}

type Tree struct {
	root Node
}

func (node *Node) AddChildren(item Node) {
	node.children = append(node.children, item)
}

func main() {
	//Generate Tree
	//Create Node Tree
	A := Node{
		data: 'A',
	}
	B := Node{
		data: 'B',
	}
	C := Node{
		data: 'C',
	}
	D := Node{
		data: 'D',
	}
	E := Node{
		data: 'E',
	}
	F := Node{
		data: 'F',
	}
	G := Node{
		data: 'G',
	}
	H := Node{
		data: 'H',
	}

	//Add Children to each Node
	B.AddChildren(D)
	B.AddChildren(E)
	C.AddChildren(F)
	C.AddChildren(G)
	C.AddChildren(H)
	A.AddChildren(B)
	A.AddChildren(C)

	tree := Tree{root: A}
	
	//Test No 1
	exNumberOne1 := 20
	exNumberOne2 := 50
	fmt.Println()
	fmt.Println("Test Jawaban No 1")
	fmt.Println("Input:", exNumberOne1, " Output:", GetDonutsSlice(exNumberOne1))
	fmt.Println("Input:", exNumberOne2, " Output:", GetDonutsSlice(exNumberOne2))
	fmt.Println()

	//Test No 2
	exNumberTwo := 7
	fmt.Println("Test Jawaban No 2")
	fmt.Println("Input", exNumberTwo, " Output:", GetRowsValue(exNumberTwo))
	fmt.Println()

	//Test No 3
	fmt.Println("Test Jawaban No 3")
	PrintNodeTree("", &tree.root)
	fmt.Println()

	exNumberFour1 := "A-C-F"
	exNumberFour1Arr := strings.Split(exNumberFour1, "-")

	exNumberFour2 := "A-B-D"
	exNumberFour2Arr := strings.Split(exNumberFour2, "-")

	exNumberFour3 := "A-B-F"
	exNumberFour3Arr := strings.Split(exNumberFour3, "-")

	//Test No 4
	fmt.Println("Test Jawaban No 4")
	fmt.Println("Input: " + exNumberFour1 + " Output:", IsValidParent(exNumberFour1Arr, len(exNumberFour1Arr), 0, &tree.root))
	fmt.Println("Input: " + exNumberFour2 + " Output:", IsValidParent(exNumberFour2Arr, len(exNumberFour2Arr), 0, &tree.root))
	fmt.Println("Input: " + exNumberFour3 + " Output:", IsValidParent(exNumberFour3Arr, len(exNumberFour3Arr), 0, &tree.root))
	fmt.Println()

	exNumberFive1 := "A-C"
	exNumberFive1Arr := strings.Split(exNumberFive1, "-")

	exNumberFive2 := "A-B-D"
	exNumberFive2Arr := strings.Split(exNumberFive2, "-")

	//Test No 5
	fmt.Println("Test Jawaban No 5")
	fmt.Println("Input: " + exNumberFive1 + " Output:", GetNodeChildren(exNumberFive1Arr, len(exNumberFive1Arr), 0, &tree.root))
	fmt.Println("Input: " + exNumberFive2 + " Output:", GetNodeChildren(exNumberFive2Arr, len(exNumberFive2Arr), 0, &tree.root))
	fmt.Println()

	//For Custom Input
	for {
		fmt.Println("======== Custom Input Test ========")
		fmt.Println("===================================")
		fmt.Println("1. Donuts")
		fmt.Println("2. A000124 of Sloane’s OEIS Formula")
		fmt.Println("3. Print Tree Path")
		fmt.Println("4. Finding Parent Tree")
		fmt.Println("5. Finding Children Node")
		fmt.Println("0. Exit")
		fmt.Println("===================================")

		fmt.Print("Choose Question Number to Test: ")

		var menu int
		fmt.Scan(&menu)

		switch menu {
		case 1:
			var value int
			fmt.Println()
			fmt.Print("Enter Value: ")
			fmt.Scan(&value)  
			fmt.Println("Input:", value, " Output:", GetDonutsSlice(value))
			fmt.Println()
			fmt.Println()
    case 2:
			var value int
			fmt.Println()
			fmt.Print("Enter Value: ")
			fmt.Scan(&value)  
			fmt.Println("Input:", value, " Output:", GetRowsValue(value))
			fmt.Println()
			fmt.Println()
		case 3:
			fmt.Println()
			PrintNodeTree("", &tree.root)
			fmt.Println()
			fmt.Println()
		case 4:
			var value string
			fmt.Print("Enter Input: ")
			fmt.Scan(&value)
			inputArr := strings.Split(value, "-")
			fmt.Println(len(inputArr))
			fmt.Println("Input: ", value, " Output:", IsValidParent(inputArr, len(inputArr), 0, &tree.root))
			fmt.Println()
			fmt.Println()
		case 5:
			var value string
			fmt.Print("Enter Input: ")
			fmt.Scan(&value)
			inputArr := strings.Split(value, "-")
			fmt.Println(len(inputArr))
			fmt.Println("Input: ", value, " Output:", GetNodeChildren(inputArr, len(inputArr), 0, &tree.root))
			fmt.Println()
			fmt.Println()
		}

		if menu == 0 {
			break
		}

	}
}

/*
	Soal Nomor 1:
*/
func GetDonutsSlice(sliceCount int ) int {
	if sliceCount == 0 {
		return 0
	}

	return sliceCount*(sliceCount + 1)/2 + 1;
}

/*
	Soal Nomor 2:
*/
func GetRowsValue(value int) string {
	if value == 0 {
		return "No Output"
	}

	var row string;
	for i := 0; i < value; i++ {
		number := i*(i + 1)/2 + 1;
    row += strconv.Itoa(number) + "-";
	}

	return row[0:len(row)-1];
}

/*
	Soal Nomor 3:
*/
func PrintNodeTree(strip string, node *Node){
	if node == nil {
		return
	}

	fmt.Println(strip + string(node.data))

	for _, element := range node.children{
		PrintNodeTree((strip + "-"), &element)
	} 

}

/*
	Soal Nomor 4:
*/
func IsValidParent(nodes []string, nodesLength int, index int, node *Node) bool {
	if node == nil {
		return false;
	}

	if string(node.data) == nodes[index]{
		if index == nodesLength-1 && len(node.children) == 0 {
			return true
		} else if index != nodesLength-1 {
			for i := range node.children {
				if string(node.children[i].data) == nodes[index+1] {
					return IsValidParent(nodes, nodesLength, index+1, &node.children[i])
				}
			}
		}
	}

	return false;
}

/*
	Soal Nomor 5:
*/
func GetNodeChildren(nodes []string, nodesLength int, index int, node *Node) string {
	if node == nil {
		return "Tidak ada Child"
	}

	if string(node.data) == nodes[index]{
		if index == nodesLength-1 && len(node.children) > 0 {
			var value string
			
			for _,element := range node.children{
				value += string(element.data) + "-"
			}

			return value[0:len(value)-1];

		} else {
			for i := range node.children {
				if string(node.children[i].data) == nodes[index+1] {
					return GetNodeChildren(nodes, nodesLength, index+1, &node.children[i])
				}
			}
		}
	}

	return "Tidak ada Child"
}
