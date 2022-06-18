package scheduler

import "time"

type user struct {
	Hours int
}

func main() {
	u := user{
		Hours: 0,
	}

	sch := NewScheduler(time.Hour*1, u)

	sch.SetMain(func(u *user) error {
		u.Hours++
		return nil
	})

	go sch.Run()

	defer sch.Close()
}
