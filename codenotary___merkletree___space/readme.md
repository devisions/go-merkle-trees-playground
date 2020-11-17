

Adding three data elements ("d0", "d1", "d2") to the storer:
```go
merkletree.Append(s, []byte("d0"))
merkletree.Append(s, []byte("d1"))
merkletree.Append(s, []byte("d2"))
```
and printing it we get the following tree:
```
            c64c                            
    46c7            
c67f    49b7    f366
```
