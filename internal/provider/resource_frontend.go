package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/matthisholleville/terraform-provider-haproxy/internal/haproxy"
	"github.com/matthisholleville/terraform-provider-haproxy/internal/haproxy/models"
)

func resourceFrontend() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFrontendCreate,
		UpdateContext: resourceFrontendUpdate,
		DeleteContext: resourceFrontendDelete,
		Schema: map[string]*schema.Schema{
			"bind_process": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Limit visibility of an instance to a certain set of processes numbers. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#bind-process",
			},
			"clflog": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable logging of HTTP request, session state and timers. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#4.2-option%20httplog",
			},
			"client_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set the maximum inactivity time on the client side. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#4-timeout%20client",
			},
			"clitcpka": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Enable or disable the sending of TCP keepalive packets on the client side. Possible value : 'enabled' or 'disabled'. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#4.2-option%20clitcpka",
			},
			"contstats": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Enable continuous traffic statistics updates. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#4-option%20contstats",
			},
			"default_backend": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify the backend to use when no 'backend' rule has been matched. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#default_backend",
			},
			"dontlognull": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Enable or disable logging of null connections. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#option%20dontlognull",
			},
			"forwardfor": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Enable insertion of the X-Forwarded-For header to requests sent to servers. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#4.2-option%20forwardfor",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Enable forwardfor. Possible value : 'enabled' or 'disabled'.",
						},
						"except": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "It is possible to disable the addition of the header for a known source address or network by adding the 'except' keyword followed by the network address.",
						},
						"header": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The keyword 'header' may be used to supply a different header name to replace the default 'X-Forwarded-For'. This can be useful where you might already have a 'X-Forwarded-For' header from a different application (e.g. stunnel), and you need preserve it.",
						},
						"ifnone": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "the keyword 'if-none' states that the header will only be added if it is not present. This should only be used in perfectly trusted environment, as this might cause a security issue if headers reaching haproxy are under the control of the end-user.",
						},
					},
				},
			},
			"http_buffer_request": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Enable or disable waiting for whole HTTP request body before proceeding. Possible value: 'enabled' or 'disabled'. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#4.2-option%20http-buffer-request",
			},
			"http_use_htx": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Enable or disable htx option. Possible value: 'enabled' or 'disabled'. https://www.haproxy.com/fr/blog/haproxy-2-0-and-beyond/",
			},
			"http_connection_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "HAProxy connection mode. Possible value : 'httpclose' or 'http-server-close' or 'http-keep-alive'",
			},
			"http_keep_alive_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set the maximum inactivity time on the client and server side for tunnels. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#4.2-timeout%20tunnel",
			},
			"http_request_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set the maximum allowed time to wait for a complete HTTP request. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#timeout%20http-request",
			},
			"httplog": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable logging of HTTP request, session state and timers. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#4.2-option%20httplog",
			},
			"log_format": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "HAProxy log format. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#8.2.4",
			},
			"log_format_sd": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "HAProxy log format. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#8.2.4",
			},
			"log_separate_errors": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "HAProxy log separate errors. Possible value 'enabled' or 'disabled'. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#8.2.5",
			},
			"log_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the log tag to use for all outgoing logs. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#4.2-log-tag",
			},
			"logasap": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Enable or disable early logging. Possible value 'enabled' or 'disabled'. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#4.2-option%20logasap",
			},
			"maxconn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Limits the sockets to this number of concurrent connections. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#5.1-maxconn",
			},
			"mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Sets the octal mode used to define access permissions on the UNIX socket. Possible value 'http' or 'tcp'. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#5.1-mode",
			},
			"monitor_fail": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Add a condition to report a failure to a monitor HTTP request. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#4-monitor%20fail",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cond": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The monitor request will fail if the condition is satisfied, and will succeed otherwise.",
						},
						"cond_test": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The monitor request will succeed only if the condition is satisfied, and will fail otherwise.",
						},
					},
				},
			},
			"monitor_uri": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Intercept a URI used by external components' monitor requests. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#4.2-monitor-uri",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Frontend name",
			},
			"stats_options": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "HAProxy stats options.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stats_enable": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If true, stats will be enable. ",
						},
						"stats_hide_version": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable statistics and hide HAProxy version reporting. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#4-stats%20hide-version",
						},
						"stats_maxconn": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "By default, the stats socket is limited to 10 concurrent connections. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#3.1-stats%20maxconn",
						},
						"stats_refresh_delay": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable statistics with automatic refresh. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#4-stats%20refresh",
						},
						"stats_show_desc": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Enable reporting of a description on the statistics page. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#4.2-stats%20show-desc",
						},
						"stats_show_legends": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable reporting additional information on the statistics page. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#4.2-stats%20show-legends",
						},
						"stats_show_node_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Enable reporting of a host name on the statistics page. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#4.2-stats%20show-node",
						},
						"stats_uri_prefix": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Enable statistics and define the URI prefix to access them. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#4.2-stats%20uri",
						},
					},
				},
			},
			"tcplog": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable advanced logging of TCP connections with session state and timers. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#4-option%20tcplog",
			},
			"unique_id_format": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Generate a unique ID for each request. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#4.2-unique-id-format",
			},
			"unique_id_header": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Add a unique ID header in the HTTP request. https://cbonte.github.io/haproxy-dconv/1.8/configuration.html#4.2-unique-id-header",
			},
		},
	}
}

func resourceFrontendCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*haproxy.Client)
	transaction, err := client.CreateTransaction(0)
	if err != nil {
		return diag.FromErr(err)
	}

	frontend := *buildFrontendFromResourceParameters(d)
	_, err = client.CreateFrontend(transaction.Id, frontend)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(frontend.Name)
	return nil
}

func resourceFrontendUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*haproxy.Client)
	transaction, err := client.CreateTransaction(0)
	if err != nil {
		return diag.FromErr(err)
	}

	frontend := *buildFrontendFromResourceParameters(d)
	_, err = client.UpdateFrontend(transaction.Id, frontend)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceFrontendDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*haproxy.Client)
	transaction, err := client.CreateTransaction(0)
	if err != nil {
		return diag.FromErr(err)
	}

	frontend := *buildFrontendFromResourceParameters(d)
	err = client.DeleteFrontend(transaction.Id, frontend)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}

func buildFrontendFromResourceParameters(d *schema.ResourceData) *models.Frontend {
	bindProcess := d.Get("bind_process").(string)
	clflog := d.Get("clflog").(bool)
	clientTimeout := d.Get("client_timeout").(int)
	clitcpka := d.Get("clitcpka").(string)
	contstats := d.Get("contstats").(string)
	defaultBackend := d.Get("default_backend").(string)
	dontLogNull := d.Get("dontlognull").(string)
	forwardForSet := d.Get("forwardFor").(map[string]interface{})
	forwardFor := &models.Forwardfor{
		Enabled: forwardForSet["enabled"].(string),
		Except:  forwardForSet["except"].(string),
		Header:  forwardForSet["header"].(string),
		Ifnone:  forwardForSet["ifnone"].(bool),
	}
	httpBufferRequest := d.Get("http_buffer_request").(string)
	httpUseHtx := d.Get("http_use_htx").(string)
	httpConnectionMode := d.Get("http_connection_mode").(string)
	httpKeepAliveTimeout := d.Get("http_keep_alive_timeout").(int)
	httpRequestTimeout := d.Get("http_request_timeout").(int)
	httpLog := d.Get("httplog").(bool)
	logFormat := d.Get("log_format").(string)
	logFormatSd := d.Get("log_format_sd").(string)
	logSeparateErrors := d.Get("log_separate_errors").(string)
	logTag := d.Get("log_tag").(string)
	logasap := d.Get("logasap").(string)
	maxconn := d.Get("maxconn").(int)
	mode := d.Get("mode").(string)
	monitorFailSet := d.Get("monitor_fail").(map[string]interface{})
	monitorFail := &models.MonitorFail{
		Cond:     monitorFailSet["cond"].(string),
		CondTest: monitorFailSet["cond_test"].(string),
	}
	monitorUri := d.Get("monitor_uri").(string)
	name := d.Get("name").(string)
	statsOptionsSet := d.Get("stats_options").(map[string]interface{})
	statsOptions := &models.StatsOptions{
		StatsEnable:       statsOptionsSet["stats_enable"].(bool),
		StatsHideVersion:  statsOptionsSet["stats_hide_version"].(bool),
		StatsMaxconn:      statsOptionsSet["stats_maxconn"].(int),
		StatsRefreshDelay: statsOptionsSet["stats_refresh_delay"].(int),
		StatsShowDesc:     statsOptionsSet["stats_show_desc"].(string),
		StatsShowLegends:  statsOptionsSet["stats_show_legends"].(bool),
		StatsShowNodeName: statsOptionsSet["stats_show_node_name"].(string),
		StatsUrilPrefix:   statsOptionsSet["stats_uri_prefix"].(string),
	}
	tcplog := d.Get("tcplog").(bool)
	uniqueIdFormat := d.Get("unique_id_format").(string)
	uniqueIdHeader := d.Get("unique_id_header").(string)

	return &models.Frontend{
		BindProcess:          bindProcess,
		Clflog:               clflog,
		ClientTimeout:        clientTimeout,
		Clitcpka:             clitcpka,
		Contstats:            contstats,
		DefaultBackend:       defaultBackend,
		Dontlognull:          dontLogNull,
		Forwardfor:           *forwardFor,
		HttpBufferRequest:    httpBufferRequest,
		HttpUseHtx:           httpUseHtx,
		HttpConnectionMode:   httpConnectionMode,
		HttpKeepAliveTimeout: httpKeepAliveTimeout,
		HttpRequestTimeout:   httpRequestTimeout,
		HttpLog:              httpLog,
		LogFormat:            logFormat,
		LogFormatSd:          logFormatSd,
		LogSeparateErrors:    logSeparateErrors,
		LogTag:               logTag,
		Logasap:              logasap,
		MaxConn:              maxconn,
		Mode:                 mode,
		MonitorFail:          *monitorFail,
		MonitorUri:           monitorUri,
		Name:                 name,
		StatsOptions:         *statsOptions,
		TcpLog:               tcplog,
		UniqueIdFormat:       uniqueIdFormat,
		UniqueIdHeader:       uniqueIdHeader,
	}
}
