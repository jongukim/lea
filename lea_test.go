package lea

import "testing"

func TestBa2w(t *testing.T) {
	w := ba2w([4]byte{0x0f, 0x1e, 0x2d, 0x3c})
	if w != 0x3c2d1e0f {
		t.Errorf("Should be 0x3c2d1e0f. Your answer was %v.", w)
	}
}

func TestW2ba(t *testing.T) {
	ba := w2ba(0x3c2d1e0f)
	if ba[0] != 0x0f || ba[1] != 0x1e || ba[2] != 0x2d || ba[3] != 0x3c {
		t.Errorf("Should be {0x0f, 0x1e, 0x2d, 0x3c}. Your answer was %v.", ba)
	}
}

func TestRol(t *testing.T) {
	cases := []struct {
		w   word
		r   uint
		ans word
	}{
		{
			0x12345678,
			4,
			0x23456781,
		},
		{
			0x12345678,
			8,
			0x34567812,
		},
		{
			0x12345678,
			16,
			0x56781234,
		},
		{
			0x12345678,
			24,
			0x78123456,
		},
		{
			0xffffffff,
			1,
			0xffffffff,
		},
		{
			0x01010101,
			0,
			0x01010101,
		},
		{
			0x12345678,
			36,
			0x23456781,
		},
	}
	for _, c := range cases {
		u := rol(c.w, c.r)
		if u != c.ans {
			t.Errorf("ROL(%v, %v) != %v (your answer was %v)", c.w, c.r, c.ans, u)
		}
	}
}

func TestRor(t *testing.T) {
	cases := []struct {
		w   word
		r   uint
		ans word
	}{
		{
			0x12345678,
			4,
			0x81234567,
		},
		{
			0x12345678,
			8,
			0x78123456,
		},
		{
			0x12345678,
			16,
			0x56781234,
		},
		{
			0x12345678,
			24,
			0x34567812,
		},
		{
			0xffffffff,
			1,
			0xffffffff,
		},
		{
			0x01010101,
			0,
			0x01010101,
		},
		{
			0x12345678,
			36,
			0x81234567,
		},
	}
	for _, c := range cases {
		u := ror(c.w, c.r)
		if u != c.ans {
			t.Errorf("ROR(%v, %v) != %v (your answer was %v)", c.w, c.r, c.ans, u)
		}
	}
}

