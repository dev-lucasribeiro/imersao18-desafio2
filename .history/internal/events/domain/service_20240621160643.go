package domain

type spotService struct{}

func NewSpotService() *spotService {
	return &spotService{}
}


func (s *spotService) GenerateSpots(event *Event, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}