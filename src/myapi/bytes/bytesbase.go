package bytes

func writeUInt16(buff []byte, data uint16) {
	for i := 0; i < 2; i++ {
		buff[i] = byte(data >> uint(i * 8))
	}
}

//rune 等同于int32,常用来处理unicode或utf-8字符
func split(r rune) bool {
	if r == 'c' {
		return true
	}
	return false
}

