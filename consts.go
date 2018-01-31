package demofile

const (
	MaxEditctBits                       = 11
	NumNetworkedEhandleSerialNumberBits = 10
	NumNetworkedEhandleBits             = MaxEditctBits + NumNetworkedEhandleSerialNumberBits
	InvalidNetworkedEhandleValue        = (1 << NumNetworkedEhandleBits) - 1
	IndexMask                           = ((1 << MaxEditctBits) - 1)
)

const weaponPrefix = "weapon_"

type (
	RoundMVPReason   byte
	Hitgroup         byte
	RoundEndReason   byte
	Team             byte
	EquipmentElement int
	EquipmentClass   int
)

const (
	MVPReason_MostEliminations RoundMVPReason = iota + 1
	MVPReason_BombDefused
	MVPReason_BombPlanted
)

const (
	HG_Generic  Hitgroup = 0
	HG_Head     Hitgroup = 1
	HG_Chest    Hitgroup = 2
	HG_Stomach  Hitgroup = 3
	HG_LeftArm  Hitgroup = 4
	HG_RightArm Hitgroup = 5
	HG_LeftLeg  Hitgroup = 6
	HG_RightLeg Hitgroup = 7
	HG_Gear     Hitgroup = 10
)

const (
	RER_TargetBombed RoundEndReason = iota + 1
	RER_VIPEscaped
	RER_VIPKilled
	RER_TerroristsEscaped
	RER_CTStoppedEscape
	RER_TerroristsStopped
	RER_BombDefused
	RER_CTWin
	RER_TerroristsWin
	RER_Draw
	RER_HostagesRescued
	RER_TargetSaved
	RER_HostagesNotRescued
	RER_TerroristsNotEscaped
	RER_VIPNotEscaped
	RER_GameStart
	RER_TerroristsSurrender
	RER_CTSurrender
)

const (
	Team_Unassigned Team = iota
	Team_Spectators
	Team_Terrorists
	Team_CounterTerrorists
)

const (
	EE_Unknown EquipmentElement = 0

	// Pistols

	EE_P2000        EquipmentElement = 1
	EE_Glock        EquipmentElement = 2
	EE_P250         EquipmentElement = 3
	EE_Deagle       EquipmentElement = 4
	EE_FiveSeven    EquipmentElement = 5
	EE_DualBarettas EquipmentElement = 6
	EE_Tec9         EquipmentElement = 7
	EE_CZ           EquipmentElement = 8
	EE_USP          EquipmentElement = 9
	EE_Revolver     EquipmentElement = 10

	// SMGs

	EE_MP7   EquipmentElement = 101
	EE_MP9   EquipmentElement = 102
	EE_Bizon EquipmentElement = 103
	EE_Mac10 EquipmentElement = 104
	EE_UMP   EquipmentElement = 105
	EE_P90   EquipmentElement = 106

	// Heavy

	EE_SawedOff EquipmentElement = 201
	EE_Nova     EquipmentElement = 202
	EE_Swag7    EquipmentElement = 203
	EE_XM1014   EquipmentElement = 204
	EE_M249     EquipmentElement = 205
	EE_Negev    EquipmentElement = 206

	// Rifles

	EE_Gallil EquipmentElement = 301
	EE_Famas  EquipmentElement = 302
	EE_AK47   EquipmentElement = 303
	EE_M4A4   EquipmentElement = 304
	EE_M4A1   EquipmentElement = 305
	EE_Scout  EquipmentElement = 306
	EE_SG556  EquipmentElement = 307
	EE_AUG    EquipmentElement = 308
	EE_AWP    EquipmentElement = 309
	EE_Scar20 EquipmentElement = 310
	EE_G3SG1  EquipmentElement = 311

	// Equipment

	EE_Zeus      EquipmentElement = 401
	EE_Kevlar    EquipmentElement = 402
	EE_Helmet    EquipmentElement = 403
	EE_Bomb      EquipmentElement = 404
	EE_Knife     EquipmentElement = 405
	EE_DefuseKit EquipmentElement = 406
	EE_World     EquipmentElement = 407

	// Grenades

	EE_Decoy      EquipmentElement = 501
	EE_Molotov    EquipmentElement = 502
	EE_Incendiary EquipmentElement = 503
	EE_Flash      EquipmentElement = 504
	EE_Smoke      EquipmentElement = 505
	EE_HE         EquipmentElement = 506
)

const (
	EC_Unknown EquipmentClass = iota
	EC_Pistols
	EC_SMG
	EC_Heavy
	EC_Rifle
	EC_Equipment
	EC_Grenade
)

