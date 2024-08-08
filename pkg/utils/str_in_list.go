package utils

// StrInList check if given str is in the list
func StrInList(list []string, item string) bool {
  for _, k := range list {
    if k == item {
      return true
    }
  }

  return false
}
