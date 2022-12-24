package service

const (
	AN       = "AN"       //申请号
	AD       = "AD"       //申请日
	ADY      = "ADY"      //申请年
	PNM      = "PNM"      //公开(公告)号
	PD       = "PD"       //公开(公告)日
	PDY      = "PDY"      //公开年
	KC       = "KC"       //公开类型
	PR       = "PR"       //优先权
	EPRD     = "EPRD"     //最早优先权日
	DAN      = "DAN"      //分案原申请号
	TI       = "TI"       //名称
	TICN     = "TICN"     //名称中
	TIEN     = "TIEN"     //名称英
	TIJP     = "TIJP"     //名称日
	TIOL     = "TIOL"     //名称OL
	ABST     = "ABST"     //摘要
	ABSTCN   = "ABSTCN"   //摘要中
	ABSTEN   = "ABSTEN"   //摘要英
	ABSTJP   = "ABSTJP"   //摘要日
	ABSTOL   = "ABSTOL"   //摘要OL
	CLM      = "CLM"      //权利要求
	CL       = "CL"       //主权项
	CLZS     = "CLZS"     //主权项字数
	CLMN     = "CLMN"     //权项数
	ICLMN    = "ICLMN"    //独权数
	DESCR    = "DESCR"    //说明书
	DPN      = "DPN"      //说明书页数
	TA       = "TA"       //名称摘要
	TAC      = "TAC"      //名称摘要权利要求书
	TACD     = "TACD"     //名称摘要权利要求书说明书
	AC       = "AC"       //摘要权利要求书
	ACD      = "ACD"      //摘要权利要求书说明书
	PA       = "PA"       //申请(专利权)人
	SPATMS   = "SPATMS"   //标准申请人
	PPA      = "PPA"      //主申请人
	PATMS    = "PATMS"    //申请人集合
	CAS      = "CAS"      //当前专利权人
	SCASTMS  = "SCASTMS"  //标准当前专利权人
	PCAS     = "PCAS"     //当前主专利权人
	CASTMS   = "CASTMS"   //当前专利权人集合
	PAS      = "PAS"      //预测专利权人
	AR       = "AR"       //申请人地址
	CAR      = "CAR"      //当前专利权人地址
	ADDRP    = "ADDRP"    //申请人（省市）
	ADDRC    = "ADDRC"    //申请人（地市）
	ADDRDC   = "ADDRDC"   //申请人（区县）
	CASP     = "CASP"     //当前专利权人（省市）
	CASC     = "CASC"     //当前专利权人（地市）
	CASD     = "CASD"     //当前专利权人（区县）
	PAN      = "PAN"      //申请人数
	PACC     = "PACC"     //申请人国家/地域
	PPACC    = "PPACC"    //主申请人国家/地域
	INN      = "INN"      //发明（设计）人
	PINN     = "PINN"     //主发明人
	INNTMS   = "INNTMS"   //发明人集合
	INNN     = "INNN"     //发明人数
	SIC      = "SIC"      //分类号
	PIC      = "PIC"      //主分类号
	SICN     = "SICN"     //分类数
	IPC      = "IPC"      //IPC分类号
	PIPC     = "PIPC"     //IPC主分类号
	IPCS     = "IPCS"     //IPC分类部
	IPCC     = "IPCC"     //IPC大类
	IPCSC    = "IPCSC"    //IPC小类
	IPCSCN   = "IPCSCN"   //IPC小类数
	ICG      = "ICG"      //IPC大组
	ICSG     = "ICSG"     //IPC小组
	PICS     = "PICS"     //主IPC部
	PICC     = "PICC"     //主IPC大类
	PICSC    = "PICSC"    //主IPC小类
	PICG     = "PICG"     //主IPC大组
	PICSG    = "PICSG"    //主IPC小组
	IPCN     = "IPCN"     //IPC分类数
	IPCBN    = "IPCBN"    //IPC部数
	LOC      = "LOC"      //LOC分类号
	PLOC     = "PLOC"     //LOC主分类号
	LOCN     = "LOCN"     //LOC分类数
	SEC      = "SEC"      //欧洲分类号
	PEC      = "PEC"      //欧洲主分类号
	ECN      = "ECN"      //欧洲分类数
	CPC      = "CPC"      //CPC分类号
	PCPC     = "PCPC"     //CPC主分类号
	CPCN     = "CPCN"     //CPC分类数
	CPCS     = "CPCS"     //CPC部
	CPCC     = "CPCC"     //CPC大类
	CPCSC    = "CPCSC"    //CPC小类
	CPCG     = "CPCG"     //CPC大组
	CPCSG    = "CPCSG"    //CPC小组
	PCPCS    = "PCPCS"    //主CPC部
	PCPCC    = "PCPCC"    //主CPC大类
	PCPCSC   = "PCPCSC"   //主CPC小类
	PCPCG    = "PCPCG"    //主CPC大组
	PCPCSG   = "PCPCSG"   //主CPC小组
	CPCSETS  = "CPCSETS"  //C-sets
	GBC      = "GBC"      //GBC分类号
	GBCN     = "GBCN"     //GBC分类数
	GBC1     = "GBC1"     //GBC门类
	GBC2     = "GBC2"     //GBC大类
	GBC3     = "GBC3"     //GBC中类
	GBC4     = "GBC4"     //GBC小类
	PGBC     = "PGBC"     //GBC主分类号
	PGBC1    = "PGBC1"    //主GBC门类
	PGBC2    = "PGBC2"    //主GBC大类
	PGBC3    = "PGBC3"    //主GBC中类
	PGBC4    = "PGBC4"    //主GBC小类
	EINDC    = "EINDC"    //EINDC分类号
	EINDCN   = "EINDCN"   //EINDC分类数
	EINDC1   = "EINDC1"   //EINDC门类
	EINDC2   = "EINDC2"   //EINDC大类
	UC       = "UC"       //UC分类号
	PUC      = "PUC"      //UC主分类号
	UCN      = "UCN"      //UC分类数
	FIC      = "FIC"      //FI分类号
	PFIC     = "PFIC"     //FI主分类号
	FICN     = "FICN"     //FI分类数
	FTERM    = "FTERM"    //FTERM分类号
	PFTERM   = "PFTERM"   //FTERM主分类号
	FTERMN   = "FTERMN"   //FTERM分类数
	IDX      = "IDX"      //DPI
	DPIT     = "DPIT"     //技术价值(DPIT)
	DPIL     = "DPIL"     //法律价值(DPIL)
	DPIE     = "DPIE"     //经济价值(DPIE)
	DPIM     = "DPIM"     //市场价值(DPIM)
	DPIS     = "DPIS"     //战略价值(DPIS)
	CO       = "CO"       //国省代码
	ED       = "ED"       //实审公告日
	GD       = "GD"       //授权公告日
	QLFQR    = "QLFQR"    //权利失效日
	AGC      = "AGC"      //专利代理机构
	AGT      = "AGT"      //代理人
	REFP     = "REFP"     //引证专利
	REFPN    = "REFPN"    //引证专利数
	REFPCN   = "REFPCN"   //引用专利国别/地域数
	REFNPN   = "REFNPN"   //引用非专利文献数
	REFBYN   = "REFBYN"   //被引证数
	REFBSCYN = "REFBSCYN" //被审查员引证数
	IAN      = "IAN"      //国际申请
	IAD      = "IAD"      //国际申请日
	IPN      = "IPN"      //国际公布
	IPD      = "IPD"      //国际公开日
	DEN      = "DEN"      //进入国家日期
	SF       = "SF"       //检索领域
	LS       = "LS"       //法律状态履历
	CLS      = "CLS"      //当前法律状态
	CLSD     = "CLSD"     //当前法律状态公告日
	CFP      = "CFP"      //海关备案专利
	RO       = "RO"       //受理局
	PRCP     = "PRCP"     //转让人
	POCP     = "POCP"     //受让人
	TSFT     = "TSFT"     //转让类型
	TEFD     = "TEFD"     //转让生效日
	TLSD     = "TLSD"     //转让法律状态公告日
	TRSF     = "TRSF"     //许可人
	ASSI     = "ASSI"     //被许可人
	LICT     = "LICT"     //许可类型
	LEFD     = "LEFD"     //许可生效日
	LCHD     = "LCHD"     //许可变更日
	LDSD     = "LDSD"     //许可解除日
	LSBD     = "LSBD"     //许可法律状态公告日
	LSOC     = "LSOC"     //许可合同状态
	LCRN     = "LCRN"     //许可合同备案号
	PLDO     = "PLDO"     //出质人
	PLDE     = "PLDE"     //质权人
	PLDT     = "PLDT"     //质押保全类型
	PEFD     = "PEFD"     //质押生效日
	PCHD     = "PCHD"     //质押变更日
	PDSD     = "PDSD"     //质押解除日
	PSBD     = "PSBD"     //质押法律状态公告日
	PSOC     = "PSOC"     //质押合同状态
	PCRN     = "PCRN"     //质押合同登记号
	LV       = "LV"       //有效性
	TZH      = "TZH"      //同族/同族号
	SFMLC    = "SFMLC"    //布局国家/地域
	ITZH     = "ITZH"     //INNOJOY同族
	IFMLN    = "IFMLN"    //INNOJOY同族数
	IFMLC    = "IFMLC"    //INNOJOY同族布局国家/地域
	IFMLCN   = "IFMLCN"   //INNOJOY布局国家/地域数
	BZ       = "BZ"       //标准专利
	WINPS    = "WINPS"    //专利奖
	WINPDY   = "WINPDY"   //获奖年份
	EAR      = "EAR"      //复审决定
	CRAI     = "CRAI"     //无效决定
	ETSIBZ   = "ETSIBZ"   //ETSI标准
	ETSIN    = "ETSIN"    //标准号
	BZC      = "BZC"      //标准申报公司
	BZD      = "BZD"      //标准申报日期
	HYLY     = "HYLY"     //行业领域
	SQRLX    = "SQRLX"    //申请人类型
	TRA      = "TRA"      //转让次数
	LIC      = "LIC"      //许可次数
	PLE      = "PLE"      //质押次数
	INV      = "INV"      //无效次数
	PERIOD   = "PERIOD"   //存活期
	RET      = "RET"      //剩余有效期
	PTAD     = "PTAD"     //期限调整PTA
	PTED     = "PTED"     //期限延长PTE
	PTEN     = "PTEN"     //药品名称
	PB       = "PB"       //药品品牌
	PI       = "PI"       //药品成分
	PF       = "PF"       //药品剂型
	PRT      = "PRT"      //药品用法
	PST      = "PST"      //药品用量
)
