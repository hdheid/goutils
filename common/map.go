package common

type Region int

// 省级行政区划编码
const (
	BeiJing     Region = 110000 // 北京
	TianJin     Region = 120000 // 天津
	HeBei       Region = 130000 // 河北
	ShanXi      Region = 140000 // 山西
	NeiMengGu   Region = 150000 // 内蒙古
	LiaoNing    Region = 210000 // 辽宁
	JiLin       Region = 220000 // 吉林
	HeiLongJing Region = 230000 // 黑龙江
	ShangHai    Region = 310000 // 上海
	JiangSu     Region = 320000 // 江苏
	ZheJiang    Region = 330000 // 浙江
	AnHui       Region = 340000 // 安徽
	FuJian      Region = 350000 // 福建
	JiangXi     Region = 360000 // 江西
	ShanDong    Region = 370000 // 山东
	HeNan       Region = 410000 // 河南
	HuBei       Region = 420000 // 湖北
	HuNan       Region = 430000 // 湖南
	GuangDong   Region = 440000 // 广东
	GuangXi     Region = 450000 // 广西
	HaiNan      Region = 460000 // 海南
	SiChuan     Region = 510000 // 四川
	GuiZhou     Region = 520000 // 贵州
	YunNan      Region = 530000 // 云南
	XiZang      Region = 540000 // 西藏
	ChongQing   Region = 500000 // 重庆
	ShanXi2     Region = 610000 // 陕西
	GanSu       Region = 620000 // 甘肃
	QingHai     Region = 630000 // 青海
	NingXia     Region = 640000 // 宁夏
	XinJiang    Region = 650000 // 新疆
	TaiWan      Region = 710000 // 台湾
	HongKong    Region = 810000 // 香港
	Macao       Region = 820000 // 澳门
)

func GetProvinceMap() map[Region]string {
	return map[Region]string{
		BeiJing:     "北京",
		TianJin:     "天津",
		HeBei:       "河北",
		ShanXi:      "山西",
		NeiMengGu:   "内蒙古",
		LiaoNing:    "辽宁",
		JiLin:       "吉林",
		HeiLongJing: "黑龙江",
		ShangHai:    "上海",
		JiangSu:     "江苏",
		ZheJiang:    "浙江",
		AnHui:       "安徽",
		FuJian:      "福建",
		JiangXi:     "江西",
		ShanDong:    "山东",
		HeNan:       "河南",
		HuBei:       "湖北",
		HuNan:       "湖南",
		GuangDong:   "广东",
		GuangXi:     "广西",
		HaiNan:      "海南",
		SiChuan:     "四川",
		GuiZhou:     "贵州",
		YunNan:      "云南",
		XiZang:      "西藏",
		ChongQing:   "重庆",
		ShanXi2:     "陕西",
		GanSu:       "甘肃",
		QingHai:     "青海",
		NingXia:     "宁夏",
		XinJiang:    "新疆",
		TaiWan:      "台湾",
		HongKong:    "香港",
		Macao:       "澳门",
	}
}

func GetProvinceKey() map[Region]string {
	return map[Region]string{
		BeiJing:     "BeiJing",
		TianJin:     "TianJin",
		HeBei:       "HeBei",
		ShanXi:      "ShanXi",
		NeiMengGu:   "NeiMengGu",
		LiaoNing:    "LiaoNing",
		JiLin:       "JiLin",
		HeiLongJing: "HeiLongJing",
		ShangHai:    "ShangHai",
		JiangSu:     "JiangSu",
		ZheJiang:    "ZheJiang",
		AnHui:       "AnHui",
		FuJian:      "FuJian",
		JiangXi:     "JiangXi",
		ShanDong:    "ShanDong",
		HeNan:       "HeNan",
		HuBei:       "HuBei",
		HuNan:       "HuNan",
		GuangDong:   "GuangDong",
		GuangXi:     "GuangXi",
		HaiNan:      "HaiNan",
		SiChuan:     "SiChuan",
		GuiZhou:     "GuiZhou",
		YunNan:      "YunNan",
		XiZang:      "XiZang",
		ChongQing:   "ChongQing",
		ShanXi2:     "ShanXi2",
		GanSu:       "GanSu",
		QingHai:     "QingHai",
		NingXia:     "NingXia",
		XinJiang:    "XinJiang",
		TaiWan:      "TaiWan",

		HongKong: "HongKong",
		Macao:    "Macao:",
	}
}

// nolint:all
const ( // 特别行政区
	C_HK string = "中国香港特别行政区"
	C_TW string = "中国台湾省"
	C_MC string = "中国澳门特别行政区"
)
