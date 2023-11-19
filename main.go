package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println("1 二分法, 2 反復法, 3 べき乗, 4 レイリー商")
	in := getCommandLineString()
	calcType, err := strconv.Atoi(strings.TrimRight(in, "\r"))
	if err != nil {
		fmt.Println(err)
		return
	}
	if calcType == 1 {
		var init_min float64
		var init_max float64
		var ipsi float64

		fmt.Println("init_min")
		in = getCommandLineString()
		init_min, err = strconv.ParseFloat(in, 64)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("init_max")
		in = getCommandLineString()
		init_max, err = strconv.ParseFloat(in, 64)
		fmt.Println("誤差")
		if err != nil {
			fmt.Println(err)
			return
		}

		in = getCommandLineString()
		ipsi, err = strconv.ParseFloat(in, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("答え: ", nibunhou(init_min, init_max, ipsi))
	} else if calcType == 2 {
		var init float64
		isAns := false
		var ans float64
		var ipsi float64

		fmt.Println("初期値")
		in = getCommandLineString()
		init, err = strconv.ParseFloat(in, 64)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("答えがあるか 1 ある, 2 ない")
		in = getCommandLineString()
		if in == "1" {
			isAns = true
		}

		if isAns {
			fmt.Println("答え")
			in = getCommandLineString()
			ans, err = strconv.ParseFloat(in, 64)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("誤差")
			in = getCommandLineString()
			ipsi, err = strconv.ParseFloat(in, 64)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		fmt.Println("答え: ", hanpukuhou(init, isAns, ans, ipsi))
	} else if calcType == 3 {
		var isAns bool
		var ans float64
		var ipsi float64
		fmt.Println("答えがあるか 1 ある, 2 ない")
		isAns = (getCommandLineString() == "1")
		if isAns {
			fmt.Println("答え")
			in = getCommandLineString()
			ans, err = strconv.ParseFloat(in, 64)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("誤差")
			in = getCommandLineString()
			ipsi, err = strconv.ParseFloat(in, 64)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("答え: ", bekizyou(true, ans, ipsi))
		} else {
			fmt.Println("答え: ", bekizyou(false, ans, ipsi))
		}

	} else if calcType == 4 {
		var isAns bool
		var ans float64
		var ipsi float64
		fmt.Println("答えがあるか 1 ある, 2 ない")
		isAns = (getCommandLineString() == "1")
		if isAns {
			fmt.Println("答え")
			in = getCommandLineString()
			ans, err = strconv.ParseFloat(in, 64)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("誤差")
			in = getCommandLineString()
			ipsi, err = strconv.ParseFloat(in, 64)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("答え: ", reiry(true, ans, ipsi))
		} else {
			fmt.Println("答え: ", reiry(false, ans, ipsi))
		}
	}
}

func getCommandLineString() string {
	in, err := bufio.NewReader(os.Stdin).ReadString('\r')
	if err != nil {
		fmt.Println(err)
		return ""
	}
	s := strings.ReplaceAll(in, "\r", "")
	s = strings.ReplaceAll(s, "\n", "")

	return s
}

func nibunhou(init_min float64, init_max float64, ipsi float64) float64 {
	a := init_min
	b := init_max

	for i := 0; i < 100; i++ {
		mid := (a + b) / 2
		fmt.Println(mid)

		midRes := Siki(mid)
		aRes := Siki(a)
		bRes := Siki(b)

		if midRes*aRes > 0 {
			a = mid
		} else if midRes*bRes > 0 {
			b = mid
		} else {
			fmt.Println("calc error")
			return 0
		}

		if b-a < ipsi {
			fmt.Printf("nibun success %d + 1回\n", i)
			return b - a/2
		}
		time.Sleep(time.Second)

	}
	return 0
}

func hanpukuhou(init float64, isAns bool, ans float64, ipsi float64) float64 {
	x := init
	if isAns {
		for i := 0; i < 100; i++ {
			fmt.Println(x)
			x = Siki(x)
			if ans-x < ipsi || x-ans < ipsi {
				fmt.Printf("hanpuku success %d + 1回\n", i)
				return x
			}
			time.Sleep(time.Second)
		}
	} else {
		for i := 0; i < 20; i++ {
			fmt.Println(x)
			x = Siki(x)
			time.Sleep(time.Second)
		}
		return x
	}

	return 0
}

func reiry(isAns bool, ans float64, ipsi float64) float64 {
	initMatrix := [][]float64{
		{1},
		{0},
	}

	aMatrix := [][]float64{
		{4, 1},
		{1, 0},
	}

	xMatrix := make([][]float64, len(aMatrix))
	xRegMatrix := make([][]float64, len(aMatrix))

	for i := 0; i < len(aMatrix); i++ {
		xMatrix[i] = make([]float64, len(initMatrix[0]))
		xRegMatrix[i] = make([]float64, len(initMatrix[0]))
		copy(xMatrix[i], initMatrix[i])
		copy(xRegMatrix[i], initMatrix[i])
	}

	if isAns {
		for i := 0; i < 100; i++ {
			//xMatrixを初期化
			for j := 0; j < len(xMatrix); j++ {
				for k := 0; k < len(xMatrix[0]); k++ {
					xMatrix[j][k] = 0
				}
			}
			//xRegMatrixとaMatrixの積を求める
			for j := 0; j < len(aMatrix); j++ {
				for k := 0; k < len(xRegMatrix[0]); k++ {
					for l := 0; l < len(xRegMatrix); l++ {
						xMatrix[j][k] += aMatrix[j][l] * xRegMatrix[l][k]
					}
				}
			}

			//xMatrixのユークリッドノルムを求める
			r := 0.0
			for j := 0; j < len(xMatrix); j++ {
				for k := 0; k < len(xMatrix[0]); k++ {
					r += pow(xMatrix[j][k], 2)
				}
			}
			r = math.Sqrt(r)

			if i > 0 {
				ue := 0.0
				for j := 0; j < len(xRegMatrix); j++ {
					for k := 0; k < len(xRegMatrix[0]); k++ {
						ue += xRegMatrix[j][k] * xMatrix[j][k]
					}
				}
				if math.Abs(ue) < ipsi {
					fmt.Printf("reiry success %d 回\n", i+1)
					return ue
				}
			}

			//xMatrixをrで割る
			for i := 0; i < len(xRegMatrix); i++ {
				for j := 0; j < len(xRegMatrix[0]); j++ {
					xRegMatrix[i][j] = xMatrix[i][j] / r
				}
			}

			fmt.Println(xRegMatrix)
			time.Sleep(time.Second)
		}
		fmt.Printf("over")
	} else {
		ue := 0.0
		for i := 0; i < 100; i++ {
			//xMatrixを初期化
			for j := 0; j < len(xMatrix); j++ {
				for k := 0; k < len(xMatrix[0]); k++ {
					xMatrix[j][k] = 0
				}
			}
			//xRegMatrixとaMatrixの積を求める
			for j := 0; j < len(aMatrix); j++ {
				for k := 0; k < len(xRegMatrix[0]); k++ {
					for l := 0; l < len(xRegMatrix); l++ {
						xMatrix[j][k] += aMatrix[j][l] * xRegMatrix[l][k]
					}
				}
			}

			//xMatrixのユークリッドノルムを求める
			r := 0.0
			for j := 0; j < len(xMatrix); j++ {
				for k := 0; k < len(xMatrix[0]); k++ {
					r += pow(xMatrix[j][k], 2)
				}
			}
			r = math.Sqrt(r)
			ue = 0.0
			if i > 0 {
				for j := 0; j < len(xRegMatrix); j++ {
					for k := 0; k < len(xRegMatrix[0]); k++ {
						ue += xRegMatrix[j][k] * xMatrix[j][k]
					}
				}
				fmt.Println(ue)
			}

			//xMatrixをrで割る
			for i := 0; i < len(xRegMatrix); i++ {
				for j := 0; j < len(xRegMatrix[0]); j++ {
					xRegMatrix[i][j] = xMatrix[i][j] / r
				}
			}

			fmt.Println(xRegMatrix)
			time.Sleep(time.Second)
		}
		return ue
	}
	return 0
}

func bekizyou(isAns bool, ans float64, ipsi float64) float64 {
	initMatrix := [][]float64{
		{1},
		{0},
	}

	aMatrix := [][]float64{
		{4, 1},
		{1, 0},
	}

	xMatrix := make([][]float64, len(aMatrix))
	xRegMatrix := make([][]float64, len(aMatrix))

	for i := 0; i < len(aMatrix); i++ {
		xMatrix[i] = make([]float64, len(initMatrix[0]))
		xRegMatrix[i] = make([]float64, len(initMatrix[0]))
		copy(xMatrix[i], initMatrix[i])
		copy(xRegMatrix[i], initMatrix[i])
	}

	if isAns {
		for i := 0; i < 100; i++ {
			//xMatrixを初期化
			for j := 0; j < len(xMatrix); j++ {
				for k := 0; k < len(xMatrix[0]); k++ {
					xMatrix[j][k] = 0
				}
			}
			//xRegMatrixとaMatrixの積を求める
			for i := 0; i < len(aMatrix); i++ {
				for j := 0; j < len(xRegMatrix[0]); j++ {
					for k := 0; k < len(xRegMatrix); k++ {
						xMatrix[i][j] += aMatrix[i][k] * xRegMatrix[k][j]
					}
				}
			}
			//xRegMatrixの最大値
			max := 0.0
			disty := 0
			distx := 0
			for i := 0; i < len(xRegMatrix); i++ {
				for j := 0; j < len(xRegMatrix[0]); j++ {
					if max < xRegMatrix[i][j] {
						max = xRegMatrix[i][j]
						disty = i
						distx = j
					}
				}
			}
			r := xMatrix[disty][distx] / max

			if (r-ans) < ipsi || (ans-r) < ipsi {
				fmt.Printf("bekizyou success %d 回\n", i+1)
				return r
			}

			//xMatrixをrで割る
			for i := 0; i < len(xRegMatrix); i++ {
				for j := 0; j < len(xRegMatrix[0]); j++ {
					xRegMatrix[i][j] = xMatrix[i][j] / r
				}
			}

			fmt.Println(xRegMatrix)
			time.Sleep(time.Second)
		}
		fmt.Printf("over")
	} else {
		r := 0.0
		for i := 0; i < 20; i++ {
			//xMatrixを初期化
			for i := 0; i < len(xMatrix); i++ {
				for j := 0; j < len(xMatrix[0]); j++ {
					xMatrix[i][j] = 0
				}
			}
			//aMatrixとxRegMatrixの積を求める
			for i := 0; i < len(aMatrix); i++ {
				for j := 0; j < len(xRegMatrix[0]); j++ {
					for k := 0; k < len(xRegMatrix); k++ {
						xMatrix[i][j] += aMatrix[i][k] * xRegMatrix[k][j]
					}
				}
			}

			//xRegMatrixの最大値の位置を求める
			max := 0.0
			disty := 0
			distx := 0
			for i := 0; i < len(xRegMatrix); i++ {
				for j := 0; j < len(xRegMatrix[0]); j++ {
					if max < xRegMatrix[i][j] {
						max = xRegMatrix[i][j]
						disty = i
						distx = j
					}
				}
			}
			r = xMatrix[disty][distx] / max

			//xMatrixをrで割る
			for i := 0; i < len(xRegMatrix); i++ {
				for j := 0; j < len(xRegMatrix[0]); j++ {
					xRegMatrix[i][j] = xMatrix[i][j] / r
				}
			}
			fmt.Println("r: ", r)
			fmt.Println(xRegMatrix)
			time.Sleep(time.Second)
		}
		return r
	}
	return 0
}

func Siki(x float64) float64 {
	//入力する式、反復法のときは変形した式を入力
	return x/2 + 1/x
}

func pow(x float64, n int) float64 {
	res := 1.0
	for i := 0; i < n; i++ {
		res *= x
	}
	return res
}

// func Siki(s string, para float64) (float64, error) {
// 	expression, err := govaluate.NewEvaluableExpression(s)
// 	if err != nil {
// 		fmt.Println(err)
// 		return 0, err
// 	}

// 	paras := map[string]interface{}{
// 		"x": para,
// 	}

// 	result, err := expression.Evaluate(paras)
// 	if err != nil {
// 		fmt.Println(err)
// 		return 0, err
// 	}

// 	return result.(float64), nil
// }
