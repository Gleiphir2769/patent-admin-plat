package charts

import (
	"fmt"
	"testing"
)

func TestBar(t *testing.T) {
	cate := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	data := []int{10, 52, 200, 334, 390, 330, 220}
	s := genBarProfile(cate, data, true)
	fmt.Println(s)
}

func TestPie(t *testing.T) {
	classes := []string{"发明", "新型", "外观"}
	data := []int{10, 52, 100}
	s := genPieProfile(classes, data)
	fmt.Println(s)
}

func TestLineSmooth(t *testing.T) {
	cate := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	data := []int{10, 52, 200, 334, 390, 330, 220}
	s := genLineSmoothProfile(cate, data, true)
	fmt.Println(s)
}

func TestVerticalBar(t *testing.T) {
	cate := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	data := []int{10, 52, 200, 334, 390, 330, 220}
	s := genVerticalBarProfile(cate, data, true)
	fmt.Println(s)
}
