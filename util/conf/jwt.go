package conf

type jWT struct {
	Secret string
}

func (j *jWT) SigningKey() []byte {
	return []byte(j.Secret)
}