func TestEncRoundKey(t *testing.T) {
	ans128 := [][6]word{
		{0x003a0fd4, 0x02497010, 0x194f7db1, 0x02497010, 0x090d0883, 0x02497010},
		{0x11fdcbb1, 0x9e98e0c8, 0x18b570cf, 0x9e98e0c8, 0x9dc53a79, 0x9e98e0c8},
		{0xf30f7bb5, 0x6d6628db, 0xb74e5dad, 0x6d6628db, 0xa65e46d0, 0x6d6628db},
		{0x74120631, 0xdac9bd17, 0xcd1ecf34, 0xdac9bd17, 0x540f76f1, 0xdac9bd17},
		{0x662147db, 0xc637c47a, 0x46518932, 0xc637c47a, 0x23269260, 0xc637c47a},
		{0xe4dd5047, 0xf694285e, 0xe1c2951d, 0xf694285e, 0x8ca5242c, 0xf694285e},
		{0xbaf8e5ca, 0x3e936cd7, 0x0fc7e5b1, 0x3e936cd7, 0xf1c8fa8c, 0x3e936cd7},
		{0x5522b80c, 0xee22ca78, 0x8a6fa8b3, 0xee22ca78, 0x65637b74, 0xee22ca78},
		{0x8a19279e, 0x6fb40ffe, 0x85c5f092, 0x6fb40ffe, 0x92cc9f25, 0x6fb40ffe},
		{0x9dde584c, 0xcb00c87f, 0x4780ad66, 0xcb00c87f, 0xe61b5dcb, 0xcb00c87f},
		{0x4fa10466, 0xf728e276, 0xd255411b, 0xf728e276, 0x656839ad, 0xf728e276},
		{0x9250d058, 0x51bd501f, 0x1cb40dae, 0x51bd501f, 0x1abf218d, 0x51bd501f},
		{0x21dd192d, 0x77c644e2, 0xcabfaa45, 0x77c644e2, 0x681c207d, 0x77c644e2},
		{0xde7ac372, 0x9436afd0, 0x10331d80, 0x9436afd0, 0xf326fe98, 0x9436afd0},
		{0xfb3ac3d4, 0x93df660e, 0x2f65d8a3, 0x93df660e, 0xdf92e761, 0x93df660e},
		{0x27620087, 0x265ef76e, 0x4fb29864, 0x265ef76e, 0x2656ed1a, 0x265ef76e},
		{0x227b88ec, 0xd0b3fa6f, 0xc86a08fd, 0xd0b3fa6f, 0xa864cba9, 0xd0b3fa6f},
		{0xf1002361, 0xe5e85fc3, 0x1f0b0408, 0xe5e85fc3, 0x488e7ac4, 0xe5e85fc3},
		{0xc65415d5, 0x51e176b6, 0xeca88bf9, 0x51e176b6, 0xedb89ece, 0x51e176b6},
		{0x9b6fb99c, 0x0548254b, 0x8de9f7c2, 0x0548254b, 0xb6b4d146, 0x0548254b},
		{0x7257f134, 0x06051a42, 0x36bcef01, 0x06051a42, 0xb649d524, 0x06051a42},
		{0xa540fb03, 0x34b196e6, 0xf7c80dad, 0x34b196e6, 0x71bc7dc4, 0x34b196e6},
		{0x8fbee745, 0xcf744123, 0x907c0a60, 0xcf744123, 0x8215ec35, 0xcf744123},
		{0x0bf6adba, 0xdf69029d, 0x5b72305a, 0xdf69029d, 0xcb47c19f, 0xdf69029d},
	}
	ans192 := [][6]word{
		{0x003a0fd4, 0x02497010, 0x194f7db1, 0x090d0883, 0x2ff5805a, 0xc2580b27},
		{0x11fdcbb1, 0x9e98e0c8, 0x18b570cf, 0x9dc53a79, 0x5c145788, 0x9771b5e5},
		{0xf30f7bb5, 0x6d6628db, 0xb74e5dad, 0xa65e46d0, 0x6f44da96, 0xf643115f},
		{0x74120631, 0xdac9bd17, 0xcd1ecf34, 0x540f76f1, 0xaa1a5bdb, 0xfbafaae7},
		{0x13f8a031, 0x34f28728, 0x31fdb409, 0x0e31481b, 0xdf498117, 0xcf9371f1},
		{0x0967c312, 0xb3484ec8, 0x3aae5b3d, 0x5a9714a0, 0xb2d4dd5f, 0x3a1fcdf7},
		{0x0ac47404, 0x59e9e54d, 0xa60dc00a, 0x566139d3, 0x898dce4f, 0x582d72dd},
		{0x77f3ea4c, 0xe2a73c8d, 0xb8f1249a, 0x6a172700, 0xbc0e539c, 0x2e46fdbb},
		{0xb4e0e98a, 0x3d028c05, 0xb8d3a050, 0xdbd67bef, 0xdf675c7a, 0x99eefbb0},
		{0xe68584f6, 0xce31ef45, 0x96c105ac, 0x2a1be677, 0x9d72b8b0, 0x33cecc54},
		{0xc22ffd76, 0x1ab7167e, 0x42bb3060, 0x7da517f5, 0x4aa0e8d3, 0x0a070c3c},
		{0xe200a765, 0xc2be17b3, 0x7f22543f, 0x3e4eb7a1, 0xc992a6f4, 0xa783c823},
		{0xc13cc747, 0xffcc8185, 0x66514e9e, 0xe4ccc199, 0xcd5c766d, 0xa004f676},
		{0x1d3a1fa6, 0xd46894ec, 0xf49c33e6, 0x782fda7e, 0x1fe6346c, 0x0ffe981c},
		{0x78b97c3d, 0x956e8ee8, 0x49ab721c, 0x2672138a, 0x037ea242, 0xce5fe8a4},
		{0x225f7158, 0x32d83e3e, 0xe118f6aa, 0x1fb83751, 0x4d27715c, 0xed2fba4e},
		{0x8dfbc56d, 0xe0a907db, 0xe4af091c, 0x5e123225, 0xd0e8d2e1, 0xcc4501fb},
		{0x8422a8f0, 0x46a12f92, 0x415152ad, 0xf55417f5, 0x38738248, 0xc6e29ded},
		{0x5723715e, 0xabfa788c, 0xc3646af7, 0x64af9186, 0x8fc855ec, 0x2bc36989},
		{0x5e6b28e3, 0xe0f5f592, 0xeb3dd108, 0x0551012a, 0x50e4221d, 0x97e85c0f},
		{0x4e258e14, 0x92298f0b, 0x771269c3, 0x6f934254, 0xc0933b6b, 0x421159b8},
		{0xd76953f4, 0x6a3e36be, 0x53b656fb, 0x610c22e0, 0x9f399330, 0xacf7e7e9},
		{0xfe0b573b, 0xcbb73085, 0x89ed67fc, 0x77014cef, 0xe1b8431f, 0xba1b4105},
		{0x06de3450, 0xb3f5b2fe, 0xdf1cec27, 0xfb22bd10, 0x8e3de6fe, 0x3d4acd27},
		{0xc5444873, 0x5bec968b, 0x8b2af393, 0x11e2f6ca, 0x9cb3694f, 0x94c56b91},
		{0x939a1a93, 0x27f101bb, 0x5381bae7, 0x48ebd1b1, 0xf6d5fca7, 0x0ca24bbc},
		{0x7b03490b, 0xde00acfb, 0xc7f8abfe, 0x410a14c1, 0xd37932a9, 0x14029327},
		{0xbd948525, 0x2c75004d, 0xc52486d5, 0x0f07e2fa, 0x1963e1fd, 0x882719c3},
	}
	ans256 := [][6]word{
		{0x003a0fd4, 0x02497010, 0x194f7db1, 0x090d0883, 0x2ff5805a, 0xc2580b27},
		{0xa83e7ef9, 0x053eca29, 0xd359f988, 0x8101a243, 0x9bbf34b3, 0x9228434f},
		{0x2efee506, 0x8b5f7bd4, 0x9991e811, 0x72dbc20c, 0x2384c97f, 0xcefee47f},
		{0xc571782c, 0x00da90b1, 0xb940a552, 0x5db79619, 0x4bc9a125, 0x5d08a419},
		{0x72de26cc, 0xd69bc26f, 0x46a7f207, 0x66ff4d81, 0xa87862fc, 0xa5f63601},
		{0x7909c4fa, 0xf3f93651, 0x72cb0bcd, 0xae69b2e3, 0x80f2ca4b, 0xf13efcce},
		{0x7869db69, 0x6b7a5b8e, 0xfefbf6b1, 0xec608c8e, 0x76e9d5d2, 0x13ca4bf6},
		{0xc5eeec7a, 0xaa42a59d, 0x1f22cd00, 0xfdd92bdc, 0xd6bbe3e8, 0x15d459ec},
		{0xcda7632a, 0x9cf01bef, 0x6596e261, 0x8c1de14c, 0x1127c3b8, 0x48b3f629},
		{0x3723d0e1, 0xfc0317ec, 0x3fdd5378, 0x0201ae1d, 0xe55db65e, 0xe4c84dbc},
		{0x3633db3f, 0xe4c24fc2, 0xbb1e1fd7, 0xa339425c, 0xfe3e1bdf, 0xd61c808d},
		{0xbdca3449, 0xbeb8aa4e, 0x145a9687, 0xeb6fcd87, 0x8b88ca72, 0x7677a84b},
		{0xd11005e9, 0x558275c5, 0xbc742819, 0x3f17e888, 0x20fcb71f, 0x60886959},
		{0x8d9446c4, 0x67d2d167, 0x855a6aef, 0x69ea517c, 0x36e48e11, 0x0d3f4e86},
		{0xbb0ede65, 0xcceecc06, 0xefc9c49f, 0x44902261, 0xbd8549c0, 0xa7e7f682},
		{0x772101e6, 0xb4b9a250, 0x6faa7b73, 0x7318b792, 0x1e57e751, 0xfd43b41c},
		{0x4ec21b5f, 0xdcfbf30b, 0xa4046947, 0xbe0e781c, 0xd74e21ac, 0x6b1f5d22},
		{0xe8b8e02b, 0x4a662d2d, 0xb50f9ca9, 0x01c98c69, 0x9eb28089, 0x216cfd3f},
		{0x92f0126b, 0x7b9961aa, 0x581f94ac, 0xab4be6dd, 0xc2a91af5, 0xfb4e8e0c},
		{0x4c2c8f04, 0x81a45991, 0x1fcb946c, 0xbccbb5b5, 0x808899cb, 0x8c1b2f89},
		{0x192061be, 0x78e5cf04, 0xf239ab5c, 0xe8471e86, 0x9e6217c7, 0xe5fdf35c},
		{0x83c3150d, 0x766887f8, 0xa1092ac7, 0x6aa6f41d, 0x16e200f9, 0x6bdc26ca},
		{0x52345706, 0xdb70d6af, 0xa8d8ffeb, 0x492ee661, 0x4cd1e991, 0xd75d8352},
		{0x85a9c5fb, 0x1e0f569e, 0x7ff7c600, 0x3f36a1d8, 0xe406ad00, 0x4ded8f16},
		{0x512bb2f4, 0x772b192c, 0x2e6168bd, 0x76af67e1, 0xd893a786, 0x3e276f69},
		{0xd11ee3ad, 0xb7f8c612, 0xd3b19318, 0x89fee4db, 0xb6c3aedd, 0x05420f90},
		{0x04f662f0, 0x8fb41a6c, 0x2f42dd5e, 0xa8ad1839, 0x46474e45, 0x46418de0},
		{0x351550c8, 0x668014f6, 0x04924365, 0x5f353d6f, 0x4eba8d76, 0x924a4318},
		{0x5aba711c, 0xa36b1398, 0x5b3e7bf4, 0x7b3a2cf9, 0x1d006ebe, 0x0d5683e5},
		{0x4f56916f, 0x215dccd2, 0x9f57886f, 0x876d1357, 0x46013d49, 0x2a4932a3},
		{0xaa285691, 0xebefe7d3, 0xe960e64b, 0xdd893f0f, 0x6a234412, 0x495d13c9},
		{0x71c683e8, 0x8069dfd0, 0x6c1a501d, 0x00699418, 0x262142f0, 0xa91a7393},
	}

	K := []byte{0x0f, 0x1e, 0x2d, 0x3c, 0x4b, 0x5a, 0x69, 0x78, 0x87, 0x96, 0xa5, 0xb4, 0xc3, 0xd2, 0xe1, 0xf0}
	RK := RoundKey(K, ENCRYPT_MODE)
	if len(RK) != 24 {
		t.Errorf("Derived round keys from a 128-bit key should have 24 rows (you have %v rows).", len(RK))
	}
	for i := 0; i < 24; i++ {
		for j := 0; j < 6; j++ {
			if RK[i][j] != ans128[i][j] {
				t.Errorf("Round key #%v for encryption is invalid. Answer: %v, Your answer %v.", i, ans128[i], RK[i])
				break
			}
		}
	}

	K = []byte{0x0f, 0x1e, 0x2d, 0x3c, 0x4b, 0x5a, 0x69, 0x78, 0x87, 0x96, 0xa5, 0xb4, 0xc3, 0xd2, 0xe1, 0xf0, 0xf0, 0xe1, 0xd2, 0xc3, 0xb4, 0xa5, 0x96, 0x87}
	RK = RoundKey(K, ENCRYPT_MODE)
	if len(RK) != 28 {
		t.Errorf("Derived round keys from a 192-bit key should have 28 rows (you have %v rows).", len(RK))
	}
	for i := 0; i < 28; i++ {
		for j := 0; j < 6; j++ {
			if RK[i][j] != ans192[i][j] {
				t.Errorf("Round key #%v for encryption is invalid. Answer: %v, Your answer %v.", i, ans192[i], RK[i])
				break
			}
		}
	}

	K = []byte{0x0f, 0x1e, 0x2d, 0x3c, 0x4b, 0x5a, 0x69, 0x78, 0x87, 0x96, 0xa5, 0xb4, 0xc3, 0xd2, 0xe1, 0xf0, 0xf0, 0xe1, 0xd2, 0xc3, 0xb4, 0xa5, 0x96, 0x87, 0x78, 0x69, 0x5a, 0x4b, 0x3c, 0x2d, 0x1e, 0x0f}
	RK = RoundKey(K, ENCRYPT_MODE)
	if len(RK) != 32 {
		t.Errorf("Derived round keys from a 256-bit key should have 32 rows (you have %v rows).", len(RK))
	}
	for i := 0; i < 32; i++ {
		for j := 0; j < 6; j++ {
			if RK[i][j] != ans256[i][j] {
				t.Errorf("Round key #%v for encryption is invalid. Answer: %v, Your answer %v.", i, ans256[i], RK[i])
				break
			}
		}
	}
}