var (
	DefaultAmmoInWeapon = map[EquipmentElement]int{
		EE_Unknown:      -1,
		EE_P2000:        13,
		EE_Glock:        20,
		EE_P250:         13,
		EE_Deagle:       7,
		EE_FiveSeven:    20,
		EE_DualBarettas: 30,
		EE_Tec9:         18,
		EE_CZ:           12,
		EE_USP:          12,
		EE_Revolver:     8,
		EE_MP7:          30,
		EE_MP9:          30,
		EE_Bizon:        64,
		EE_Mac10:        30,
		EE_UMP:          25,
		EE_P90:          50,
		EE_SawedOff:     7,
		EE_Nova:         8,
		EE_Swag7:        5,
		EE_XM1014:       7,
		EE_M249:         100,
		EE_Negev:        150,
		EE_Gallil:       35,
		EE_Famas:        25,
		EE_AK47:         30,
		EE_M4A4:         30,
		EE_M4A1:         20,
		EE_Scout:        10,
		EE_SG556:        30,
		EE_AUG:          30,
		EE_AWP:          10,
		EE_Scar20:       20,
		EE_G3SG1:        20,
		EE_Zeus:         1,
		EE_Kevlar:       0,
		EE_Helmet:       0,
		EE_Bomb:         1,
		EE_Knife:        0,
		EE_DefuseKit:    1,
		EE_World:        0,
		EE_Decoy:        1,
		EE_Molotov:      1,
		EE_Incendiary:   1,
		EE_Flash:        1,
		EE_Smoke:        1,
		EE_HE:           1,
	}

	MoveSpeedsInWeapon = map[EquipmentElement]int{
		EE_Unknown:      500,
		EE_P2000:        240,
		EE_Glock:        240,
		EE_P250:         240,
		EE_Deagle:       230,
		EE_FiveSeven:    240,
		EE_DualBarettas: 240,
		EE_Tec9:         240,
		EE_CZ:           240,
		EE_USP:          240,
		EE_Revolver:     220,
		EE_MP7:          220,
		EE_MP9:          240,
		EE_Bizon:        240,
		EE_Mac10:        240,
		EE_UMP:          230,
		EE_P90:          230,
		EE_SawedOff:     210,
		EE_Nova:         220,
		EE_Swag7:        225,
		EE_XM1014:       215,
		EE_M249:         195,
		EE_Negev:        195,
		EE_Gallil:       215,
		EE_Famas:        220,
		EE_AK47:         215,
		EE_M4A4:         225,
		EE_M4A1:         225,
		EE_Scout:        230,
		EE_SG556:        210,
		EE_AUG:          220,
		EE_AWP:          200,
		EE_Scar20:       215,
		EE_G3SG1:        215,
		EE_Zeus:         240,
		EE_Bomb:         250,
		EE_Knife:        250,
		EE_Decoy:        245,
		EE_Molotov:      245,
		EE_Incendiary:   245,
		EE_Flash:        245,
		EE_Smoke:        245,
		EE_HE:           245,
	}

	PricesOfWeapons = map[EquipmentElement]int{
		EE_Unknown:      0,
		EE_P2000:        200,
		EE_Glock:        200,
		EE_P250:         250,
		EE_Deagle:       700,
		EE_FiveSeven:    500,
		EE_DualBarettas: 500,
		EE_Tec9:         500,
		EE_CZ:           500,
		EE_USP:          200,
		EE_Revolver:     700,
		EE_MP7:          1700,
		EE_MP9:          1250,
		EE_Bizon:        1400,
		EE_Mac10:        1050,
		EE_UMP:          1200,
		EE_P90:          2350,
		EE_SawedOff:     1200,
		EE_Nova:         1200,
		EE_Swag7:        1800,
		EE_XM1014:       2000,
		EE_M249:         5200,
		EE_Negev:        2000,
		EE_Gallil:       2000,
		EE_Famas:        2250,
		EE_AK47:         2700,
		EE_M4A4:         3100,
		EE_M4A1:         3100,
		EE_Scout:        1700,
		EE_SG556:        3000,
		EE_AUG:          3300,
		EE_AWP:          4750,
		EE_Scar20:       5000,
		EE_G3SG1:        5000,
		EE_Zeus:         200,
		EE_Decoy:        50,
		EE_Molotov:      400,
		EE_Incendiary:   600,
		EE_Flash:        200,
		EE_Smoke:        300,
		EE_HE:           300,
	}
)
