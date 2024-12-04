package algorithms

import "testing"

func TestQuickSort(t *testing.T) {
  unsorted := []int{7,1,5,4}
  n := len(unsorted)

  sorted := []int{1,4,5,7}
  QuickSort(unsorted, 0, n - 1)
  got := unsorted 
  want := sorted
  
  for i := range want {
    if got[i] != want[i] {
      t.Errorf("Slices are not sorted. got: %d, want %d", got, want)
    }
  }
}
