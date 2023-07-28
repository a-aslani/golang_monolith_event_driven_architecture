package domain

type Utils interface {
	Hash(str string) (string, error)
	CompareHash(hash, str string) bool
}
