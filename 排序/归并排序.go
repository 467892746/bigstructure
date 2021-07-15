package main


func MergeSort(l *SqList)  {
	var b [maxsize]int
	Msort(&l.r, &b, 0, l.length - 1)
	l.r = b
}

func Msort(sr, tr1 *[maxsize]int, s, t int)  {
	var tr2 [maxsize]int
	if s == t {
		tr1[s] = sr[s]
	}else {
		m := (s + t) / 2
		Msort(sr, &tr2, s, m)
		Msort(sr, &tr2, m + 1, t)
		Merges(&tr2, tr1, s, m, t)
	}

}

func Merges(sr, tr *[maxsize]int, i, m, n int)  {
	j := m + 1
	k := i
	for i <=m && j <=n {
		if sr[i] < sr[j] {
			tr[k] = sr[i]
			i++
		}else {
			tr[k] = sr[j]
			j++
		}
		k++
	}
	if i <= m {
		for l := 0; l <= m -i; l ++ {
			tr[k + l] = sr[i + l]
		}
	} // 将剩余的SR[i..m]复制到TR
	if j <= n {
		for l := 0; l <= n -j; l ++ {
			tr[k + l] = sr[j + l]
		}
	} // 将剩余的SR[j..n]复制到TR
}
