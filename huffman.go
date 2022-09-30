package huffman

import "fmt"

// search every rune
func searchAndSortWeight(s string) ([]rune, []uint) {
	rs := make([]rune, 0, len(s))
	ws := make([]uint, 0, len(s))
	for _, r := range s {
		for i, lr := range rs {
			// found
			if lr == r {
				nw := ws[i] + 1
				wl := len(ws)
				// not max weight, j = ?
				// nw = 1 + 1 = 2
				// 1 [1] 2 2 3 3
				//    +
				// 1 [2] 2 2 3 3
				//    ↑      j
				// 1  _  2 2 3 3
				//      ← ←
				// 1 2 2  _  3 3
				//        ↓
				// 1 2 2 [2] 3 3
				// max weight, j = wl
				// nw = 2 + 1 = 2
				// 1 1 [2] 2 3 3
				//      +
				// 1 1 [3] 2 3 3
				//      ↑  k
				// 1 1  _  2 3 3
				//        ← ← ←
				// 1 1 2 3 3  _
				//            ↓
				// 1 1 2 3 3 [3]
				var j int
				for j = i + 1; j != wl; j++ {
					if nw < ws[j] {
						break
					}
				}
				// swap
				for k := i + 1; k != j; k++ {
					ws[k-1] = ws[k]
					rs[k-1] = rs[k]
				}
				// insert
				ws[j-1] = nw
				rs[j-1] = r
				goto NEXT
			}
		}
		// not found
		rs = append(append(make([]rune, 0, len(s)), r), rs...)
		ws = append(append(make([]uint, 0, len(s)), 1), ws...)
	NEXT:
	}
	return rs, ws
}

type tNode struct {
	v rune
	w uint
	l *tNode
	r *tNode
}

func makeHuffmanTree(rs []rune, ws []uint) *tNode {
	nl := len(rs)
	if nl < 2 {
		return nil
	}
	rootSlice := []*tNode{
		&tNode{
			w: ws[0] + ws[1],
			l: &tNode{v: rs[0], w: ws[0]},
			r: &tNode{v: rs[1], w: ws[1]},
		},
	}
	for i := 2; i != nl; i++ {
		var ln, rn *tNode
		for _, root := range rootSlice {

		}

		// if ws[i] <  {

		// }

		// if root.w <= ws[i] {
		// 	ln = root
		// 	rn = &tNode{v: rs[i], w: ws[i]}
		// } else {
		// 	ln = &tNode{v: rs[i], w: ws[i]}
		// 	rn = root
		// }
		// root = &tNode{
		// 	w: ln.w + rn.w,
		// 	l: ln,
		// 	r: rn,
		// }
	}
	return root
}

func makeHuffmanCode(n *tNode) map[rune]uint {
	var d uint = 2
	m := make(map[rune]uint)
	for n != nil {
		if n.r != nil {
			if n.r.l != nil || n.r.r != nil {
				panic(fmt.Sprintf("n.r.l %v is not nil or n.r.r %+v is not nil", n.r.l, n.r.r))
			}
			if n.r.v == 0 {
				panic("n.r.v is 0")
			}
			m[n.r.v] = (1 << (d - 1)) - 2
		}
		if n.l != nil {
			if n.l.v != 0 {
				m[n.l.v] = (1 << (d - 1)) - 1
			}
		}
		d++
		n = n.l
	}
	return m
}
