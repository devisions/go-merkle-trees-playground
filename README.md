# go-merkle-trees-playground

A playground for Go-based Merkle Tree implementations

## Contents

| implementation | github<br/>stars/src | source / notes |
| :--- | --- | :--- |
| `cbergoon___merkletree` | [219](https://github.com/cbergoon/merkletree) | It has a simple API and implementation.<br/>It stores the data in the leafs, not outside the tree. |
| `codenotary___merkletree` | [21](https://github.com/codenotary/merkletree) | A Merkle Tree implementation that follows RFC 6962. |
| `laser___go-merkle-tree` | [7](https://github.com/laser/go-merkle-tree)<br/> | It uses prefixes for `Leaf`s and `Node`s<br/>(it also uses `Branch`es for linking nodes).<br/>It keeps a balanced binary tree<br/>(duplicated nodes, for any single child).<br/>It includes _proof_ for auditing and verification purposes. |