func TestEncRound128Key(t *testing.T) {
	P := []byte{0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f}
	var X [4]word
	for i := 0; i < 4; i++ {
		var tmp [4]byte
		copy(tmp[:], P[i*4:(i+1)*4])
		X[i] = ba2w(tmp)
	}
	Xans := [24][4]word{
		{0x13121110, 0x17161514, 0x1b1a1918, 0x1f1e1d1c},
		{0x0f079051, 0x693d668d, 0xe5edcfd4, 0x13121110},
		{0x3fc44a2d, 0xf767ea2a, 0xa0b67cf0, 0x0f079051},
		{0x99e912cd, 0x906fd05d, 0x4d293e55, 0x3fc44a2d},
		{0x43048c71, 0x5faa8d15, 0xdfc687fb, 0x99e912cd},
		{0x862a337d, 0x419f623d, 0x4b97dd8a, 0x43048c71},
		{0x055b3a34, 0xa2eb0f67, 0xaf9873ba, 0x862a337d},
		{0x38875cb8, 0x19f1c052, 0x02e13d1c, 0x055b3a34},
		{0xf1ddbcca, 0x2c031302, 0x8a5f86d6, 0x38875cb8},
		{0xf770a17e, 0xc47d9365, 0x2df8cda7, 0xf1ddbcca},
		{0x58a898f4, 0xdb57aa1e, 0x20d820a4, 0xf770a17e},
		{0x11c9f487, 0xbf079d6e, 0x28c10b82, 0x58a898f4},
		{0xa7e4a0e4, 0xe8e97f62, 0x47727e5f, 0x11c9f487},
		{0xd1ea924a, 0x2298587f, 0xf2afc1d0, 0xa7e4a0e4},
		{0x7e91cf8c, 0xfcca259f, 0x86ab69cf, 0xd1ea924a},
		{0x809fd3e9, 0xef492067, 0x536df05e, 0x7e91cf8c},
		{0x2b54eee2, 0x98b175f9, 0xd9c14ac4, 0x809fd3e9},
		{0x63eb48a2, 0x7ad2716d, 0x783a355e, 0x2b54eee2},
		{0x4b34e264, 0x101d5f00, 0x7fee2017, 0x63eb48a2},
		{0xba42cf9e, 0xd156295c, 0xb88c1f9d, 0x4b34e264},
		{0x970433ea, 0xa0d420cb, 0x4b96b2c1, 0xba42cf9e},
		{0x49facf18, 0x6f1fe3c2, 0x3744e7b8, 0x970433ea},
		{0xd1527e90, 0x6ce66afe, 0x1d55c7f1, 0x49facf18},
		{0xfd8b6404, 0x8675df3b, 0xe4b9d73f, 0xd1527e90},
	}
	K := []byte{0x0f, 0x1e, 0x2d, 0x3c, 0x4b, 0x5a, 0x69, 0x78, 0x87, 0x96, 0xa5, 0xb4, 0xc3, 0xd2, 0xe1, 0xf0}
	RK := RoundKey(K, ENCRYPT_MODE)
	for i := 0; i < 24; i++ {
		for j := 0; j < 4; j++ {
			if X[j] != Xans[i][j] {
				t.Errorf("Invalid result: round #%v for encryption. Answer: %v, yours: %v", i, Xans[i], X)
				return
			}
		}
		X = EncRound(X, RK[i])
	}
}

