package dto

type InnojoySearchResult struct {
	ReturnValue int    `json:"ReturnValue"`
	ErrorInfo   string `json:"ErrorInfo"`
	Option      struct {
		GUID        string `json:"GUID"`
		ReturnValue int    `json:"ReturnValue"`
		Meta        struct {
			PassTimeSearch int `json:"PassTimeSearch"`
			PassTimeRead   int `json:"PassTimeRead"`
			PassTime       int `json:"PassTime"`
		} `json:"Meta"`
		PatentList    []*PatentDetail `json:"PatentList"`
		Count         int             `json:"Count"`
		FamilyCount   int             `json:"FamilyCount"`
		IsFamilyMerge bool            `json:"isFamilyMerge"`
		ResultSection []struct {
			Database string `json:"Database"`
			DbName   string `json:"DbName"`
			Sort     int    `json:"Sort"`
			Display  bool   `json:"Display"`
			Count    int    `json:"Count"`
			Group    int    `json:"Group"`
		} `json:"ResultSection"`
	} `json:"Option"`
	Option2       string `json:"Option2"`
	Param         int    `json:"param"`
	TotalCount    int    `json:"TotalCount"`
	TotalCount2   int    `json:"TotalCount2"`
	Page          int    `json:"page"`
	Total         int    `json:"total"`
	Records       int    `json:"records"`
	IsGetCheckPic bool   `json:"IsGetCheckPic"`
	IsWriteMsg    bool   `json:"isWriteMsg"`
}

