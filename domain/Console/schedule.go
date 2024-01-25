package Console

type Schedule struct {
	ScheduleMap ScheduleMap
}

func NewSchedule() *Schedule {
	schedule := Schedule{}
	scheduleMap := make(ScheduleMap, 0)
	schedule.ScheduleMap = scheduleMap
	return &schedule
}

func (s *Schedule) AddSpec(handle ScheduleFunc) {
	s.ScheduleMap["spec"] = handle
}

func (s *Schedule) AddFlags(handle ScheduleFunc) {
	s.ScheduleMap["flags"] = handle
}

func (s *Schedule) AddFn(handle ScheduleFunc) {
	s.ScheduleMap["fn"] = handle
}
