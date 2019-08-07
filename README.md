# AVL Tree in Go

Just an exercize to see how hard it is to get AVL
trees working.

# Build it

    % go build avltree.go
    % go build avlfuzz.go

# Run it

    % ./avlfuzz 15
    [7 3 8 14 9 2 13 1 11 12 4 10 5 0 6]
    0
    1
    2
    ...


    % ./avltree 7 3 8 14 9 2 13 1 11 12 > avl.dot
    % dot -Tpng -o avl.png avl.dot
    ... view avl.png
