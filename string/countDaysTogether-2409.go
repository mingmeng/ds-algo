package string

var days = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {

	// leave Alice>= arriveBob >=arrive Alice
	// 先判断两个日子是否有相交
	// 条件为 bob 的某一个日期落在了alice的日期区间内
	dateProcess(arriveAlice)

	return 0
}

func dateProcess(date string) Date {
	m := int(date[0]-'0')*10 + int(date[1]-'0')
	d := int(date[3]-'0')*10 + int(date[4]-'0')
	return Date{
		Month: m,
		Day:   d,
	}
}

type Date struct {
	Month int
	Day   int
}

// d - date
func (d Date) Sub(date Date) int {
	if d.Month == date.Month {
		return d.Day - date.Day
	} else if d.Month > date.Month {
		var subs = d.Day

		for i := date.Month; i < d.Month-1; i++ {
			subs += days[i]
		}

		subs += days[date.Month-1] - date.Day

		return subs
	} else {
		var subs = date.Day

		for i := d.Month; i < date.Month-1; i++ {
			subs += days[i]
		}

		subs += days[d.Month-1] - d.Day
		return subs
	}
}

// -1 means small like d = 1.1 date =  12.1 will output it
// 0 means equal like d= 1.1 date = 1.1 will output it
// 1 means large like d = 12.1 date = 1.1 will output it
func (d Date) Compare(date Date) int {
	if date.Month == d.Month && date.Day == d.Day {
		return 0
	} else if date.Month > d.Month {
		return -1
	} else if d.Month > date.Month {
		return 1
	} else {
		if d.Day-date.Day > 0 {
			return 1
		} else {
			return -1
		}
	}
}