type PatentDetail struct {
	No         int           `json:"NO,omitempty"`
	Rno        int           `json:"RNO,omitempty"`
	Ftotal     int           `json:"FTOTAL,omitempty"`
	Fp         int           `json:"FP,omitempty"`
	Kc         string        `json:"KC,omitempty"`
	An         string        `json:"AN,omitempty"`
	ExamAN     string        `json:"ExamAN,omitempty"`
	Ad         string        `json:"AD,omitempty"`
	Pnm        string        `json:"PNM,omitempty"`
	Pnm2       string        `json:"PNM2,omitempty"`
	Pd         string        `json:"PD,omitempty"`
	Ti         string        `json:"TI,omitempty"`
	Abst       string        `json:"ABST,omitempty"`
	Cl         string        `json:"CL,omitempty"`
	Pa         string        `json:"PA,omitempty"`
	Cas        string        `json:"CAS,omitempty"`
	Ar         string        `json:"AR,omitempty"`
	Inn        string        `json:"INN,omitempty"`
	Pr         string        `json:"PR,omitempty"`
	Eprd       string        `json:"EPRD,omitempty"`
	Co         string        `json:"CO,omitempty"`
	Dan        string        `json:"DAN,omitempty"`
	Agc        string        `json:"AGC,omitempty"`
	Agt        string        `json:"AGT,omitempty"`
	Sic        string        `json:"SIC,omitempty"`
	Pic        string        `json:"PIC,omitempty"`
	Cpc        string        `json:"CPC,omitempty"`
	Cpcsets    string        `json:"CPCSETS,omitempty"`
	Ian        string        `json:"IAN,omitempty"`
	Ipn        string        `json:"IPN,omitempty"`
	Den        string        `json:"DEN,omitempty"`
	Dpn        string        `json:"DPN,omitempty"`
	Clm        string        `json:"CLM,omitempty"`
	Descr      string        `json:"DESCR,omitempty"`
	Db         string        `json:"DB,omitempty"`
	Ady        string        `json:"ADY,omitempty"`
	Pdy        string        `json:"PDY,omitempty"`
	Sicn       string        `json:"SICN,omitempty"`
	Ipcc       string        `json:"IPCC,omitempty"`
	Ipcsc      string        `json:"IPCSC,omitempty"`
	Picc       string        `json:"PICC,omitempty"`
	Picsc      string        `json:"PICSC,omitempty"`
	Picg       string        `json:"PICG,omitempty"`
	Picsg      string        `json:"PICSG,omitempty"`
	Ppa        string        `json:"PPA,omitempty"`
	Iad        string        `json:"IAD,omitempty"`
	Ipd        string        `json:"IPD,omitempty"`
	Qlfqr      string        `json:"QLFQR,omitempty"`
	Refp       string        `json:"REFP,omitempty"`
	Pan        int           `json:"PAN,omitempty"`
	Innn       int           `json:"INNN,omitempty"`
	Ifiid      int           `json:"IFIID,omitempty"`
	Fterm      string        `json:"FTERM,omitempty"`
	Sec        string        `json:"SEC,omitempty"`
	Fmln       int           `json:"FMLN,omitempty"`
	Ifmln      int           `json:"IFMLN,omitempty"`
	Sfmln      int           `json:"SFMLN,omitempty"`
	Sfmlcn     int           `json:"SFMLCN,omitempty"`
	Idx        int           `json:"IDX,omitempty"`
	Patms      string        `json:"PATMS,omitempty"`
	Inntms     string        `json:"INNTMS,omitempty"`
	Isnewdata  bool          `json:"ISNEWDATA,omitempty"`
	Clmn       int           `json:"CLMN,omitempty"`
	Refbyn     string        `json:"REFBYN,omitempty"`
	Refpn      int           `json:"REFPN,omitempty"`
	InnojoyDpi int           `json:"INNOJOY_DPI,omitempty"`
	Pinn       string        `json:"PINN,omitempty"`
	Pacc       string        `json:"PACC,omitempty"`
	Dpit       int           `json:"DPIT,omitempty"`
	Dpil       int           `json:"DPIL,omitempty"`
	Dpie       int           `json:"DPIE,omitempty"`
	Dpim       int           `json:"DPIM,omitempty"`
	Dpis       int           `json:"DPIS,omitempty"`
	Refpnl     int           `json:"REFPNL,omitempty"`
	Refpnf     int           `json:"REFPNF,omitempty"`
	Refpcn     int           `json:"REFPCN,omitempty"`
	Refbynl    int           `json:"REFBYNL,omitempty"`
	Refbynf    int           `json:"REFBYNF,omitempty"`
	Refnpn     int           `json:"REFNPN,omitempty"`
	Sicsn      int           `json:"SICSN,omitempty"`
	Siccn      int           `json:"SICCN,omitempty"`
	Fmlcn      int           `json:"FMLCN,omitempty"`
	Kv         int           `json:"KV,omitempty"`
	Pgn        int           `json:"PGN,omitempty"`
	Clmin      int           `json:"CLMIN,omitempty"`
	Lic        int           `json:"LIC,omitempty"`
	Cls        string        `json:"CLS,omitempty"`
	Clsd       string        `json:"CLSD,omitempty"`
	Ple        int           `json:"PLE,omitempty"`
	Tra        int           `json:"TRA,omitempty"`
	Inv        int           `json:"INV,omitempty"`
	Ppacc      string        `json:"PPACC,omitempty"`
	Addrp      string        `json:"ADDRP,omitempty"`
	Addrc      string        `json:"ADDRC,omitempty"`
	Addrdc     string        `json:"ADDRDC,omitempty"`
	Car        string        `json:"CAR,omitempty"`
	Ipcbn      string        `json:"IPCBN,omitempty"`
	Ipcn       string        `json:"IPCN,omitempty"`
	Ipcscn     string        `json:"IPCSCN,omitempty"`
	Iclmn      string        `json:"ICLMN,omitempty"`
	Clzs       string        `json:"CLZS,omitempty"`
	Ipc        string        `json:"IPC,omitempty"`
	Pipc       string        `json:"PIPC,omitempty"`
	Icg        string        `json:"ICG,omitempty"`
	Icsg       string        `json:"ICSG,omitempty"`
	Loc        string        `json:"LOC,omitempty"`
	Locn       string        `json:"LOCN,omitempty"`
	Cpcn       string        `json:"CPCN,omitempty"`
	Ecn        string        `json:"ECN,omitempty"`
	Uc         string        `json:"UC,omitempty"`
	Ucn        string        `json:"UCN,omitempty"`
	Fic        string        `json:"FIC,omitempty"`
	Ficn       string        `json:"FICN,omitempty"`
	Refbscyn   string        `json:"REFBSCYN,omitempty"`
	Bztz       int           `json:"BZTZ,omitempty"`
	Pcas       string        `json:"PCAS,omitempty"`
	ImageList  []interface{} `json:"imageList,omitempty"`
	Ptad       int           `json:"PTAD,omitempty"`
	Pas        string        `json:"PAS,omitempty"`
	Gbc        string        `json:"GBC,omitempty"`
	Gbcn       int           `json:"GBCN,omitempty"`
	Gbc1       string        `json:"GBC1,omitempty"`
	Gbc2       string        `json:"GBC2,omitempty"`
	Gbc3       string        `json:"GBC3,omitempty"`
	Gbc4       string        `json:"GBC4,omitempty"`
	Pgbc       string        `json:"PGBC,omitempty"`
	Pgbc1      string        `json:"PGBC1,omitempty"`
	Pgbc2      string        `json:"PGBC2,omitempty"`
	Pgbc3      string        `json:"PGBC3,omitempty"`
	Pgbc4      string        `json:"PGBC4,omitempty"`
	Eindc      string        `json:"EINDC,omitempty"`
	Eindc1     string        `json:"EINDC1,omitempty"`
	Eindc2     string        `json:"EINDC2,omitempty"`
	Eindcn     string        `json:"EINDCN,omitempty"`
	Casp       string        `json:"CASP,omitempty"`
	Casc       string        `json:"CASC,omitempty"`
	Casd       string        `json:"CASD,omitempty"`
	Ro         string        `json:"RO,omitempty"`
	Cfp        int           `json:"CFP,omitempty"`
	Spatms     string        `json:"SPATMS,omitempty"`
	Scastms    string        `json:"SCASTMS,omitempty"`
	Eprd0      string        `json:"eprd,omitempty"`
	Clm0       string        `json:"clm,omitempty"`
	Descr0     string        `json:"descr,omitempty"`
	Ady0       string        `json:"ady,omitempty"`
	Pdy0       string        `json:"pdy,omitempty"`
	Ipcc0      string        `json:"ipcc,omitempty"`
	Ipcsc0     string        `json:"ipcsc,omitempty"`
	Picc0      string        `json:"picc,omitempty"`
	Picsc0     string        `json:"picsc,omitempty"`
	Picg0      string        `json:"picg,omitempty"`
	Picsg0     string        `json:"picsg,omitempty"`
	Qlfqr0     string        `json:"qlfqr,omitempty"`
	Clsd0      string        `json:"clsd,omitempty"`
	Ppacc0     string        `json:"ppacc,omitempty"`
	Addrp0     string        `json:"addrp,omitempty"`
	Addrc0     string        `json:"addrc,omitempty"`
	Addrdc0    string        `json:"addrdc,omitempty"`
	Car0       string        `json:"car,omitempty"`
	Ipcbn0     string        `json:"ipcbn,omitempty"`
	Ipcn0      string        `json:"ipcn,omitempty"`
	Iclmn0     string        `json:"iclmn,omitempty"`
	Clzs0      string        `json:"clzs,omitempty"`
	Ipc0       string        `json:"ipc,omitempty"`
	Pipc0      string        `json:"pipc,omitempty"`
	Icg0       string        `json:"icg,omitempty"`
	Icsg0      string        `json:"icsg,omitempty"`
	Loc0       string        `json:"loc,omitempty"`
	Locn0      string        `json:"locn,omitempty"`
	Cpcn0      string        `json:"cpcn,omitempty"`
	Ecn0       string        `json:"ecn,omitempty"`
	Uc0        string        `json:"uc,omitempty"`
	Ucn0       string        `json:"ucn,omitempty"`
	Fic0       string        `json:"fic,omitempty"`
	Ficn0      string        `json:"ficn,omitempty"`
	Refbscyn0  string        `json:"refbscyn,omitempty"`

	IsClaimed bool `json:"isClaimed"`
	IsFocused bool `json:"isFocused"`
	PatentId  int  `json:"patentId"`
}

type SearchReq struct {
	Token              string              `json:"token"`
	PatentSearchConfig *PatentSearchConfig `json:"patentSearchConfig"`
}

type PatentSearchConfig struct {
	GUID      string `json:"GUID"`
	Action    string `json:"Action"`
	Query     string `json:"Query"`
	Database  string `json:"Database"`
	Page      string `json:"Page"`
	PageSize  string `json:"PageSize"`
	Sortby    string `json:"Sortby"`
	FieldList string `json:"FieldList"`
}

type StatisticReq struct {
	Token              string              `json:"token"`
	PatentSearchConfig *PatentSearchConfig `json:"patentSearchConfig"`
	AnalyseConfig      *AnalyseConfig      `json:"analyseConfig"`
	Language           string              `json:"language"`
}

type AnalyseConfig struct {
	AID string `json:"AID"`
}

type LoginReq struct {
	UserConfig UserConfig `json:"userConfig"`
}

type UserConfig struct {
	EMail    string `json:"EMail"`
	Password string `json:"Password"`
}

type LoginResp struct {
	ReturnValue int    `json:"ReturnValue"`
	Option      string `json:"Option"`
	ErrorInfo   string `json:"ErrorInfo"`
}