func TestDecRound128Key(t *testing.T) {
	C := [16]byte{0x9f, 0xc8, 0x4e, 0x35, 0x28, 0xc6, 0xc6, 0x18, 0x55, 0x32, 0xc7, 0xa7, 0x04, 0x64, 0x8b, 0xfd}
	var X [4]word
	for i := 0; i < 4; i++ {
		var buf [4]byte
		copy(buf[:], C[i*4:(i+1)*4])
		X[i] = ba2w(buf)
	}
	Xans := [24][4]word{
		{0x354ec89f, 0x18c6c628, 0xa7c73255, 0xfd8b6404},
		{0xfd8b6404, 0x8675df3b, 0xe4b9d73f, 0xd1527e90},
		{0xd1527e90, 0x6ce66afe, 0x1d55c7f1, 0x49facf18},
		{0x49facf18, 0x6f1fe3c2, 0x3744e7b8, 0x970433ea},
		{0x970433ea, 0xa0d420cb, 0x4b96b2c1, 0xba42cf9e},
		{0xba42cf9e, 0xd156295c, 0xb88c1f9d, 0x4b34e264},
		{0x4b34e264, 0x101d5f00, 0x7fee2017, 0x63eb48a2},
		{0x63eb48a2, 0x7ad2716d, 0x783a355e, 0x2b54eee2},
		{0x2b54eee2, 0x98b175f9, 0xd9c14ac4, 0x809fd3e9},
		{0x809fd3e9, 0xef492067, 0x536df05e, 0x7e91cf8c},
		{0x7e91cf8c, 0xfcca259f, 0x86ab69cf, 0xd1ea924a},
		{0xd1ea924a, 0x2298587f, 0xf2afc1d0, 0xa7e4a0e4},
		{0xa7e4a0e4, 0xe8e97f62, 0x47727e5f, 0x11c9f487},
		{0x11c9f487, 0xbf079d6e, 0x28c10b82, 0x58a898f4},
		{0x58a898f4, 0xdb57aa1e, 0x20d820a4, 0xf770a17e},
		{0xf770a17e, 0xc47d9365, 0x2df8cda7, 0xf1ddbcca},
		{0xf1ddbcca, 0x2c031302, 0x8a5f86d6, 0x38875cb8},
		{0x38875cb8, 0x19f1c052, 0x02e13d1c, 0x055b3a34},
		{0x055b3a34, 0xa2eb0f67, 0xaf9873ba, 0x862a337d},
		{0x862a337d, 0x419f623d, 0x4b97dd8a, 0x43048c71},
		{0x43048c71, 0x5faa8d15, 0xdfc687fb, 0x99e912cd},
		{0x99e912cd, 0x906fd05d, 0x4d293e55, 0x3fc44a2d},
		{0x3fc44a2d, 0xf767ea2a, 0xa0b67cf0, 0x0f079051},
		{0x0f079051, 0x693d668d, 0xe5edcfd4, 0x13121110},
	}
	K := []byte{0x0f, 0x1e, 0x2d, 0x3c, 0x4b, 0x5a, 0x69, 0x78, 0x87, 0x96, 0xa5, 0xb4, 0xc3, 0xd2, 0xe1, 0xf0}
	RK := RoundKey(K, DECRYPT_MODE)
	for i := 0; i < 24; i++ {
		for j := 0; j < 4; j++ {
			if X[j] != Xans[i][j] {
				t.Errorf("Round result #%v for decryption is invalid. Answer: %v, your answer %v.", i, Xans[i], X)
				return
			}
		}
		X = DecRound(X, RK[i])
	}
}

