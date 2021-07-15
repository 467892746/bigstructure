package main

// SelectSort 简单选择排序
func SelectSort(l *SqList)  {
	for i := 0; i <l.length; i++ {
		min := i
		for j := i + 1; j < l.length; j ++ {
			if l.r[min] > l.r[j] {
				min = j
			}
		}
		if min != i {
			Swap(l, i, min)
		}
	}
}


