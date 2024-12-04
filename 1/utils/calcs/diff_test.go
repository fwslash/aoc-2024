package calcs

import "testing"

func TestArrDiff(t *testing.T) {
  arr1 := []int{1,2,3}
  arr2 := []int{2,3,4}

  got := ArrDiff(arr1, arr2)
  want := 3

  if got != want {
    t.Errorf("Arr Diff did not return correct number. Got: %d, Want; %d ", got, want)
  }
}
