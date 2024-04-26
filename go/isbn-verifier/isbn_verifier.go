package isbn

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//(d₁ * 10 + d₂ * 9 + d₃ * 8 + d₄ * 7 + d₅ * 6 + d₆ * 5 + d₇ * 4 + d₈ * 3 + d₉ * 2 + d₁₀ * 1) mod 11 == 0
func IsValidISBN(isbn string) bool {
  if isbn == ""{
    return false
  }

  if strings.Count(isbn, "-") == 3 && len(isbn) != 13{
    fmt.Println("dashes/ len not 13")
    return false
  }

  if strings.Count(isbn, "-") == 0 && len(isbn) != 10{
    fmt.Println("no dashes/ len not 10")
    return false
  }
 
  isbn = strings.ReplaceAll(isbn, "-", "")
  result := 0
  count := 10

  for i, v := range isbn{
    if !unicode.IsDigit(v){
      if i != 9{
        return false
      }
      if v == 'X'{
        v = 10
      } else {
        return false
      }
    }
    v, _ := strconv.Atoi(string(v)) 
    result += (v * (count - i))
  }

  return (result % 11 == 0)
}
