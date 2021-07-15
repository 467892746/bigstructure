package main

func HeapSort(l *SqList)  {
	for i := (l.length - 1) / 2; i >= 0; i-- {
		HeapAdjust(l, i, l.length - 1)
	}
	for i := l.length - 1; i >=0; i-- {
		Swap(l, i, 0)
		HeapAdjust(l, 0, i - 1 )
	}
}

func HeapAdjust(l *SqList, s, m int)  {
	temp := l.r[s]
	for j := 2 * s + 1; j <= m; j = j * 2 +1 {
		if j < m && l.r[j] < l.r[j + 1]{
			j++
		}
		if temp >= l.r[j] {
			break
		}
		l.r[s] = l.r[j]
		s = j
	}
	l.r[s] = temp
}