package main

type ResSrvStatus struct {
	MinecraftNet     string `json:"minecraft.net"`
	SessionMinecraft string `json:"session.minecraft.net"`
	AccountMojang    string `json:"account.mojang.com"`
	AuthServMojang   string `json:"authserver.mojang.com"`
	SessionMojang    string `json:"sessionserver.mojang.com"`
	APIMojang        string `json:"api.mojang.com"`
	TextureMinecraft string `json:"textures.minecraft.net"`
	Mojang           string `json:"mojang.com"`
	Online 			 string `json:"online"`
}

type MCSrvStatus struct {
	Ip       string        `json:"ip"`
	Port     int           `json:"port"`
	Debug    DebugStruct   `json:"debug"`
	MOTD     MOTDStruct    `json:"motd"`
	Players  PlayersStruct `json:"players"`
	Version  string        `json:"version"`
	Online   bool          `json:"online"`
	Protocal int           `json:"protocal"`
	Hostname string        `json:"hostname"`
	Icon     string        `json:"icon"`
	Software string        `json:"software"`
}

type DebugStruct struct {
	Ping          bool `json:"ping"`
	Query         bool `json:"query"`
	Srv           bool `json:"srv"`
	QueryMismatch bool `json:"querymismatch"`
	IpInSrv       bool `json:"ipinsrv"`
	CNameInSrv    bool `json:"cnameinsrv"`
	AnimatedMOTD  bool `json:"animatedmotd"`
	CacheTime     int  `json:"cachetime"`
	ApiVersion    int  `json:"apiversion"`
}

type MOTDStruct struct {
	Raw   []string `json:"raw"`
	Clean []string `json:"clean"`
	Html  []string `json:"html"`
}

type PlayersStruct struct {
	Online int      `json:"online"`
	Max    int      `json:"max"`
	List   []string `json:"list"`
}
