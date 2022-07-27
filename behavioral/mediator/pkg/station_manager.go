package pkg

type StationManager struct {
	IsPlatformFree bool
	Queue          []Vehicle
}

func NewStationManager() *StationManager {
	return &StationManager{
		IsPlatformFree: true,
	}
}

func (s StationManager) CanArrive(vehicle Vehicle) bool {
	if s.IsPlatformFree {
		s.IsPlatformFree = false
		return true
	}
	s.Queue = append(s.Queue, vehicle)
	return false
}

func (s StationManager) NotifyAboutGo() {
	if !s.IsPlatformFree {
		s.IsPlatformFree = true
	}
	if len(s.Queue) > 0 {
		firstTrainInQueue := s.Queue[0]
		s.Queue = s.Queue[1:]
		firstTrainInQueue.PermitArrive()
	}
}
