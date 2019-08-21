package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	//"path/filepath"
	//"time"
)

type round struct {
	Alice string `json:"alice"`
	Bob string `json:"bob"`
	Result int `json:"result"`
}

func readFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}
	return data
}

func main() {
	var raw map[string][]round
	var answer map[string][]round

	basePath := os.Getenv("HOME")
	inputData := readFile(filepath.Join(basePath, "match.json"))
	err := json.Unmarshal(inputData, &raw)
	if err != nil {
		panic(err)
	}
	questions := raw["matches"]

	//b, c := raw["matches"][0].Alice[1], raw["matches"][0].Alice[0]
	//fmt.Println(raw["matches"][0], len(raw["matches"][0].Alice), byte(b), byte(c))

	outputData := readFile(filepath.Join(basePath, "result.json"))
	err = json.Unmarshal(outputData, &answer)
	if err != nil {
		panic(err)
	}
	//correct := answer["matches"]
	number := len(questions)
	//fmt.Println(number)
	alices:=[]string{}
	aliceColor := [][]int{}
	bobs:=[]string{}
	bobColor := [][]int{}
	for i:=0; i<number; i++ {
		a, ac := preprocess(questions[i].Alice)
		alices = append(alices, a)
		aliceColor = append(aliceColor, ac)

		b, bc := preprocess(questions[i].Bob)
		bobs = append(bobs, b)
		bobColor = append(bobColor, bc)
	}
	//mgr1 := newCardBuf()
	//mgr2 := newCardBuf()
	mgr := newSimpleCards()

	start := time.Now()
	for i:=0; i<number; i++ {
		//alice := process(questions[i].Alice)
		//bob := process(questions[i].Bob)
		//questions[i].Result = compareResult(alice, bob)
		//aliceMode, aliceV := mgr.process(questions[i].Alice)
		//bobMode, bobV := mgr.process(questions[i].Bob)
		//TestAdd(mgr1, mgr2, questions[i].Alice, questions[i].Bob)
		//aliceMode, aliceV, bobMode, bobV := Process(mgr1, mgr2, questions[i].Alice, questions[i].Bob)
		aliceMode, aliceV := mgr.process(alices[i], aliceColor[i])
		bobMode, bobV := mgr.process(bobs[i], bobColor[i])
		questions[i].Result = compareResult(aliceMode, aliceV, bobMode, bobV)
	}
	fmt.Println("5 cards without ghost cost: ", time.Now().Sub(start))

	//for i:=0; i<len(questions); i++ {
	//	if questions[i].Result != correct[i].Result {
	//		fmt.Println(questions[i], " need ", correct[i])
	//	}
	//}

	////////////////////////// 7cards//////////////////////////////
	inputData = readFile(filepath.Join(basePath, "7cards.json"))
	err = json.Unmarshal(inputData, &raw)
	if err != nil {
		panic(err)
	}
	questions = raw["matches"]

	//b, c := raw["matches"][0].Alice[1], raw["matches"][0].Alice[0]
	//fmt.Println(raw["matches"][0], len(raw["matches"][0].Alice), byte(b), byte(c))

	outputData = readFile(filepath.Join(basePath, "7cars_ret.json"))
	err = json.Unmarshal(outputData, &answer)
	if err != nil {
		panic(err)
	}
	//correct = answer["matches"]
	mgr1 := newCardBuf()
	mgr2 := newCardBuf()

	number = len(questions)
	start = time.Now()
	for i:=0; i<number; i++ {
		//alice := process(questions[i].Alice)
		//bob := process(questions[i].Bob)
		//questions[i].Result = compare(alice, bob)
		//aliceMode, aliceV := mgr.process(questions[i].Alice)
		//bobMode, bobV := mgr.process(questions[i].Bob)
		aliceMode, aliceV, bobMode, bobV := Process(mgr1, mgr2, questions[i].Alice, questions[i].Bob)
		questions[i].Result = compareResult(aliceMode, aliceV, bobMode, bobV)
	}
	fmt.Println("7 cards without ghost cost: ", time.Now().Sub(start))

	//for i:=0; i<len(questions); i++ {
	//	if questions[i].Result != correct[i].Result {
	//		fmt.Println(questions[i], " need ", correct[i])
	//	}
	//}
	////////////////////////5 with ghost//////////////////////
	inputData = readFile(filepath.Join(basePath, "five_cards_with_ghost.json"))
	err = json.Unmarshal(inputData, &raw)
	if err != nil {
		panic(err)
	}
	questions = raw["matches"]
	mgr1 = newCardBuf()
	mgr2 = newCardBuf()

	number = len(questions)
	start = time.Now()
	for i:=0; i<number; i++ {
		aliceMode, aliceV, bobMode, bobV := Process(mgr1, mgr2, questions[i].Alice, questions[i].Bob)
		questions[i].Result = compareResult(aliceMode, aliceV, bobMode, bobV)
	}
	fmt.Println("5 cards with ghost cost: ", time.Now().Sub(start))
	//////////////////////////7 with ghost///////////////////////////////////
	inputData = readFile(filepath.Join(basePath, "7cards_with_ghost.json"))
	err = json.Unmarshal(inputData, &raw)
	if err != nil {
		panic(err)
	}
	questions = raw["matches"]
	mgr1 = newCardBuf()
	mgr2 = newCardBuf()

	number = len(questions)
	start = time.Now()
	for i:=0; i<number; i++ {
		aliceMode, aliceV, bobMode, bobV := Process(mgr1, mgr2, questions[i].Alice, questions[i].Bob)
		questions[i].Result = compareResult(aliceMode, aliceV, bobMode, bobV)
	}
	fmt.Println("7 cards with ghost cost: ", time.Now().Sub(start))
}

