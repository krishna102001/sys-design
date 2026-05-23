## MERKLE-TREE
merkle tree is also know as hash tree. it is a cryptographic tree data structure used to efficently and securly verify the
large dataset whether its got corrupted or tampered by someone.
merkle tree are built from bottom to up approach where each leaf node is a data and if there is 
odd number or node data then we will simply make the copy of the last node so its will become even
because the construction of tree will from the bottom to top where two leaf node paired up are hashed and make the new node. applying like this we will reach to one node which will became the root node of the merkle tree.
