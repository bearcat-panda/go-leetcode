package __Z字形变换

import (
	"fmt"
	"testing"
)

//将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
//
// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
//
//
//P   A   H   N		1 5 9 13			| 4 (3+3-2)
//A P L S I I G     2 4 6 8 10 12 14	| 3
//Y   I   R			3 7 11				|4
//
// 计算公式
//14/(3+3/2) * 2 + 14%(3+3/2) > 4
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
//
// 请你实现这个将字符串进行指定行数变换的函数：
//
//
//string convert(string s, int numRows);
//
//
//
// 示例 1：
//
//
//输入：s = "PAYPALISHIRING", numRows = 3
//输出："PAHNAPLSIIGYIR"
//
//示例 2：
//
//
//输入：s = "PAYPALISHIRING", numRows = 4
//输出："PINALSIGYAHRPI"
//解释：
//P     I    N		1 7 13			|6 (4+4-2)
//A   L S  I G		2 6 8 12 14     |4 2 4 2
//Y A   H R			3 5 9 11		|2 4 2
//P     I			4 10			|5
//
//
// 示例 3：
//
//
//输入：s = "A", numRows = 1
//输出："A"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 由英文字母（小写和大写）、',' 和 '.' 组成
// 1 <= numRows <= 1000
//
// Related Topics 字符串
// 👍 983 👎 0

func TestConvert(t *testing.T) {
	var s = "PAYPALISHIRING"
	fmt.Println(convertTwo(s, 3), "PAHNAPLSIIGYIR", convertTwo(s, 3) == "PAHNAPLSIIGYIR")
	fmt.Println(convertThree(s, 3), "PAHNAPLSIIGYIR", convertThree(s, 3) == "PAHNAPLSIIGYIR")
	fmt.Println(convertTwo(s, 4), "PINALSIGYAHRPI", convertTwo(s, 4) == "PINALSIGYAHRPI")
	s = "A"
	fmt.Println(convertTwo(s, 1))
}

//创建一个对应的数组. 然后把数据填充进去
//遍历数组. 取出数据
// 计算公式
//14/(3+3/2) * 2 + 14%(3+3/2) > 4
//行数为一维数组的个数 numRows
//每一个一维数组的个数为
/**
字符串长度length
行数为row
因为是Z. 一个不完整V 左转90度 就是2列 v = (row+row/2)
完整的列数 length / v * (row-1)
不完整的列数 length % v > row ? 1:0
*/
//写入
/**
0 0
1 0
2 0
...
row 0

row/2 1

0 2
1 2
2 2
...
row 2

row/2 3

0 4
1 4
2 4
...
row 4

*/
func convertOne(s string, numRows int) string {
	//不完整的v的长度
	v := (numRows + numRows/2)
	length := len(s)
	//完整列的长度
	columnComplete := length / v * (numRows - 1)
	//不完整列的长度
	var columnNotComplete int
	if length%v > numRows {
		columnNotComplete = 1
	} else {
		columnNotComplete = 2
	}

	//总长度
	columnComplete += columnNotComplete

	/*var array = make([][]int, numRows)
	for i, v := range array {
		for j, _ := range v {

		}
	}*/

	return s

}

//第一种方法太浪费时间了
//找规律
//numRows=3
//P   A   H   N		1   5 	9 	  13			| 4 (3+3-2)
//A P L S I I G     2 4 6 8 10 12 14			| 2
//Y   I   R			3 	7 	11					|4
//numRows=4
//P     I    N		1 	   7   	  13			|6 (4+4-2)
//A   L S  I G		2 	6 8    12 14     	|4 2 4 2
//Y A   H R			3 5   9  11				|2 4 2
//P     I			4     10				|5

//numRows=5
//P	    H			1       9
//A   S	I			2     8 10
//Y  I  R			3 	7   11
//P	L   I G			4 6	    12	14
//A     N			5       13

func convertTwo(s string, numRows int) string {
	//如果是两行最多也就是 length/2列
	//可以不计算精确的列数. 就按最大的给
	if numRows == 1 {
		return s
	}

	//字符串长度
	length := len(s)
	//用来存放字符串
	var array = make([][]uint8, numRows) //一位数组
	for i, v := range array {            //二维数组
		v = make([]uint8, length/2)
		array[i] = v
	}

	var count = 0
	for i := 0; i < length/2; i++ { //按位置进行赋值
		for j := 0; j < numRows; j++ {
			if count > length-1 {
				break
			}
			//从上到下
			if i%(numRows-1) == 0 {
				array[j][i] = s[count]
				count++
			} else { //从下到上
				x := (numRows - 1) - (count % (numRows - 1))
				array[x][i] = s[count]
				count++
				break
			}
		}
	}

	var sArray = []uint8{} //按顺序遍历数组里面的值
	for _, v := range array {
		for _, v1 := range v {
			if v1 != 0 {
				sArray = append(sArray, v1)
			}
		}
	}

	return string(sArray)
}

func convertThree(s string, numRows int) string {
	// 不满足，提前返回
	if len(s) <= 2 || numRows == 1 {
		return s
	}
	// 保存最终结果
	var res string
	// 保存每次的结果
	var tmp = make([]string, numRows)
	// 初始位置
	curPos := 0
	// 用来标记是否该转向了
	shouldTurn := -1
	// 遍历每个元素
	for _, val := range s {
		// 添加进tmp里面，位置为curPos
		tmp[curPos] += string(val)
		// 如果走到头或者尾，就该转向了
		// 因为就在numRows的长度里面左右震荡
		if curPos == 0 || curPos == numRows-1 {
			// 转向
			shouldTurn = -shouldTurn
		}
		// curPos 用来标记tmp里面我们该去哪个位置
		curPos += shouldTurn
	}
	// 这时tmp里面已经保存了数据了，我们只需要转换一下输出格式
	for _, val := range tmp {
		res += val
	}
	// 最后的输出
	return res
}

/*性能什么的就先放一边，理解是第一位。
清晰思路，不搞看不懂的玄学操作。

numRows = 3
tmp[0]:LCIR
tmp[1]:ETOESIIG
tmp[2]:EDHN
来回上下扫，所以tmp[0]先放L，tmp[1]放E，tmp[2]放E
这时候，扫到底了，往上扫，tmp[1]放T，tmp[0]放C
这时候，扫到顶了，往下扫，tmp[1]放O，tmp[2]放D
这时候，扫到底了，继续。。。
填充的顺序就像cos函数的正数部分*/

func convertFour(s string, numRows int) string {
	if len(s) < 2 || numRows == 1 {
		return s
	}

	//最终结果
	var res string
	//每行的结果
	var tmp = make([]string, numRows)
	//字符初始位置
	current := 0
	//转向后是+还是- 默认为-
	turn := -1

	for _, v := range s {
		tmp[current] = string(v) //赋值

		//判断是否需要转向
		if current == 0 || current == numRows-1 {
			turn = -turn
		}

		current += turn
	}

	// 这时tmp里面已经保存了数据了，我们只需要转换一下输出格式
	for _, val := range tmp {
		res += val
	}
	// 最后的输出
	return res

}
