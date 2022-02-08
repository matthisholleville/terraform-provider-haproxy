package models

type Frontend struct {
	BindProcess          string       `json:"bind_process,omitempty"`
	Clflog               bool         `json:"clflog,omitempty"`
	ClientTimeout        int          `json:"client_timeout,omitempty"`
	Clitcpka             string       `json:"clitcpka,omitempty"`
	Contstats            string       `json:"contstats,omitempty"`
	DefaultBackend       string       `json:"default_backend,omitempty"`
	Dontlognull          string       `json:"dontlognull,omitempty"`
	Forwardfor           Forwardfor   `json:"forwardfor,omitempty"`
	HttpBufferRequest    string       `json:"http-buffer-request,omitempty"`
	HttpUseHtx           string       `json:"http-use-htx,omitempty"`
	HttpConnectionMode   string       `json:"http_connection_mode,omitempty"`
	HttpKeepAliveTimeout int          `json:"http_keep_alive_timeout,omitempty"`
	HttpRequestTimeout   int          `json:"http_request_timeout,omitempty"`
	HttpLog              bool         `json:"httplog,omitempty"`
	LogFormat            string       `json:"log_format,omitempty"`
	LogFormatSd          string       `json:"log_format_sd,omitempty"`
	LogSeparateErrors    string       `json:"log_separate_errors,omitempty"`
	LogTag               string       `json:"log_tag,omitempty"`
	Logasap              string       `json:"logasap,omitempty"`
	MaxConn              int          `json:"maxconn,omitempty"`
	Mode                 string       `json:"mode,omitempty"`
	MonitorFail          MonitorFail  `json:"monitor_fail,omitempty"`
	MonitorUri           string       `json:"monitor_uri,omitempty"`
	Name                 string       `json:"name"`
	StatsOptions         StatsOptions `json:"stats_options,omitempty"`
	TcpLog               bool         `json:"tcplog,omitempty"`
	UniqueIdFormat       string       `json:"unique_id_format,omitempty"`
	UniqueIdHeader       string       `json:"unique_id_header,omitempty"`
}

type Forwardfor struct {
	Enabled string `json:"enabled"`
	Except  string `json:"except,omitempty"`
	Header  string `json:"header,omitempty"`
	Ifnone  bool   `json:"ifnone,omitempty"`
}

type MonitorFail struct {
	Cond     string `json:"cond"`
	CondTest string `json:"cond_test"`
}

type StatsOptions struct {
	StatsEnable       bool   `json:"stats_enable,omitempty"`
	StatsHideVersion  bool   `json:"stats_hide_version,omitempty"`
	StatsMaxconn      int    `json:"stats_maxconn,omitempty"`
	StatsRefreshDelay int    `json:"stats_refresh_delay,omitempty"`
	StatsShowDesc     string `json:"stats_show_desc,omitempty"`
	StatsShowLegends  bool   `json:"stats_show_legends,omitempty"`
	StatsShowNodeName string `json:"stats_show_node_name,omitempty"`
	StatsUrilPrefix   string `json:"stats_uri_prefix,omitempty"`
}
