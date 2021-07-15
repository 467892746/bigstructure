package main

//BubbleSort0 冒泡排序初级版本(交换排序)
func BubbleSort0(l *SqList)  {
	for i := 0; i < l.length; i ++ {
		for j := i +1; j < l.length; j ++ {
			if l.r[i] > l.r[j] {
				Swap(l, i, j)
			}
		}
	}
}

//BubbleSort1  冒泡排序
func BubbleSort1(l *SqList)  {
	for i := 0; i < l.length; i ++ {
		for j := l.length - 2; j >= i; j -- {
			if l.r[j] > l.r[j + 1] {
				Swap(l, j, j + 1)
			}
		}
	}
}

// BubbleSort2 冒泡排序优化
func BubbleSort2(l *SqList)  {
	flag := true
	for i := 0; i < l.length - 1 && flag; i++ {
		flag = false
		for j := l.length - 2; j >= i; j -- {
			if l.r[j] > l.r[j + 1] {
				Swap(l, j, j + 1)
				flag = true
			}
		}
	}
}