func TestEncRound192Key(t *testing.T) {
	P := []byte{0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f}
	var X [4]word
	for i := 0; i < 4; i++ {
		var tmp [4]byte
		copy(tmp[:], P[i*4:(i+1)*4])
		X[i] = ba2w(tmp)
	}
	Xans := [28][4]word{
		{0x23222120, 0x27262524, 0x2b2a2928, 0x2f2e2d2c},
		{0x0f085091, 0x030483d2, 0xbe4ab9ef, 0x23222120},
		{0x23fc7579, 0x99fa0bb5, 0x92d65065, 0x0f085091},
		{0x1e64758b, 0x6b19e366, 0x3edbb998, 0x23fc7579},
		{0x8da45638, 0xd886dfdd, 0x2da2b83c, 0x1e64758b},
		{0xa29dfd15, 0xd8687adf, 0xb89c47b4, 0x8da45638},
		{0x34e43c2e, 0xb6268ba7, 0x584086d7, 0xa29dfd15},
		{0xdf6e285b, 0x88f26855, 0x198fbb0c, 0x34e43c2e},
		{0xe62dde25, 0xdd1cdf46, 0xb8049544, 0xdf6e285b},
		{0xd715e465, 0x0e4d136e, 0x35bc93a5, 0xe62dde25},
		{0x1ab97de4, 0xa5c19c64, 0xcfd627b0, 0xd715e465},
		{0x1a155930, 0x4ccf6ee2, 0x8c5136f7, 0x1ab97de4},
		{0x0eef4d0d, 0x9f3065e1, 0x405fc8b9, 0x1a155930},
		{0xa0dd5c61, 0xfcefa1a4, 0x48e2adc3, 0x0eef4d0d},
		{0xdcf21fcc, 0xf9ca084f, 0x0b02cdd8, 0xa0dd5c61},
		{0xdfd53021, 0x2eee92c5, 0xeedfe48b, 0xdcf21fcc},
		{0x81dce833, 0x4e0af1c2, 0x3abac76b, 0xdfd53021},
		{0x9646ef75, 0x607a7771, 0x9fbc48ec, 0x81dce833},
		{0x7f40d072, 0xac609c27, 0x5de1c810, 0x9646ef75},
		{0xfd0bae5f, 0x35429a83, 0x11f5e49f, 0x7f40d072},
		{0x2feb9af2, 0x0799218a, 0xe5374a5f, 0xfd0bae5f},
		{0xfd86cfee, 0xa7d97a82, 0x7c97ed23, 0x2feb9af2},
		{0xadd0adf1, 0xe09057e1, 0xccd95f65, 0xfd86cfee},
		{0x06c45cfe, 0x392aaa1d, 0xae9fd56c, 0xadd0adf1},
		{0xf3032315, 0xb1df9d75, 0x1627928d, 0x06c45cfe},
		{0xf4eec840, 0x6a15d699, 0x2392c666, 0xf3032315},
		{0xb353eb6a, 0xad286c22, 0x5a9d146d, 0xf4eec840},
		{0xf2c67476, 0x44333e44, 0x6d5a1045, 0xb353eb6a},
	}
	K := []byte{0x0f, 0x1e, 0x2d, 0x3c, 0x4b, 0x5a, 0x69, 0x78, 0x87, 0x96, 0xa5, 0xb4, 0xc3, 0xd2, 0xe1, 0xf0, 0xf0, 0xe1, 0xd2, 0xc3, 0xb4, 0xa5, 0x96, 0x87}
	RK := RoundKey(K, ENCRYPT_MODE)
	for i := 0; i < 28; i++ {
		for j := 0; j < 4; j++ {
			if X[j] != Xans[i][j] {
				t.Errorf("Round result #%v for encryption is invalid. Answer: %v, your answer %v.", i, Xans[i], X)
				break
			}
		}
		X = EncRound(X, RK[i])
	}
}