//func main() {
//	//fmt.Println(process("AsKsQsJsJsKsQs"))
//	//fmt.Println(process("KsKsQsJsQsKsQs"))
//	//fmt.Println(process("KsKsQsJsTsKs8s"))
//	//fmt.Println(process("KsKsKsJsTsKs8s"))
//	//arr := []int{3,44,38,5,47,15,36,26,15,27,2,46,28,4,19,50,48}
//	//fmt.Println(quickSort(arr))
//	//a := process("5c2s6dThQsJh9d")
//	//b := process("Ac7s5c2s6dThQs")
//	//fmt.Println(a, b, compare(a, b))
//	//mgr := newCardBuf()
//	//m, v := mgr.process("Ac7dJs6c5c")
//	//fmt.Println(tabExplain[m], v)
//	//
//	//m, v = mgr.process("7c8d5dJd9d")
//	//fmt.Println(tabExplain[m], v)
//	//
//	//m, v = mgr.process("AsKsQsJsJsKsQs")
//	//fmt.Println(tabExplain[m], v)
//	//
//	//m, v = mgr.process("Xn8dKsKd9h")
//	//fmt.Println(tabExplain[m], v)
//	////
//	//m, v = mgr.process("6s5h4c3s2c")
//	//fmt.Println(tabExplain[m], v)
//	//
//	//m, v = mgr.process("As2h3s4c5s")
//	//fmt.Println(tabExplain[m], v)
//	//ret := make(map[string]seqCards)
//	//create5cardsTable(ret, divideStr("A5432"), "", 5, flush)
//	//ret := createFullTableWithoutGhost()
//	//fmt.Println(ret)
//	//ret = createFullTableWithGhost()
//	//fmt.Println(ret)
//	//create32Table()
//	//ret := []string{}
//	//permutate([]string{"J","J","T","K","Q"}, 0, 5, &ret)
//	//fmt.Println(ret)
//	ret := createFullTableWithoutGhost()
//	create32Table(ret)
//	create3Table(ret)
//	create4Table(ret)
//	createCouple2Table(ret)
//	createCoupleTable(ret)
//	fmt.Println(ret["627T3"])
//	//fmt.Println(len(ret))
//	ret1 := make(map[string]seqCards)
//	create3Table(ret1)
//
//	cb := newSimpleCards()
//	cards := "3A2AA"
//	start := time.Now()
//	for i:=0; i<20000; i++ {
//		//if _, ok := ret["3hAh2cAsAc"]; ok {
//		////if _, ok := ret1["3hAh2cAsAc"]; ok {
//		//}
//		cb.add5Card(cards, []int{2,1,0,0})
//	}
//	fmt.Println(time.Now().Sub(start))
//
////	mgr := newSimpleCards()
////	m, v := mgr.process("3hAh2cAsAc")
////	m1, v1 := mgr.process("6h6s6cJhAd")
////	fmt.Println(tabExplain[m], v, tabExplain[m1], v1)
////	fmt.Println(compareResult(m, v, m1, v1))
//}

func preprocess(raw string) (string, []int) {
	str := ""
	color := []int{0,0,0,0}
	for i:=0;i<len(raw);i+=2 {
		str += string(raw[i])
		se := colorTable[string(raw[i+1])]
		color[se]++
	}
	return str, color
}

var table map[string]int = map[string]int {
	"2":2, "3":3, "4":4,"5":5,"6":6,"7":7,"8":8,"9":9,"T":10,"J":11, "Q":12,"K":13,"A":14,
}

var invTable map[int]string = map[int]string {
	2:"2", 3:"3", 4:"4", 5:"5",6:"6", 7:"7", 8:"8", 9:"9", 10:"T", 11:"J", 12:"Q", 13:"K", 14:"A",
}

const (
	same = true
	diff = false

	over = 1
	less = 2
	equal = 0

)

const (
	royal = iota //皇家同花顺
	flush        //同花顺
	four         //4条
	threeTwo     //3+2
	suit         //同花
	sequence     //顺子
	three        //3张
	couple2      //两对
	couple       //一对
	alone        //散
)

var tabExplain []string = []string{"royal", "flush", "four", "3+2", "suit", "sequence", "3", "2+2", "2", "sand"}
//todo: 状态模式
func checkKind(good []int, pos int, oldState int) int {
	switch oldState {
	case alone:
		return couple
	case couple:
		if pos>=2 {
			if good[pos-2] == good[pos-1] {
				return three
			} else if good[pos-2] != good[pos-1] {
				return couple2
			}
		} else {
			return couple2
		}
	case couple2:
		if pos>=2 {
			if good[pos-2] == good[pos-1] {
				return threeTwo
			} else {
				return couple2
			}
		} else {
			return couple2
		}
	case three:
		if pos>=3 {
			if good[pos-2] == good[pos-1] && good[pos-3] == good[pos-1] {
				return four
			} else if good[pos-2] == good[pos-1] && good[pos-3] != good[pos-1] {
				return threeTwo
			} else if good[pos-2] != good[pos-1] {
				return threeTwo
			}
		} else {
			return threeTwo
		}
	case threeTwo:
		if pos>=3 {
			if good[pos-2] == good[pos-1] && good[pos-3] == good[pos-1] {
				return four
			} else {
				return threeTwo
			}
		} else {
			return threeTwo
		}
	case four:
		return four
	}

	return oldState
}

func sort(good []int, length int, oldState int) int {
	//bubble sort
	for i:=length; i>0; i-- {
		if good[i] > good[i-1] {
			//swap
			good[i], good[i-1] = good[i-1], good[i]
		} else if(good[i] == good[i-1]) {
			// check same cards
			return checkKind(good, i, oldState)
		} else {
			return oldState
		}
	}
	return oldState
}

type cardSort struct {
	good []int
	buck map[int][]int
	kind int
	isSame bool
}

func process(cards string) *cardSort {
	//ghost := false
	length := len(cards)
	good := make([]int, length>>1)
	for i:=1; i<(length>>1); i++ {
		good[i]=0
	}

	var buck map[int][]int
	good[0] = table[string(cards[0])]
	lastColor := string(cards[1])
	isSame := same
	kind := alone

	for i:=2; i<length; i+=2 {
		//if string(cards[i]) == "X" {
		//	//ghost card
		//	ghost = true
		//	continue
		//}
		good[i>>1] = table[string(cards[i])]
		kind = sort(good, i>>1, kind)
		if lastColor != string(cards[i+1]) {
			// color same?
			isSame = diff
		}
	}

	// 顺子？
	if kind == alone {
		i := 0
		for i=0; i< (length>>1)-1; i++ {
			if good[i] != good[i+1]+1 {
				if good[i]==14 {
					// check A5432或A765432
					if good[i+1] == 5 || ((good[i+1] == 7)&&(length==14)) {
						continue
					}
				}
				break
			}
		}
		if i == (length>>1)-1 {
			if isSame == same {
				if good[0] == table["A"] && good[1] == table["K"] {
					kind = royal
				} else {
					kind = flush
				}

			} else {
				kind = sequence
			}
		} else {
			if isSame == same {
				kind = suit
			}
		}
	} else {
		// 存在成对的牌， 记录之
		buck = tuningSeq(good)
	}
	return &cardSort{good, buck, kind, isSame}
}

