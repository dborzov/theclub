package main

// serialized trees in quasy lisp syntax
const (
	ex1 = `(3,(1,1,1),(8,5,13))`
	ex2 = `(8,(3,(1,1,2),5),13)`
)


func loadBranch(s []byte) *Tree, err {
  if len(s) == 0 {
    return nil, errors.New("empty string")
  }
  if s[0] != byte(`(`) {
    val, err := strconv.Atoi(string(s))
    if err != nil {
      return nil, err
    }
    return *Tree{Value: val}, nil
  }
  strings


}
