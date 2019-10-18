# graphtheory
*An opensource Go library for manipulating and analysing graphs*  

<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-6%25-brightgreen.svg?longCache=true&style=flat)</a>  

This library is a work in progress, but provides a framework for playing around with different ideas in graph theory.  

My main goal is experimenting with different ways of implementing graph theory algorithms concurrently (i.e., concurrent algorithms for testing graph isomorphism or counting valid colorings).  

## Sample program

```go
package main

import (
  "fmt"
  "graphtheory/graph"
)

func main() {
  //Generate all graphs with 6 vertices
  graphs := graph.AllNGraphs(6)

  // Print out the graphs (displayed as adjecency matrices)
  for _, g := range graphs {
    fmt.Println(g)
  }
}
```  