func tuningSeq(good []int) map[int][]int {
	buck := make(map[int][]int, 4) // 1,2,3,4张牌
	list := make([]int, 15)
	length := len(good)
	for i:=0; i<length; i++ {
		list[good[i]]++
	}

	for i:=2; i<15; i++ {
		if list[i] > 0 { // number exist
			if v, ok := buck[list[i]]; ok { // 多对3张或2张牌
				if list[i] == 3 {
					small := i
					if v[0] < i {
						buck[list[i]] = []int{i} //大的保存
						small = v[0]
					}
					// degrade
					buck[2] = []int{small}
					buck[1] = append(buck[1], small)

				} else if list[i] == 2 {
					if len(v) == 2 {
						//degrade
						small := i
						if v[0] < small {
							small = v[0]
							v[0] = i
						}
						if v[1] < small {
							small, v[1] = v[1], small
						}
						buck[2] = v
						buck[1] = append(buck[1], []int{small, small}...)
					} else {
						buck[2] = append(buck[2], i)
					}
				} else {
					buck[1] = append(buck[1], i)
				}
			} else {
				buck[list[i]] = []int{i}
			}
		}
	}
	if len(buck[3])>0 && len(buck[2])>1 {
		// 3+2 tuning
		big := buck[2][0]
		small := buck[2][1]
		if small > big {
			big, small = small, big
		}
		buck[2] = []int{big}
		buck[1] = append(buck[1], []int{small, small}...)
	}
	return buck
}

func quickSort(in []int) []int { //ascending
	length := len(in)
	if length < 2 {
		return in
	} else if length == 2 {
		if in[0] > in[1] {
			in[0], in[1] = in[1], in[0]
		}
		return in
	}
	pivot := in[0]
	index := 0
	i:=0
	j:=length-1


	for i<j {
		for i<j {
			if in[j] < pivot {
				in[index] = in[j]
				index = j
				j--
				break
			} else {
				j--
			}
		}
		for i<j {
			if in[i] > pivot {
				in[index] = in[j]
				index = i
				i++
				break
			} else {
				i++
			}
		}

	}

	in[index] = pivot

	dL := quickSort(in[:index])
	dR := quickSort(in[index+1:])

	dL = append(dL, pivot)
	dL = append(dL, dR...)
	return dL
}

func quickSortDes(in []int, length int) []int { //descending
	if length < 2 {
		return in
	} else if length == 2 {
		if in[0] < in[1] {
			in[0], in[1] = in[1], in[0]
		}
		return in
	}
	pivot := in[0]
	index := 0
	i:=0
	j:=length-1


	for i<j {
		for i<j {
			if in[j] > pivot {
				in[index] = in[j]
				index = j
				j--
				break
			} else {
				j--
			}
		}
		for i<j {
			if in[i] < pivot {
				in[index] = in[j]
				index = i
				i++
				break
			} else {
				i++
			}
		}

	}

	in[index] = pivot

	dL := quickSort(in[:index])
	dR := quickSort(in[index+1:])

	dL = append(dL, pivot)
	dL = append(dL, dR...)
	return dL
}

func compare(a, b *cardSort) int {
	ret := equal
	if a.kind < b.kind {
		return over
	} else if a.kind > b.kind {
		return less
	} else {
		if a.kind == flush || a.kind == sequence {
			if (a.good[0] > b.good[0] && a.good[0] != table["A"]) || (a.good[0] != table["A"] && b.good[0] == table["A"]) {
				return over
			} else if a.good[0] < b.good[0] {
				return less
			}
		} else if a.kind == suit || a.kind == alone {
			for i:=0; i<len(a.good); i++ {
				if a.good[i] > b.good[i] {
					return over
				} else if a.good[i] < b.good[i] {
					return less
				}
			}
		} else if a.kind == threeTwo {
			a3 := a.buck[3]
			b3 := b.buck[3]
			a2 := a.buck[2]
			b2 := b.buck[2]
			a1, ok1 := a.buck[1]
			b1, _ := b.buck[1]
			if a3[0] > b3[0] {
				return over
			} else if a3[0] < b3[0] {
				return less
			} else {
				if a2[0] > b2[0] {
					return over
				} else if a2[0] < b2[0] {
					return less
				} else {
					if ok1 {
						// 7张牌, 5张牌不走
						a1s := quickSort(a1)
						b1s := quickSort(b1)
						if a1s[1] > b1s[1] {
							return over
						} else if a1s[1] < b1s[1] {
							return less
						} else {
							if a1s[0] > b1s[0] {
								return over
							} else if a1s[1] < b1s[1] {
								return less
							}
						}
					}
				}
			}
		} else if a.kind == three {
			a3 := a.buck[3][0]
			b3 := b.buck[3][0]
			if a3 > b3 {
				return over
			} else if a3 < b3 {
				return less
			} else {
				a1, _ := a.buck[1]
				b1, _ := b.buck[1]
				a1s := quickSort(a1)
				b1s := quickSort(b1)
				for i := len(a1) - 1; i >= 0; i-- {
					if a1s[i] > b1s[i] {
						return over
					} else if a1s[i] < b1s[i] {
						return less
					}
				}
			}
		} else if a.kind == four {
			a4 := a.buck[4][0]
			b4 := b.buck[4][0]
			if a4>b4 {
				return over
			} else if a4<b4 {
				return less
			} else {
				a1 := a.buck[1]
				b1 := b.buck[1]
				if len(a1) == 1 { // 5张
					if a1[0] > b1[0] {
						return over
					} else if a1[0] < b1[0] {
						return less
					}
				} else if len(a1) == 3 { // 7张
					a1s := quickSort(a1)
					b1s := quickSort(b1)
					for i:=2; i>=0; i-- {
						if a1s[i] > b1s[i] {
							return over
						} else if a1s[i] < b1s[i] {
							return less
						}
					}
				}
			}
		} else if a.kind == couple2 {
			a2 := quickSort(a.buck[2])
			b2 := quickSort(b.buck[2])
			if a2[1] > b2[1] {
				return over
			} else if a2[1] < b2[1] {
				return less
			} else {
				if a2[0] > b2[0] {
					return over
				} else if a2[0] < b2[0] {
					return less
				} else {
					a1 := a.buck[1]
					b1 := b.buck[1]
					if len(a1) == 1 {
						//5张
						if a1[0] > b1[0] {
							return over
						} else if a1[0] < b1[0] {
							return less
						}
					} else if len(a1) == 3 {
						//7张
						a1s := quickSort(a1)
						b1s := quickSort(b1)
						for i:=2; i>=0; i-- {
							if a1s[i] > b1s[i] {
								return over
							} else if a1s[i] < b1s[i] {
								return less
							}
						}
					}
				}
			}

		} else if a.kind == couple {
			a2 := a.buck[2][0]
			b2 := b.buck[2][0]
			if a2 > b2 {
				return over
			} else if a2 < b2 {
				return less
			} else {
				a1s := quickSort(a.buck[1])
				b1s := quickSort(b.buck[1])
				for i:=len(a1s)-1; i>=0;i-- {
					if a1s[i] > b1s[i] {
						return over
					} else if a1s[i] < b1s[i] {
						return less
					}
				}
			}
		}
	}
	return ret
}

