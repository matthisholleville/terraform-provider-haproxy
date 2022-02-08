package provider

import (
	"context"

	"github.com/avast/retry-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/matthisholleville/terraform-provider-haproxy/internal/haproxy"
	"github.com/matthisholleville/terraform-provider-haproxy/internal/haproxy/models"
)

func resourceFrontend() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFrontendCreate,
		ReadContext:   resourceFrontendRead,
		UpdateContext: resourceFrontendUpdate,
		DeleteContext: resourceFrontendDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
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

func resourceFrontendRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*haproxy.Client)

	frontend := models.Frontend{
		Name: d.Id(),
	}

	result, err := client.GetFrontend(frontend)
	if err != nil {
		return diag.FromErr(err)

	}

	d.Set("name", result.Name)
	d.Set("bind_process", result.BindProcess)
	d.Set("clflog", result.Clflog)
	d.Set("client_timeout", result.ClientTimeout)
	d.Set("clitcpka", result.Clitcpka)
	d.Set("contstats", result.Contstats)
	d.Set("default_backend", result.DefaultBackend)
	d.Set("dontlognull", result.Dontlognull)
	d.Set("forwardfor", result.Forwardfor)
	d.Set("http_buffer_request", result.HttpBufferRequest)
	d.Set("http_use_htx", result.HttpUseHtx)
	d.Set("http_connection_mode", result.HttpConnectionMode)
	d.Set("http_keep_alive_timeout", result.HttpKeepAliveTimeout)
	d.Set("http_request_timeout", result.HttpRequestTimeout)
	d.Set("httplog", result.HttpLog)
	d.Set("log_format", result.LogFormat)
	d.Set("log_format_sd", result.LogFormatSd)
	d.Set("log_separate_errors", result.LogSeparateErrors)
	d.Set("log_tag", result.LogTag)
	d.Set("logasap", result.Logasap)
	d.Set("maxconn", result.MaxConn)
	d.Set("mode", result.Mode)
	d.Set("monitor_fail", result.MonitorFail)
	d.Set("monitor_uri", result.MonitorUri)

	return nil
}

func resourceFrontendCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*haproxy.Client)
	frontend := *buildFrontendFromResourceParameters(d)

	err := retry.Do(
		func() error {
			configuration, err := client.GetConfiguration()
			if err != nil {
				return err
			}
			transaction, err := client.CreateTransaction(configuration.Version)
			if err != nil {
				return err
			}

			_, err = client.CreateFrontend(transaction.Id, frontend)
			if err != nil {
				return err
			}
			_, err = client.CommitTransaction(transaction.Id)
			if err != nil {
				return err
			}
			return nil
		},
	)

	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(frontend.Name)
	return nil
}

func resourceFrontendUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*haproxy.Client)

	frontend := *buildFrontendFromResourceParameters(d)
	err := retry.Do(
		func() error {
			configuration, err := client.GetConfiguration()
			if err != nil {
				return err
			}
			transaction, err := client.CreateTransaction(configuration.Version)
			if err != nil {
				return err
			}
			_, err = client.UpdateFrontend(transaction.Id, frontend)
			if err != nil {
				return err
			}
			_, err = client.CommitTransaction(transaction.Id)
			if err != nil {
				return err
			}
			return nil
		},
	)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceFrontendDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*haproxy.Client)

	frontend := *buildFrontendFromResourceParameters(d)

	err := retry.Do(
		func() error {
			configuration, err := client.GetConfiguration()
			if err != nil {
				return err
			}
			transaction, err := client.CreateTransaction(configuration.Version)
			if err != nil {
				return err
			}

			err = client.DeleteFrontend(transaction.Id, frontend)
			if err != nil {
				return err
			}
			_, err = client.CommitTransaction(transaction.Id)
			if err != nil {
				return err
			}
			return nil
		},
	)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func buildFrontendFromResourceParameters(d *schema.ResourceData) *models.Frontend {
	frontend := &models.Frontend{}
	if v, ok := d.GetOk("bind_process"); ok {
		frontend.BindProcess = v.(string)
	}

	if v, ok := d.GetOk("clflog"); ok {
		frontend.Clflog = v.(bool)
	}

	if v, ok := d.GetOk("client_timeout"); ok {
		frontend.ClientTimeout = v.(int)
	}

	if v, ok := d.GetOk("clitcpka"); ok {
		frontend.Clitcpka = v.(string)
	}

	if v, ok := d.GetOk("contstats"); ok {
		frontend.Contstats = v.(string)
	}

	if v, ok := d.GetOk("default_backend"); ok {
		frontend.DefaultBackend = v.(string)
	}

	if v, ok := d.GetOk("dontlognull"); ok {
		frontend.Dontlognull = v.(string)
	}

	if v, ok := d.GetOk("forwardFor"); ok {
		forwardFor := models.Forwardfor{}
		forwardFor.Enabled = v.(map[string]interface{})["enabled"].(string)
		forwardFor.Except = v.(map[string]interface{})["except"].(string)
		forwardFor.Header = v.(map[string]interface{})["header"].(string)
		forwardFor.Ifnone = v.(map[string]interface{})["ifnone"].(bool)
	}

	if v, ok := d.GetOk("http_buffer_request"); ok {
		frontend.HttpBufferRequest = v.(string)
	}

	if v, ok := d.GetOk("http_use_htx"); ok {
		frontend.HttpUseHtx = v.(string)
	}

	if v, ok := d.GetOk("http_connection_mode"); ok {
		frontend.HttpConnectionMode = v.(string)
	}

	if v, ok := d.GetOk("http_keep_alive_timeout"); ok {
		frontend.HttpKeepAliveTimeout = v.(int)
	}

	if v, ok := d.GetOk("http_request_timeout"); ok {
		frontend.HttpRequestTimeout = v.(int)
	}

	if v, ok := d.GetOk("httplog"); ok {
		frontend.HttpLog = v.(bool)
	}

	if v, ok := d.GetOk("log_format"); ok {
		frontend.LogFormat = v.(string)
	}

	if v, ok := d.GetOk("log_format_sd"); ok {
		frontend.LogFormatSd = v.(string)
	}

	if v, ok := d.GetOk("log_separate_errors"); ok {
		frontend.LogSeparateErrors = v.(string)
	}

	if v, ok := d.GetOk("log_tag"); ok {
		frontend.LogTag = v.(string)
	}

	if v, ok := d.GetOk("logasap"); ok {
		frontend.Logasap = v.(string)
	}

	if v, ok := d.GetOk("maxconn"); ok {
		frontend.MaxConn = v.(int)
	}

	if v, ok := d.GetOk("mode"); ok {
		frontend.Mode = v.(string)
	}

	if v, ok := d.GetOk("monitor_fail"); ok {
		monitorFail := models.MonitorFail{}
		monitorFail.Cond = v.(map[string]interface{})["cond"].(string)
		monitorFail.CondTest = v.(map[string]interface{})["cond_test"].(string)
	}

	if v, ok := d.GetOk("monitor_uri"); ok {
		frontend.MonitorUri = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		frontend.Name = v.(string)
	}

	if v, ok := d.GetOk("stats_options"); ok {
		statsOptions := models.StatsOptions{}
		statsOptions.StatsEnable = v.(map[string]interface{})["stats_enable"].(bool)
		statsOptions.StatsHideVersion = v.(map[string]interface{})["stats_hide_version"].(bool)
		statsOptions.StatsMaxconn = v.(map[string]interface{})["stats_maxconn"].(int)
		statsOptions.StatsRefreshDelay = v.(map[string]interface{})["stats_refresh_delay"].(int)
		statsOptions.StatsShowDesc = v.(map[string]interface{})["stats_show_desc"].(string)
		statsOptions.StatsShowLegends = v.(map[string]interface{})["stats_show_legends"].(bool)
		statsOptions.StatsShowNodeName = v.(map[string]interface{})["stats_show_node_name"].(string)
		statsOptions.StatsUrilPrefix = v.(map[string]interface{})["stats_uri_prefix"].(string)
	}

	if v, ok := d.GetOk("tcplog"); ok {
		frontend.TcpLog = v.(bool)
	}

	if v, ok := d.GetOk("unique_id_format"); ok {
		frontend.UniqueIdFormat = v.(string)
	}

	if v, ok := d.GetOk("unique_id_header"); ok {
		frontend.UniqueIdHeader = v.(string)
	}

	return frontend
}
