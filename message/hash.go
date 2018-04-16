package message

type Hash [32]uint8

func ZeroHash() Hash {
	var h [32]uint8
	return h
}

func GenesisHash() Hash {
	return [32]uint8{51, 70, 246, 60, 245, 178, 94, 20, 173, 221, 136, 85, 226, 117, 87, 132, 229, 94, 97, 44, 213, 133, 97, 200, 202, 24, 215, 207, 108, 168, 111, 75}
}
