package huffman

import (
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"
)

func Test_searchAndSortWeight(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  []rune
		want1 []uint
	}{
		// TODO: Add test cases.
		{
			"test case 1: a2 b2 c2",
			args{s: "aabbcc"},
			[]rune{'a', 'b', 'c'},
			[]uint{2, 2, 2},
		},
		{
			"test case 2: a1 b2 c3",
			args{s: "abbccc"},
			[]rune{'a', 'b', 'c'},
			[]uint{1, 2, 3},
		},
		{
			"test case 3: a3 b2 c1",
			args{s: "aaabbc"},
			[]rune{'c', 'b', 'a'},
			[]uint{1, 2, 3},
		},
		{
			"test case 4: a1 b3 c2",
			args{s: "abbbcc"},
			[]rune{'a', 'c', 'b'},
			[]uint{1, 2, 3},
		},
		{
			"test case 5: a1 b2 c2",
			args{s: "abbcc"},
			[]rune{'a', 'b', 'c'},
			[]uint{1, 2, 2},
		},
		{
			"test case 6: a2 b2 c1",
			args{s: "aabbc"},
			[]rune{'c', 'a', 'b'},
			[]uint{1, 2, 2},
		},
		{
			"test case 7: ABAABACD",
			args{s: "ABAABACD"},
			[]rune{'D', 'C', 'B', 'A'},
			[]uint{1, 1, 2, 4},
		},
		{
			"test case 8: ",
			args{s: func() string {
				rand.Seed(time.Now().UnixNano())
				s := "BBBGGGGGGGGGGGGDDDDDDDCCCCAAEEEEEEEEFFFFFFFFFFF"
				b := strings.Builder{}
				for len(s) != 0 {
					i := rand.Intn(len(s))
					b.WriteByte(s[i])
					if i == len(s)-1 {
						s = s[:i]
					} else {
						s = s[:i] + s[i+1:]
					}
				}
				return b.String()
			}()},
			[]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G'},
			[]uint{2, 3, 4, 7, 8, 11, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := searchAndSortWeight(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchAndSortWeight() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("searchAndSortWeight() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_makeHuffmanTree(t *testing.T) {
	type args struct {
		rs []rune
		ws []uint
	}
	tests := []struct {
		name string
		args args
		want *tNode
	}{
		// TODO: Add test cases.
		{
			"test case 1: abbccc",
			args{
				rs: func() []rune {
					rs, _ := searchAndSortWeight("abbccc")
					return rs
				}(),
				ws: func() []uint {
					_, ws := searchAndSortWeight("abbccc")
					return ws
				}(),
			},
			&tNode{
				w: 6,
				l: &tNode{
					w: 3,
					l: &tNode{v: 'a', w: 1},
					r: &tNode{v: 'b', w: 2},
				},
				r: &tNode{v: 'c', w: 3},
			},
		},
		{
			"test case 2: ABAABACD",
			args{
				rs: func() []rune {
					rs, _ := searchAndSortWeight("ABAABACD")
					return rs
				}(),
				ws: func() []uint {
					_, ws := searchAndSortWeight("ABAABACD")
					return ws
				}(),
			},
			&tNode{
				w: 8,
				l: &tNode{
					w: 4,
					l: &tNode{
						w: 2,
						l: &tNode{v: 'D', w: 1},
						r: &tNode{v: 'C', w: 1},
					},
					r: &tNode{v: 'B', w: 2},
				},
				r: &tNode{v: 'A', w: 4},
			},
		},
		{
			"test case 3: BBBGGGGGGGGGGGGDDDDDDDCCCCAAEEEEEEEEFFFFFFFFFFF",
			args{
				rs: func() []rune {
					rs, _ := searchAndSortWeight("BBBGGGGGGGGGGGGDDDDDDDCCCCAAEEEEEEEEFFFFFFFFFFF")
					return rs
				}(),
				ws: func() []uint {
					_, ws := searchAndSortWeight("BBBGGGGGGGGGGGGDDDDDDDCCCCAAEEEEEEEEFFFFFFFFFFF")
					return ws
				}(),
			},
			&tNode{
				w: 48,
				l: &tNode{},
				r: &tNode{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeHuffmanTree(tt.args.rs, tt.args.ws); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeHuffmanTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeHuffmanCode(t *testing.T) {
	type args struct {
		n *tNode
	}
	tests := []struct {
		name string
		args args
		want map[rune]uint
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				n: makeHuffmanTree(
					func() []rune {
						rs, _ := searchAndSortWeight("ABAABACD")
						return rs
					}(),
					func() []uint {
						_, ws := searchAndSortWeight("ABAABACD")
						return ws
					}(),
				)},
			map[rune]uint{
				'A': 0b0,
				'B': 0b10,
				'C': 0b110,
				'D': 0b111,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeHuffmanCode(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeHuffmanCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
