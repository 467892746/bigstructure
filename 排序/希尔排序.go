package main

func ShellSort(s *SqList)  {
	var sentry,j int

	increment := s.length
	for increment > 1 {
		increment = increment / 3 + 1
		for i := increment ; i < s.length; i++ {
			sentry = s.r[i]
			for j = i - increment;   j >= 0 && sentry < s.r[j]; j-= increment {
				s.r[j + increment] = s.r[j]
			}
			if j + increment !=  i {
				s.r[j + increment] = sentry

			}
		}
	}
}
