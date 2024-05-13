package pets

type Pet interface {
	Feed(food string) string
	GiveAttention(activity string) string
	IsHungry() bool
}
