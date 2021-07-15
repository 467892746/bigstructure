package main

func InsertSort(l *SqList)  {
	var sentry int
	for i := 1; i < l.length; i++ {
			sentry = l.r[i]
			var j int
			for j = i -1; l.r[j] > sentry && j >= 0; j-- {
				l.r[j + 1]= l.r[j]
			}
		if j + 1 !=  i {
			l.r[j + 1] = sentry

		}
	}
}
