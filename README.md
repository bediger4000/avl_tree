# AVL Tree in Go

An exercize to see how hard it is to get AVL trees working.
I got insert working in about 2 days of fooling around in my free time.
I did some development on a Mac, some on Linux.
Go makes portability easy.

I studied [this class handout](https://courses.cs.washington.edu/courses/cse373/06sp/handouts/lecture12.pdf)
to understand it all.

# Build it

    % cd $GOPATH/src
    % git clone git@github.com:bediger4000/avl_tree.git
    % cd avl_tree
    % go build avltree.go
    % go build avlfuzz.go

# Run it

    % ./avlfuzz 15
    [7 3 8 14 9 2 13 1 11 12 4 10 5 0 6]
    0
    1
    2
    ...

`avlfuzz` creates a slice (default length 10, but you can set it on command line)
of integers from 0 to N, shuffles them, then builds a binary search tree using
the AVL algorithm to insert the integers.


    % ./avltree 7 3 8 14 9 2 13 1 11 12 > avl.dot
    % dot -Tpng -o avl.png avl.dot
    ... view avl.png

`avltree` uses a list of integers on the command line to create a [GraphViz]()
representation of a binary search tree.
It inserts the integers from the list into a binary tree using the AVL algorithm.
You can use `dot` (part of GraphViz) to render `avltree` output into PNG, or PostScript
or whatever.

## Example

After building it, you can get an image to see like this:

    % ./avltree 5 7 9 2 6 8 4 3 0 1 > avl.dot
    % dot -Tpng -o avl.png avl.dot

![AVL binary search tree](avl.png?raw=true)

Empty/null child nodes are represented as small black dots,
to keep GraphViz algorithms from rearranging everything.
I believe I got that idea from [here](https://eli.thegreenplace.net/2009/11/23/visualizing-binary-trees-with-graphviz).

Each oval node has a label like "3,2/0".
That means the node has a data value of 3, a depth of 2,
and a "balance factor" of 0.
Since leaf nodes have a depth of 0, and two layers of
nodes appear under "3,2/0", the depth is correct.
The balance factor is right depth - left depth.
Node 3 has 2 child nodes, 1 and 5, both of depth 1,
so a balance factor of 0 for node 3 is correct.