const (
	spades=iota
	hearts
	diamonds
	clubs
)

var colorTable map[string]int = map[string]int{"S":spades, "H":hearts, "D":diamonds, "C":clubs, "s":spades, "h":hearts, "d":diamonds, "c":clubs}

type cardColor struct {
	count int
	cards []int
}

type seqCards struct {
	mode int
	max []int
}

type cardBuf struct {
	s_order []int
	//single map[int]struct{}
	single int
	s_cnt int
	d_order []int
	//double map[int]struct{}
	double int
	d_cnt int
	t_order []int
	//tripple map[int]struct{}
	tripple int
	t_cnt int
	//card4 map[int]struct{}
	card4 int
	cardState int
	//color map[int]int
	//color []int
	color []cardColor
	same bool
	tabNoGhost map[string]int
	tabGhost map[string]int

	tabFullNoGhost map[string]seqCards
}

func newCardBuf() *cardBuf {
	cb := &cardBuf{cardState:alone, same:true}
	cb.s_order = make([]int, 7)
	//cb.single = make(map[int]struct{})
	cb.d_order = make([]int, 3)
	//cb.double = make(map[int]struct{})
	cb.t_order = make([]int, 2)
	//cb.tripple = make(map[int]struct{})
	//cb.card4 = make(map[int]struct{})
	//cb.color = make(map[int]int)
	//cb.color = make([]int, 4)
	cb.color = make([]cardColor, 4)
	cb.color[spades].cards = make([]int, 7)
	cb.color[clubs].cards = make([]int, 7)
	cb.color[hearts].cards = make([]int, 7)
	cb.color[diamonds].cards = make([]int, 7)

	cb.tabGhost = createSequenceTableWithGhost()
	cb.tabNoGhost = createSequenceTableWithoutGhost()
	cb.tabFullNoGhost = createFullTableWithoutGhost()

	return cb
}

func (cb *cardBuf) clear() {
	//cb.s_order = []int{}
	cb.s_cnt = 0
	cb.single = 0//make(map[int]struct{})
	//cb.d_order = []int{}
	cb.d_cnt = 0
	cb.double = 0//make(map[int]struct{})
	//cb.t_order = []int{}
	cb.t_cnt = 0
	cb.tripple = 0//make(map[int]struct{})
	cb.card4 = 0//make(map[int]struct{})
	//cb.color = make(map[int]int)
	//for i:=2; i<15; i++ {
	//	cb.color[i] = 0
	//}
	cb.color[diamonds].count, cb.color[clubs].count, cb.color[hearts].count, cb.color[spades].count = 0,0,0,0
}

func bubbleSort(good []int, length int) {
	//bubble sort
	for i:=length-1; i>0; i-- {
		if good[i] > good[i-1] {
			//swap
			good[i], good[i-1] = good[i-1], good[i]
		} else if(good[i] == good[i-1]) {
			// nothing
		} else {
			return
		}
	}
}

func sortCard(good []int, length int) {
	//fmt.Println(good, length)
	for j:=0; j<length-1; j++ {
		max := good[j]
		pivot := j
		for i := j+1; i < length; i++ {
			if good[i] > max {
				max = good[i]
				pivot = i
			}
		}
		good[pivot] = good[j]
		good[j] = max
		//fmt.Println(good)
	}
}

func stringContainedInSlice(str string, sli []string) bool {
	for _, s := range sli {
		if str == s {
			return true
		}
	}
	return false
}

func divideStr(s string) (ret []string) {
	for i:=0; i<len(s); i++ {
		ret = append(ret, string(s[i]))
	}
	return
}

func createCardsTable(num int, tab map[string]seqCards, all []string, part string, max []int, mode int) {
	if len(part) == num {
		sc := seqCards{mode, max}
		tab[part] = sc
		return
	}
	//fmt.Println(all)
	ss := divideStr(part)
	//fmt.Println(ss)
	//fmt.Println(part)
	for i:=0; i<len(all); i++ {
		temp := part
		if len(ss) == 0 || !stringContainedInSlice(all[i], ss) {
			temp += all[i]
			createCardsTable(num, tab, all, temp, max, mode)
		}
	}
}

func selectCards(num int, part string, ret *[]string) {
	if len(part) == num {
		//fmt.Println("+++++", part)
		*ret = append(*ret, part)
		return
	}
	all := []string{"A","K","Q","J","T","9","8","7","6","5","4","3","2"}
	ss := divideStr(part)
	for i:=0; i<len(all); i++ {
		temp := part
		if len(ss) == 0 || !stringContainedInSlice(all[i], ss) {
			temp += all[i]
			//fmt.Println(temp)
			selectCards(num, temp, ret)
		}
	}
}

func isSwap(str []string, begin, end int) bool {
	for i:=begin; i<end; i++ {
		if str[i] == str[end] {
			return false
		}
	}
	return true
}

func permutate(sample []string, begin, end int, kinds *[]string) {
	if begin==end {
		*kinds = append(*kinds, strings.Join(sample,""))
		return
	}

	for i:=begin;i<end;i++ {
		if isSwap(sample, begin, i) {
			sample[begin], sample[i] = sample[i], sample[begin]
			permutate(sample, begin+1, end, kinds)
			sample[begin], sample[i] = sample[i], sample[begin]
		}
	}
}