func TestEncRound256Key(t *testing.T) {
	P := []byte{0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f}
	var X [4]word
	for i := 0; i < 4; i++ {
		var tmp [4]byte
		copy(tmp[:], P[i*4:(i+1)*4])
		X[i] = ba2w(tmp)
	}
	Xans := [32][4]word{
		{0x33323130, 0x37363534, 0x3b3a3938, 0x3f3e3d3c},
		{0x0f0810d1, 0x030583d2, 0xa246bdef, 0x33323130},
		{0xe370475a, 0x379d1cd0, 0x7b627f7b, 0x0f0810d1},
		{0xa212c114, 0xc5be3591, 0x435bb556, 0xe370475a},
		{0x90bcb059, 0x94df55a0, 0xd8e15ef6, 0xa212c114},
		{0x4e5cc849, 0xf484b5d8, 0xef0fc663, 0x90bcb059},
		{0xa520787d, 0xae3db194, 0xfa2feb17, 0x4e5cc849},
		{0x231a5d45, 0xf338ad75, 0x9d4b9850, 0xa520787d},
		{0xdd744e80, 0x0a6568a0, 0x3f9c93a9, 0x231a5d45},
		{0xd141f34e, 0x311ba7ed, 0xb34c9f6f, 0xdd744e80},
		{0xf5a76166, 0x3e00a130, 0xb1f9a58d, 0xd141f34e},
		{0xaf52973c, 0xc4befd35, 0xaae4a642, 0xf5a76166},
		{0x3df5e119, 0xb8937ebb, 0xb4a7a6ab, 0xaf52973c},
		{0xede0ddb3, 0x2c84bd26, 0x2c86c203, 0x3df5e119},
		{0x960f7157, 0x477a5b5a, 0x29659f76, 0xede0ddb3},
		{0x2c8d1d71, 0xe0b54ae6, 0xfbdd003c, 0x960f7157},
		{0x720a9b5f, 0x18bf274a, 0x0a1af597, 0x2c8d1d71},
		{0x1aa88202, 0xc3867edc, 0xc49ce291, 0x720a9b5f},
		{0xe16c34f7, 0x69defa8b, 0x15b2990f, 0x1aa88202},
		{0xc7837b0b, 0xcf85d76f, 0x17203201, 0xe16c34f7},
		{0xa3061bb3, 0xbbe1ce55, 0x00a3f8e9, 0xc7837b0b},
		{0x54f6bcfa, 0xc195ea5b, 0xb8280ef0, 0xa3061bb3},
		{0x662f351e, 0x49995ddc, 0x4ef48970, 0x54f6bcfa},
		{0x09db178e, 0x4748e08a, 0x30ba1411, 0x662f351e},
		{0x751113cb, 0x9a425ee2, 0x200fee63, 0x09db178e},
		{0x47d21a23, 0x08561dff, 0x86131859, 0x751113cb},
		{0xf7aaf6ac, 0x4f5eac5b, 0xf4247a5b, 0x47d21a23},
		{0x8e952768, 0x3de52e9b, 0x367ed97c, 0xf7aaf6ac},
		{0xcb641a2d, 0x8d161a90, 0xdbd4a137, 0x8e952768},
		{0xb6e87380, 0x93b8b779, 0xc9530e82, 0xcb641a2d},
		{0x48bd3559, 0x5ad96ae7, 0x2e0feb8b, 0xb6e87380},
		{0x97e1f927, 0x853a0309, 0x487c41fc, 0x48bd3559},
	}
	K := []byte{0x0f, 0x1e, 0x2d, 0x3c, 0x4b, 0x5a, 0x69, 0x78, 0x87, 0x96, 0xa5, 0xb4, 0xc3, 0xd2, 0xe1, 0xf0, 0xf0, 0xe1, 0xd2, 0xc3, 0xb4, 0xa5, 0x96, 0x87, 0x78, 0x69, 0x5a, 0x4b, 0x3c, 0x2d, 0x1e, 0x0f}
	RK := RoundKey(K, ENCRYPT_MODE)
	for i := 0; i < 32; i++ {
		for j := 0; j < 4; j++ {
			if X[j] != Xans[i][j] {
				t.Errorf("Round result #%v for encryption is invalid. Answer: %v, your answer %v.", i, Xans[i], X)
				break
			}
		}
		X = EncRound(X, RK[i])
	}
}

