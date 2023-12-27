结构体定义如下:
```python
type OpenPort struct {
	AgentID   string
	AgentIP   string
	AgentPort int
	Proto     string
}

type TgtPot struct {
	PotIP   string
	PotPort int
	// IsRealDev bool
}

type SpecPotInfo struct {
	Proto       string
	NeedPayload bool
}

type Strategy struct {
	OpenPort OpenPort `json:"OpenPort"`
	Realdev  TgtPot   `json:"Realdev"`
	Emudev   TgtPot   `json:"Emudev"`
}

type ServerCfg struct {
	Strategys   *[]Strategy
	SpecPotList *map[int]SpecPotInfo
}

var SrvCfg ServerCfg
```

解析文件示例:
```python
strategys:
- openport:
    agentid: 4ecdb6c3fda547b9b7ff3fb0a220804e
    agentip: 10.19.201.231
    agentport: 11121
    proto: tcp
  realdev:
    potip: "10.19.196.133"
    potport: 80
    proto: "http"
  emudev:
    potip: 10.19.196.36
    potport: 80
    potid: "1"
- openport:
    agentid: 4ecdb6c3fda547b9b7ff3fb0a220804e
    agentip: 10.19.201.231
    agentport: 23
    proto: tcp
  realdev:
    potip: ""
    potport: 0
    proto: ""
  emudev:
    potip: 10.19.196.40
    potport: 22
specpotlist:
  9091:
    proto: spec
    needpayload: true
```