func create32Table(result map[string]seqCards) {
	raw := []string{}
	selectCards(2, "", &raw)
	//fmt.Println(raw)
	for _, item := range raw {
		material := []string{string(item[0]),string(item[0]),string(item[0]),string(item[1]),string(item[1])}
		kind := seqCards{threeTwo, []int{table[string(item[0])], table[string(item[1])]}}

		possible := []string{}
		permutate(material, 0, 5, &possible)
		for _, p := range possible {
			result[p]= kind
		}
	}
}

func create4Table(result map[string]seqCards) {
	raw := []string{}
	selectCards(2, "", &raw)
	//fmt.Println(raw)
	for _, item := range raw {
		material := []string{string(item[0]),string(item[0]),string(item[0]),string(item[0]),string(item[1])}
		kind := seqCards{four, []int{table[string(item[0])], table[string(item[1])]}}

		possible := []string{}
		permutate(material, 0, 5, &possible)
		for _, p := range possible {
			result[p]= kind
		}
	}
}

func create3Table(result map[string]seqCards) {
	raw := []string{}
	selectCards(3, "", &raw)
	//fmt.Println(raw)
	for _, item := range raw {
		material := []string{string(item[0]),string(item[0]),string(item[0]),string(item[1]),string(item[2])}
		max, min := table[string(item[1])], table[string(item[2])]
		if max < min {
			max, min = min, max
		}
		kind := seqCards{three, []int{table[string(item[0])], max, min}}

		possible := []string{}
		permutate(material, 0, 5, &possible)
		for _, p := range possible {
			result[p]= kind
		}
	}
}

func createCouple2Table(result map[string]seqCards) {
	raw := []string{}
	selectCards(3, "", &raw)
	//fmt.Println(raw)
	for _, item := range raw {
		material := []string{string(item[0]),string(item[0]),string(item[1]),string(item[1]),string(item[2])}
		max, min := table[string(item[0])], table[string(item[1])]
		if max < min {
			max, min = min, max
		}
		kind := seqCards{couple2, []int{max, min, table[string(item[2])]}}

		possible := []string{}
		permutate(material, 0, 5, &possible)
		for _, p := range possible {
			result[p]= kind
		}
	}
}

func createCoupleTable(result map[string]seqCards) {
	raw := []string{}
	selectCards(4, "", &raw)
	//fmt.Println(raw)
	for _, item := range raw {
		material := []string{string(item[0]),string(item[0]),string(item[1]),string(item[2]),string(item[3])}
		seq := []int{table[string(item[1])], table[string(item[2])], table[string(item[3])]}
		sortCard(seq, 3)
		sorted := []int{table[string(item[0])]}
		kind := seqCards{couple, append(sorted, seq...)}

		possible := []string{}
		permutate(material, 0, 5, &possible)
		for _, p := range possible {
			result[p]= kind
		}
	}
}

func createFullTableWithoutGhost() map[string]seqCards {
	ret := make(map[string]seqCards)
	smallTab := createSequenceTableWithoutGhost()

	for k, v := range smallTab {
		max := table[string(k[0])]
		if k=="A5432" {
			max = 5
		}
		createCardsTable(5, ret, divideStr(k), "", []int{max}, v)
	}
	return ret
}

func createFullTableWithGhost() map[string]seqCards {
	ret := make(map[string]seqCards)
	smallTab := createSequenceTableWithGhost()
	for k, v := range smallTab {
		alphabets := divideStr(k)
		max := table[string(k[0])]
		if stringContainedInSlice(k, []string{"A543", "A432", "A532", "A542"}) {
			max = 5
		} else {
			if table[alphabets[0]]==table[alphabets[1]]+1 && table[alphabets[1]]==table[alphabets[2]]+1 && table[alphabets[2]]==table[alphabets[3]]+1 {
				//顺序的数列，赖子做首数
				max = table[alphabets[0]]+1
			}
		}
		createCardsTable(4, ret, alphabets, "", []int{max}, v)
	}
	return ret
}

func createSequenceTableWithoutGhost() map[string]int {
	table := make(map[string]int)
	str := []string{"K","Q","J","T","9","8","7","6","5","4","3","2"}
	table["AKQJT"] = royal
	table["A5432"] = flush
	for i:=0; i<8; i++ {
		key:=str[i]+str[i+1]+str[i+2]+str[i+3]+str[i+4]
		table[key] = flush
	}
	//fmt.Println(table)
	return table
}

func createSequenceTableWithGhost() map[string]int {
	table := make(map[string]int)
	str := []string{"Q","J","T","9","8","7","6","5","4","3","2"}
	table["AKQJ"] = royal
	table["AQJT"] = royal
	table["AKQT"] = royal
	table["AKJT"] = royal
	table["KQJT"] = royal

	table["KJT9"] = flush
	table["KQT9"] = flush
	table["KQJ9"] = flush
	table["QJT9"] = flush

	table["A543"] = flush
	table["A432"] = flush
	table["A532"] = flush
	table["A542"] = flush

	for i:=0; i<6; i++ {
		for j:=0; j<5; j++ {
			key := ""
			for k:=0; k<5; k++ {
				if k!=j {
					key += str[i+k]
				}
			}
			table[key] = flush
		}
	}
	return table
}
func (cb *cardBuf) addCard(cards string, length int) bool {
	ghost := false
	for i:=0; i<length; i+=2 {
		if string(cards[i]) == "X" {
			ghost = true
			continue
		}
		card := table[string(cards[i])]
		se := colorTable[string(cards[i+1])]
		// 记录颜色
		//cb.color[card] |= 1<<uint(se)
		//cb.color[se] |= (1<<uint(card))
		cb.color[se].cards[cb.color[se].count] = card
		cb.color[se].count++
		//bubbleSort(cb.color[se].cards, cb.color[se].count)

		off := 1 << uint(card)
		//fmt.Println(card, se, cb.color[card])
		if (cb.single & off) == 0 { //if _, ok := cb.single[card]; !ok {
			//1张 not exist
			//cb.single[card]= struct{}{}
			cb.single |= off
			// sort
			//cb.s_order = append(cb.s_order, card)
			cb.s_order[cb.s_cnt] = card
			cb.s_cnt++
			//fmt.Println(cb.s_order, cb.s_cnt)
			//bubbleSort(cb.s_order, cb.s_cnt)
		} else if (cb.double & off) == 0 { //else if _, ok := cb.double[card]; !ok {
			//对 not exist
			//cb.double[card]= struct{}{}
			cb.double |= off
			// sort
			//cb.d_order = append(cb.d_order, card)
			cb.d_order[cb.d_cnt] = card
			cb.d_cnt++
			//bubbleSort(cb.d_order, cb.d_cnt)
		} else if (cb.tripple & off) == 0 { //else if _, ok := cb.tripple[card]; !ok {
			// 3张 not exist
			//cb.tripple[card] = struct{}{}
			cb.tripple |= off
			// sort
			//cb.t_order = append(cb.t_order, card)
			cb.t_order[cb.t_cnt] = card
			cb.t_cnt++
			//bubbleSort(cb.t_order, cb.t_cnt)
		} else {
			// 4张
			cb.card4 = card
		}
	}
	return ghost
}

