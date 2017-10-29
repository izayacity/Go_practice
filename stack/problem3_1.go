package stack

/*
3.1 Three in One: Describe how you could use a single array to implement three stacks
*/
type SliceMultiStack struct {
	slice    []*stackNode // slice holding nodes, so we can set nils
	sizes    []int        // number of elements in each stack
	stacks   int          // number of stacks
	stackMax int          // size of each stack
}