func TestEncDec128(t *testing.T) {
	K := []byte{0x0f, 0x1e, 0x2d, 0x3c, 0x4b, 0x5a, 0x69, 0x78, 0x87, 0x96, 0xa5, 0xb4, 0xc3, 0xd2, 0xe1, 0xf0}
	P := [16]byte{0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f}

	CANS := [16]byte{0x9f, 0xc8, 0x4e, 0x35, 0x28, 0xc6, 0xc6, 0x18, 0x55, 0x32, 0xc7, 0xa7, 0x04, 0x64, 0x8b, 0xfd}
	RKE := RoundKey(K, ENCRYPT_MODE)
	C := Encrypt(P, RKE)
	for i := 0; i < 16; i++ {
		if C[i] != CANS[i] {
			t.Errorf("Invalid encryption. Answer: %v, yours %v", CANS, C)
			return
		}
	}

	RKD := RoundKey(K, DECRYPT_MODE)
	DP := Decrypt(C, RKD)
	for i := 0; i < 16; i++ {
		if DP[i] != P[i] {
			t.Errorf("Decryption was invalid. Answer: %v, yours: %v", P, DP)
			return
		}
	}
}

func TestEncDec192(t *testing.T) {
	K := []byte{0x0f, 0x1e, 0x2d, 0x3c, 0x4b, 0x5a, 0x69, 0x78, 0x87, 0x96, 0xa5, 0xb4, 0xc3, 0xd2, 0xe1, 0xf0, 0xf0, 0xe1, 0xd2, 0xc3, 0xb4, 0xa5, 0x96, 0x87}
	P := [16]byte{0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f}
	CANS := []byte{0x6f, 0xb9, 0x5e, 0x32, 0x5a, 0xad, 0x1b, 0x87, 0x8c, 0xdc, 0xf5, 0x35, 0x76, 0x74, 0xc6, 0xf2}
	RKE := RoundKey(K, ENCRYPT_MODE)
	C := Encrypt(P, RKE)
	for i := 0; i < 16; i++ {
		if C[i] != CANS[i] {
			t.Errorf("Invalid encryption. Answer: %v, yours %v", CANS, C)
			return
		}
	}
	RKD := RoundKey(K, DECRYPT_MODE)
	DP := Decrypt(C, RKD)
	for i := 0; i < 16; i++ {
		if DP[i] != P[i] {
			t.Errorf("Decryption was invalid. Answer: %v, yours: %v", P, DP)
			return
		}
	}
}