func isRoyalorFlushorSeq (dat int, iSame bool) int {
	if dat == 14 && iSame {
		return royal
	} else if iSame {
		return flush
	}
	return sequence
}

// 搜索4张颜色相同
//func (cb *cardBuf) check4ColorSame(dat []int, length int) (bool, int) {
//	num := 0
//	right := make([]int, length)
//	for i:=0; i<length-1; i++ {
//		for j:=i+1; j<length; j++ {
//			if cb.color[dat[i]]&cb.color[dat[j]] > 0 {
//				//颜色相同
//				right[i]++
//				right[j]++
//			}
//		}
//	}
//	pos := 100
//	for i:=0; i<length; i++ {
//		if right[i] > 2 { // 至少有3张与其相同颜色， 即4张颜色相同
//			num++
//			if pos > i {
//				pos = i // 最大牌位置
//			}
//		}
//	}
//	if num > 1 { // 有至少两组是4张相同颜色
//		return true, pos
//	}
//
//	return false, -1
//}

func calOnes(dat int) (int, []int){
	ones := 0
	red := make([]int, 32)
	j:=0
	//for i:=31; i>=0; i-- {
	for i:=14; i>=2; i-- {
		if ((1<<uint(i)) & dat) > 0 {
			ones++
			red[j] = i
			j++
		}
	}
	return ones, red
}

func (cb *cardBuf) check4ColorSame() (bool, int, []int) {
	if cb.color[spades].count > 3 {
		return true, cb.color[spades].count, cb.color[spades].cards
	}
	if cb.color[hearts].count > 3 {
		return true, cb.color[hearts].count, cb.color[hearts].cards
	}
	if cb.color[clubs].count > 3 {
		return true, cb.color[clubs].count, cb.color[clubs].cards
	}
	if cb.color[diamonds].count > 3 {
		return true, cb.color[diamonds].count, cb.color[diamonds].cards
	}
	return false, 0, nil
}

func (cb *cardBuf) check5ColorSame() (bool, int, []int) {
	if cb.color[spades].count > 4 {
		return true, cb.color[spades].count, cb.color[spades].cards
	}
	if cb.color[hearts].count > 4 {
		return true, cb.color[hearts].count, cb.color[hearts].cards
	}
	if cb.color[clubs].count > 4 {
		return true, cb.color[clubs].count, cb.color[clubs].cards
	}
	if cb.color[diamonds].count > 4 {
		return true, cb.color[diamonds].count, cb.color[diamonds].cards
	}
	return false, 0, nil
}
//func (cb *cardBuf) check4ColorSame() (bool, int, []int) {
//	s, ss := calOnes(cb.color[spades])
//	if s > 3 {
//		return true, s, ss
//	}
//	d, dd := calOnes(cb.color[diamonds])
//	if d > 3 {
//		return true, d, dd
//	}
//	c, cc := calOnes(cb.color[clubs])
//	if c > 3 {
//		return true, c, cc
//	}
//	h, hh := calOnes(cb.color[hearts])
//	if h > 3 {
//		return true, h, hh
//	}
//	return false, 0, nil
//}

//func (cb *cardBuf) check5ColorSame() (bool, int, []int) {
//	s, ss := calOnes(cb.color[spades])
//	if s > 4 {
//		return true, s, ss
//	}
//	d, dd := calOnes(cb.color[diamonds])
//	if d > 4 {
//		return true, d, dd
//	}
//	c, cc := calOnes(cb.color[clubs])
//	if c > 4 {
//		return true, c, cc
//	}
//	h, hh := calOnes(cb.color[hearts])
//	if h > 4 {
//		return true, h, hh
//	}
//	return false, 0, nil
//}
// 搜索5张颜色相同
//func (cb *cardBuf) check5ColorSame(dat []int, length int) (bool, int) {
//	num := 0
//	right := make([]int, length)
//	for i:=0; i<length-1; i++ {
//		for j:=i+1; j<length; j++ {
//			//fmt.Println(dat[i], dat[j])
//			if cb.color[dat[i]]&cb.color[dat[j]] > 0 {
//				//颜色相同
//				right[i]++
//				right[j]++
//			}
//		}
//	}
//	//fmt.Println(cb.color)
//	//fmt.Println(right)
//	pos := 100
//	for i:=0; i<length; i++ {
//		if right[i] > 3 { // 至少有4张与其相同颜色， 即5张颜色相同
//			num++
//			if pos > i {
//				pos = i // 最大牌位置
//			}
//		}
//	}
//	if num > 1 { // 有至少两组是4张相同颜色
//		return true, pos
//	}
//
//	return false, -1
//}

func combineKey(dat ...int) string {
	length := len(dat)
	str := ""
	for i:=0;i<length;i++ {
		str += invTable[dat[i]]
	}
	return str
}

