package sensors

type UserPairEvent struct {
	BaseEvent
}

func (e UserPairEvent) Name() string {
	return "track_signup"
}

func (e UserPairEvent) Properties() map[string]interface{} {
	return Struct2Map(e)
}

