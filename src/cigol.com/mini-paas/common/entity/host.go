package entity

type Host struct {
	// field should begin with uppercase letter, or it will not be exported
	Name     string `json:"name"`
	Ip       string `json:"ip"`
	SshPort  int    `json:"sshPort"`
	UserName string `json:"userName"`
	Passwd   string `json:"passwd"`
}

func NewHost(name, ip string, sshPort int, userName, passwd string) Host {
	return Host{Name: name, Ip: ip, SshPort: sshPort, UserName: userName, Passwd: passwd}
}