func TestEncDec256(t *testing.T) {
	K := []byte{0x0f, 0x1e, 0x2d, 0x3c, 0x4b, 0x5a, 0x69, 0x78, 0x87, 0x96, 0xa5, 0xb4, 0xc3, 0xd2, 0xe1, 0xf0, 0xf0, 0xe1, 0xd2, 0xc3, 0xb4, 0xa5, 0x96, 0x87, 0x78, 0x69, 0x5a, 0x4b, 0x3c, 0x2d, 0x1e, 0x0f}
	P := [16]byte{0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f}
	CANS := []byte{0xd6, 0x51, 0xaf, 0xf6, 0x47, 0xb1, 0x89, 0xc1, 0x3a, 0x89, 0x00, 0xca, 0x27, 0xf9, 0xe1, 0x97}
	RKE := RoundKey(K, ENCRYPT_MODE)
	C := Encrypt(P, RKE)
	for i := 0; i < 16; i++ {
		if C[i] != CANS[i] {
			t.Errorf("Invalid encryption. Answer: %v, yours %v", CANS, C)
			return
		}
	}
	RKD := RoundKey(K, DECRYPT_MODE)
	DP := Decrypt(C, RKD)
	for i := 0; i < 16; i++ {
		if DP[i] != P[i] {
			t.Errorf("Decryption was invalid. Answer: %v, yours: %v", P, DP)
			return
		}
	}
}
