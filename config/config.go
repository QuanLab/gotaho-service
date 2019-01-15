package config

type Config struct {
	MySQL      MySQL      `json:"MySQL,omitempty"`
	Server     Server     `json:"Server,omitempty"`
	Carte      Carte      `json:"Carte,omitempty"`
	MailServer MailServer `json:"MailServer,omitempty"`
}

type MySQL struct {
	Username  string `json:"Username,omitempty"`
	Password  string `json:"Password,omitempty"`
	Host      string `json:"Host,omitempty"`
	Port      int    `json:"Port,omitempty"`
	Database  string `json:"Database,omitempty"`
	Parameter string `json:"Parameter,omitempty"`
}

type Server struct {
	Port             int    `json:"Port,omitempty"`
	BaseContextPath  string `json:"BaseContextPath,omitempty"`
	PrivatePassPhase string `json:"PrivatePassPhase,omitempty"`
	Salt             string `json:"Salt,omitempty"`
}

type Carte struct {
	KettleStatus string `json:"KettleStatus,omitempty"`
	StartJobURL  string `json:"StartJobURL,omitempty"`
	StopJobURL   string `json:"StopJobURL,omitempty"`
	JobStatusURL string `json:"JobStatusURL,omitempty"`
	JobImageURL  string `json:"JobImageURL,omitempty"`
}

type MailServer struct {
	Domain     string `json:"Domain,omitempty"`
	Server     string `json:"Server,omitempty"`
	Port       int    `json:"Port,omitempty"`
	Email      string `json:"Email,omitempty"`
	Password   string `json:"Password,omitempty"`
	SenderName string `json:"SenderName,omitempty"`
}
