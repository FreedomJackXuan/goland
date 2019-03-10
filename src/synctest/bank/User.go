package bank

type Bank struct {
	Name string
	Passwd string
}

func (b *Bank) SetName(name string)  {
	b.Name = name
}
func (b * Bank) GetName() string {
	return b.Name
}
func (b *Bank) SetPasswd(passwd string)  {
	b.Passwd = passwd
}
func (b *Bank) GetPasswd() string{
	return b.Passwd
}

