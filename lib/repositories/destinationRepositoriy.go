package repositories

type DestinationRepository interface {
	GetDestinationPlatforms(uuid string)
}