func (cb *cardBuf) checkSingleCards(ghost bool) (mode int, ret int) {
	dat := cb.s_order
	length := cb.s_cnt
	if length < 4 || (length == 4 && !ghost) {
		// 有赖子得满4张 || 无赖子得满5张
		return alone, dat[0]
	}

	// 颜色
	if ghost {
		if iSame, index, candidate := cb.check4ColorSame(); iSame {
			for i:=0; i<index-3; i++ {
				v, ok := cb.tabGhost[combineKey(candidate[i], candidate[i+1], candidate[i+2], candidate[i+3])]
				if ok {
					if candidate[i] == 14 && v == flush {
						// A5432
						return v, 5
					}
					if candidate[i] < 14 && v == flush && (candidate[i+3]-candidate[i] == 3) {
						// 赖子插入位置， 头部
						return v, candidate[i] + 1
					}
					return v, candidate[i]
				} else {
					return suit, candidate[i]
				}
			}
		}
	} else {
		if iSame, index, candidate := cb.check5ColorSame(); iSame {
			for i:=0; i<index-4; i++ {
				v, ok := cb.tabNoGhost[combineKey(candidate[i], candidate[i+1], candidate[i+2], candidate[i+3], candidate[i+4])]
				if ok {
					if candidate[i] == 14 && v == flush {
						// A5432
						return v, 5
					}
					return v, candidate[i]
				} else {
					return suit, candidate[i]
				}
			}
		}
	}

	// 颜色不同，是否顺子
	//hole := 0
	//pos := make([]int, length)
	//pos[0] = -1 // head
	//j:=1
	//for i:=0; i<length-1; i++ {
	//	if dat[i] != dat[i+1]+1 {
	//		hole++
	//		pos[j]=i
	//		j++
	//	}
	//}
	//pos[j]=length-1 // tail

	if ghost {
		// 有赖子
		//if pos[1] >=3 {
		//	//开头4张连续，插入赖子凑足5张
		//	if dat[0]<14 {
		//		ret=dat[0]+1
		//	} else {
		//		//开头A
		//		ret=dat[0]
		//	}
		//	mode = isRoyalorFlushorSeq(ret, sameColor)
		//	return
		//}
		for i:=0; i<length-3;i++ {
			v, ok := cb.tabGhost[combineKey(dat[i], dat[i+1], dat[i+2], dat[i+3])]
			if ok {
				if dat[i]==14 && v==flush {
					// A5432
					return sequence, 5
				}
				if v == royal {
					return sequence, 14
				}
				if dat[i]<14 && (dat[i+3] - dat[i] == 3) {
					// 赖子插入位置， 头部
					return sequence, dat[i]+1
				}
				return sequence, dat[i]
			}
		}
	} else {
		for i:=0; i<length-4;i++ {
			v, ok := cb.tabNoGhost[combineKey(dat[i], dat[i+1], dat[i+2], dat[i+3], dat[i+4])]
			if ok {
				if dat[i]==14 && v==flush {
					// A5432
					return sequence, 5
				}
				return sequence, dat[i]
			}
		}
	}

	return alone, dat[0]
}

func (cb *cardBuf) checkBomb(ghost bool) (bool, int) {
	if cb.card4 > 0 {
		return true, cb.card4
	}
	if ghost {
		if cb.t_cnt > 0 {
			return true, cb.t_order[0]
		}
	}
	return false, 0
}

func (cb *cardBuf) checkThreeTwo(ghost bool) (bool, int, int) {
	if cb.t_cnt > 0  {
		z3 := cb.t_order[0]
		max := 0
		for i:=0; i<cb.d_cnt; i++ {
			if z3 != cb.d_order[i] {
				max = cb.d_order[i]
				break
			}
		}
		if ghost {
			//3+1+赖子
			for i:=0; i<cb.s_cnt; i++ {
				if cb.s_order[i] != z3 {
					if max < cb.s_order[i] {
						max = cb.s_order[0]
						break
					}
				}
			}
		}
		if max > 0 {
			return true, z3, max
		}
	} else {
		if ghost {
			// 2+2+赖子
			if cb.d_cnt > 1 {
				return true, cb.d_order[0], cb.d_order[1]
			}
		}
	}

	return false, 0, 0
}

func (cb *cardBuf) checkThree(ghost bool) (bool, int) {
	max := 0
	if cb.t_cnt > 0 {
		max = cb.t_order[0]
	}
	if ghost {
		for i:=0; i<cb.d_cnt; i++ {
			if max < cb.d_order[i] {
				max = cb.d_order[i]
			}
		}
	}
	if max > 0 {
		return true, max
	}
	return false, 0
}

func (cb *cardBuf) check2Couple(ghost bool) (bool, int, int) {
	first := 0
	second := 0
	//fmt.Println(cb.d_order)
	if cb.d_cnt > 1 {
		first, second = cb.d_order[0], cb.d_order[1]
	} else if cb.d_cnt > 0 {
		first = cb.d_order[0]
	}
	if ghost {
		for i:=0; i<cb.s_cnt; i++ {
			if cb.s_order[i] > first {
				first, second = cb.s_order[i], first
				break
			} else if cb.s_order[i] > second {
				second = cb.s_order[i]
				break
			} else {
				break
			}
		}
	}
	if first > 0 && second > 0 {
		return true, first, second
	}
	return false, 0, 0
}

func (cb *cardBuf) checkCouple(ghost bool) (bool, int) {
	max := 0
	if cb.d_cnt > 0 {
		max = cb.d_order[0]
	}
	if ghost {
		if cb.s_order[0] > max {
			return true, cb.s_order[0]
		}
	}
	if max > 0 {
		return true, max
	}
	return false, 0
}

func (cb *cardBuf) checkType(ghost bool) (int, []int) {
	m2, v2 := cb.checkBomb(ghost)
	if m2 {
		ret := []int{v2, 0}
		for i:=0; i<len(cb.s_order); i++ {
			if cb.s_order[i] != v2 {
				ret[1]=cb.s_order[i]
				return four, ret
			}
		}
	}

	m3, v31, v32 := cb.checkThreeTwo(ghost)
	if m3 {
		return threeTwo, []int{v31, v32}
	}

	// full version
	m1, v1 := cb.checkSingleCards(ghost)
	if m1 < four {
		return m1, []int{v1}
	}
	// 5 cards without ghost
	//m1, v1 := cb.check5SingleCardsOnlyWithoutGhost()
	//if m1 < four {
	//	return m1, []int{v1}
	//}

	if m1 == suit {
		ret := []int{v1, 0, 0, 0, 0}
		j := 1
		for i:=0; i<len(cb.s_order); i++ {
			if v1 != cb.s_order[i] {
				ret[j] = cb.s_order[i]
				j++
				if j==5 {
					return m1, ret
				}
			}
		}
	}
	if m1 == sequence {
		return m1, []int{v1}
	}

	m4, v4 := cb.checkThree(ghost)
	if m4 {
		ret := []int{v4, 0, 0}
		j := 1
		for i:=0; i<len(cb.s_order); i++ {
			if cb.s_order[i] != v4 {
				ret[j] = cb.s_order[i]
				j++
				if j== 3 {
					return three, ret
				}
			}
		}
	}

	m5, v51, v52 := cb.check2Couple(ghost)
	if m5 {
		ret := []int{v51, v52, 0}
		for i:=0; i<len(cb.s_order); i++ {
			if cb.s_order[i] != v51 && cb.s_order[i] != v52 {
				ret[2] = cb.s_order[i]
				return couple2, ret
			}
		}
	}

	m6, v6 := cb.checkCouple(ghost)
	if m6 {
		ret := []int{v6, 0, 0, 0}
		j := 1
		for i:=0; i<len(cb.s_order); i++ {
			if cb.s_order[i] != v6 {
				ret[j] = cb.s_order[i]
				j++
				if j== 4 {
					return couple, ret
				}
			}
		}
	}

	ret := []int{0, 0, 0, 0, 0}
	for i:=0; i<5; i++ {
		ret[i] = cb.s_order[i]
	}

	if ghost {
		ret[0] = 14
	}
	return alone, ret
}

