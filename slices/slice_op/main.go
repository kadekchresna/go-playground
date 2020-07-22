package main

import "fmt"

func main() {
	fmt.Println("eyy")
	bs := []int{1, 2, 3, 4}
	ts := make([]int, 1, 1)

	// Copy
	copy(ts, bs)
	fmt.Println(bs, cap(bs), len(bs))
	fmt.Println(ts, cap(ts), len(ts))
	fmt.Println()

	// Append
	as := append([]int{}, []int{2, 1, 3, 34}...)
	fmt.Println(as, cap(as), len(as))
	fmt.Println()

	// Cut
	cs := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(cs, cap(cs), len(cs))
	cs = append(cs[:2], cs[3:]...)
	fmt.Println(cs, cap(cs), len(cs))
	fmt.Println()

	// Delete
	fmt.Println(bs)
	dso := append(bs[:1], bs[2:]...)
	fmt.Println(dso, cap(dso), len(dso))
	fmt.Println()

	// Delete
	// fmt.Println(copy(bs[1:], bs[2:]))
	tmp := []int{1, 23, 45, 2, 1, 3, 2}
	fmt.Println(tmp, cap(tmp), len(tmp))
	fmt.Println(tmp[1:])
	fmt.Println(tmp[2:])
	fmt.Println(copy(tmp[1:], tmp[2:]))
	fmt.Println(tmp[1:])
	fmt.Println(tmp, cap(tmp), len(tmp))

	dsp := append(tmp[:6])
	// dsp := append(tmp[:1+copy(tmp[1:], tmp[2:])])
	fmt.Println(dsp, cap(dsp), len(dsp))
	fmt.Println()

	// Extend
	es := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(es, cap(es), len(es))
	aes := append(es, make([]int, 2)...)
	fmt.Println(aes, cap(aes), len(aes))
	fmt.Println()

	fmt.Println(es, cap(es), len(es))
	eaes := append(es[:3], append(make([]int, 3), es[4:]...)...)
	fmt.Println(eaes, cap(eaes), len(eaes))
	fmt.Println()

}