func TestAdd(cb1, cb2 *cardBuf, cards1, cards2 string) {
	length := len(cards1)
	//ghost1 := false
	//ghost2 := false

		//if string(cards1[i]) == "X" {
		//	ghost1 = true
		//	continue
		//}
		//if string(cards2[i]) == "X" {
		//	ghost2 = true
		//	continue
		//}
		cb1.addCard(cards1, length)
		cb2.addCard(cards2, length)

	cb1.clear()
	cb2.clear()
}

func Process(cb1, cb2 *cardBuf, cards1, cards2 string) (int, []int, int, []int) {
	length := len(cards1)
	ghost1 := cb1.addCard(cards1, length)
	ghost2 := cb2.addCard(cards2, length)

	//bubbleSort(cb1.t_order, cb1.t_cnt)
	sortCard(cb1.d_order, cb1.d_cnt)
	sortCard(cb1.s_order, cb1.s_cnt)
	//bubbleSort(cb2.t_order, cb2.t_cnt)
	sortCard(cb2.d_order, cb2.d_cnt)
	sortCard(cb2.s_order, cb2.s_cnt)

	mode1, v1 := cb1.checkType(ghost1)
	cb1.clear()
	mode2, v2 := cb2.checkType(ghost2)
	cb2.clear()
	return mode1, v1, mode2, v2
}

func (cb *cardBuf) process(cards string) (int, []int) {
	length := len(cards)
	ghost := cb.addCard(cards, length)
	//fmt.Println(cb.s_order, cb.d_order, cb.t_order)
	sortCard(cb.d_order, cb.d_cnt)
	sortCard(cb.s_order, cb.s_cnt)
	//fmt.Println(cb.s_order, cb.d_order, cb.t_order)
	mode, v := cb.checkType(ghost)
	cb.clear()
	return mode, v
}

func compareResult(mode1 int, v1 []int, mode2 int, v2 []int) int {
	if mode1 < mode2 {
		return over
	} else if mode1 > mode2 {
		return less
	} else {
		switch mode1 {
		case royal:
			return equal
		//case flush:
		//	if v1[0]  > v2[0] {
		//		return over
		//	} else if v1[0] < v2[0] {
		//		return less
		//	} else {
		//		return equal
		//	}
		//case four:
		//	if v1[0]  > v2[0] {
		//		return over
		//	} else if v1[0] < v2[0] {
		//		return less
		//	} else {
		//		if v1[1] > v2[1] {
		//			return over
		//		} else if v1[1] < v2[1] {
		//			return less
		//		}
		//		return equal
		//	}
		//case threeTwo:
		//
		//case suit:
		//
		//case sequence:
		//
		//case three:
		//
		//case couple2:
		//
		//case couple:

		default:
			for i:=0;i<len(v1);i++ {
				if v1[i] > v2[i] {
					return over
				} else if v1[i] < v2[i] {
					return less
				}
			}
		}
	}
	return equal
}

func (cb *cardBuf) check5SingleCardsOnlyWithoutGhost()  (mode int, ret int) {
	dat := cb.s_order
	length := cb.s_cnt
	if length != 5 {
		// 无赖子得满5张
		return alone, dat[0]
	}
	if iSame, _, candidate := cb.check5ColorSame(); iSame {
			v, ok := cb.tabFullNoGhost[combineKey(candidate[0], candidate[1], candidate[2], candidate[3], candidate[4])]
			if ok {
				return v.mode, v.max[0]
			} else {
				return suit, candidate[0]
			}
	}
	v, ok := cb.tabFullNoGhost[combineKey(dat[0], dat[1], dat[2], dat[3], dat[4])]
	if ok {
		return sequence, v.max[0]
	}
	return alone, dat[0]
}


type SimpleCards struct {
	cards seqCards
	buf []int
	color []int
	table5 map[string]seqCards
}

func newSimpleCards() *SimpleCards {
	ret := &SimpleCards{}
	ret.buf = make([]int, 5)
	ret.table5 = createFullTableWithoutGhost()
	create32Table(ret.table5)
	create4Table(ret.table5)
	create3Table(ret.table5)
	createCouple2Table(ret.table5)
	createCoupleTable(ret.table5)
	return ret
}

func (cb *SimpleCards) add5Card(cards string, se []int) {
	cb.color = se

	if v, ok := cb.table5[cards]; ok {
		cb.cards = v
	} else {
		for i:=0; i<5; i++ {
			card := table[string(cards[i])]
			cb.buf[i]=card
		}
		sortCard(cb.buf, 5)
		cb.cards = seqCards{alone, cb.buf}
	}

}

func (cb *SimpleCards) checkColor() bool {
	if cb.color[spades] == 5 || cb.color[hearts] == 5 ||cb.color[diamonds] == 5 ||cb.color[clubs] == 5 {
		return true
	}
	return false
}

func (cb *SimpleCards) checkType() (int, []int) {
	dat := []int{}
	if cb.checkColor() {
		if cb.cards.mode == royal || cb.cards.mode == flush {
			return cb.cards.mode, append(dat, cb.cards.max...)
		}
		return suit, append(dat, cb.cards.max...)
	}

	if cb.cards.mode == royal || cb.cards.mode == flush {
		return sequence, append(dat, cb.cards.max...)
	}

	return cb.cards.mode, append(dat, cb.cards.max...)
}

func (cb *SimpleCards) process(hand string, se []int) (int, []int) {
	cb.add5Card(hand, se)
	//return 0, []int{}
	return cb.checkType()
}