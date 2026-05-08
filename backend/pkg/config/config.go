package config

import (
	"net/url"
	"reflect"
	"regexp"
	"strings"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/vxcontrol/cloud/anonymizer/patterns"
)

type Config struct {
	// === Core System Configuration ===
	DatabaseURL string `env:"DATABASE_URL" envDefault:"postgres://pentagiuser:pentagipass@pgvector:5432/pentagidb?sslmode=disable"`
	Debug       bool   `env:"DEBUG" envDefault:"false"`
	DataDir     string `env:"DATA_DIR" envDefault:"./data"`
	AskUser     bool   `env:"ASK_USER" envDefault:"false"`

	// === Container Runtime Configuration ===
	DockerInside                 bool   `env:"DOCKER_INSIDE" envDefault:"false"`
	DockerNetAdmin               bool   `env:"DOCKER_NET_ADMIN" envDefault:"false"`
	DockerSocket                 string `env:"DOCKER_SOCKET"`
	DockerNetwork                string `env:"DOCKER_NETWORK"`
	DockerPublicIP               string `env:"DOCKER_PUBLIC_IP" envDefault:"0.0.0.0"`
	DockerWorkDir                string `env:"DOCKER_WORK_DIR"`
	DockerDefaultImage           string `env:"DOCKER_DEFAULT_IMAGE" envDefault:"debian:latest"`
	DockerDefaultImageForPentest string `env:"DOCKER_DEFAULT_IMAGE_FOR_PENTEST" envDefault:"vxcontrol/kali-linux"`

	// === API Server Configuration ===
	ServerPort   int    `env:"SERVER_PORT" envDefault:"8080"`
	ServerHost   string `env:"SERVER_HOST" envDefault:"0.0.0.0"`
	ServerUseSSL bool   `env:"SERVER_USE_SSL" envDefault:"false"`
	ServerSSLKey string `env:"SERVER_SSL_KEY"`
	ServerSSLCrt string `env:"SERVER_SSL_CRT"`

	// === Frontend Static Assets ===
	StaticURL   *url.URL `env:"STATIC_URL"`
	StaticDir   string   `env:"STATIC_DIR" envDefault:"./fe"`
	CorsOrigins []string `env:"CORS_ORIGINS" envDefault:"*"`

	// === Authentication & Session Security ===
	CookieSigningSalt string `env:"COOKIE_SIGNING_SALT"`

	// === Web Scraper Service Endpoints ===
	ScraperPublicURL  string `env:"SCRAPER_PUBLIC_URL"`
	ScraperPrivateURL string `env:"SCRAPER_PRIVATE_URL"`

	// === LLM Provider: DeepSeek (Core Provider) ===
	DeepSeekAPIKey    string `env:"DEEPSEEK_API_KEY"`
	DeepSeekServerURL string `env:"DEEPSEEK_SERVER_URL" envDefault:"https://api.deepseek.com"`
	DeepSeekProvider  string `env:"DEEPSEEK_PROVIDER"`

	// === Vector Embedding Configuration ===
	EmbeddingURL           string `env:"EMBEDDING_URL"`
	EmbeddingKey           string `env:"EMBEDDING_KEY"`
	EmbeddingModel         string `env:"EMBEDDING_MODEL"`
	EmbeddingStripNewLines bool   `env:"EMBEDDING_STRIP_NEW_LINES" envDefault:"true"`
	EmbeddingBatchSize     int    `env:"EMBEDDING_BATCH_SIZE" envDefault:"512"`
	EmbeddingProvider      string `env:"EMBEDDING_PROVIDER" envDefault:"deepseek"`

	// === Chain Summarization Engine ===
	SummarizerPreserveLast   bool `env:"SUMMARIZER_PRESERVE_LAST" envDefault:"true"`
	SummarizerUseQA          bool `env:"SUMMARIZER_USE_QA" envDefault:"true"`
	SummarizerSumHumanInQA   bool `env:"SUMMARIZER_SUM_MSG_HUMAN_IN_QA" envDefault:"false"`
	SummarizerLastSecBytes   int  `env:"SUMMARIZER_LAST_SEC_BYTES" envDefault:"51200"`
	SummarizerMaxBPBytes     int  `env:"SUMMARIZER_MAX_BP_BYTES" envDefault:"16384"`
	SummarizerMaxQASections  int  `env:"SUMMARIZER_MAX_QA_SECTIONS" envDefault:"10"`
	SummarizerMaxQABytes     int  `env:"SUMMARIZER_MAX_QA_BYTES" envDefault:"65536"`
	SummarizerKeepQASections int  `env:"SUMMARIZER_KEEP_QA_SECTIONS" envDefault:"1"`

	// === Search Engine: DuckDuckGo (Core Search) ===
	DuckDuckGoEnabled    bool   `env:"DUCKDUCKGO_ENABLED" envDefault:"true"`
	DuckDuckGoRegion     string `env:"DUCKDUCKGO_REGION"`
	DuckDuckGoSafeSearch string `env:"DUCKDUCKGO_SAFESEARCH"`
	DuckDuckGoTimeRange  string `env:"DUCKDUCKGO_TIME_RANGE"`

	// === OAuth Callback Configuration ===
	PublicURL string `env:"PUBLIC_URL" envDefault:""`

	// === AI Assistant Mode Configuration ===
	AssistantUseAgents                bool `env:"ASSISTANT_USE_AGENTS" envDefault:"false"`
	AssistantSummarizerPreserveLast   bool `env:"ASSISTANT_SUMMARIZER_PRESERVE_LAST" envDefault:"true"`
	AssistantSummarizerLastSecBytes   int  `env:"ASSISTANT_SUMMARIZER_LAST_SEC_BYTES" envDefault:"76800"`
	AssistantSummarizerMaxBPBytes     int  `env:"ASSISTANT_SUMMARIZER_MAX_BP_BYTES" envDefault:"16384"`
	AssistantSummarizerMaxQASections  int  `env:"ASSISTANT_SUMMARIZER_MAX_QA_SECTIONS" envDefault:"7"`
	AssistantSummarizerMaxQABytes     int  `env:"ASSISTANT_SUMMARIZER_MAX_QA_BYTES" envDefault:"76800"`
	AssistantSummarizerKeepQASections int  `env:"ASSISTANT_SUMMARIZER_KEEP_QA_SECTIONS" envDefault:"3"`

	// === Network Proxy Settings ===
	ProxyURL string `env:"PROXY_URL"`

	// SSL Trusted CA Certificate Path (for external communication with LLM backends)
	ExternalSSLCAPath   string `env:"EXTERNAL_SSL_CA_PATH" envDefault:""`
	ExternalSSLInsecure bool   `env:"EXTERNAL_SSL_INSECURE" envDefault:"false"`

	// HTTP client timeout in seconds for external API calls (LLM providers, search tools, etc.)
	// A value of 0 means no timeout (not recommended).
	HTTPClientTimeout int `env:"HTTP_CLIENT_TIMEOUT" envDefault:"600"`

	// === Agent Execution Monitoring ===
	ExecutionMonitorEnabled        bool `env:"EXECUTION_MONITOR_ENABLED" envDefault:"false"`
	ExecutionMonitorSameToolLimit  int  `env:"EXECUTION_MONITOR_SAME_TOOL_LIMIT" envDefault:"5"`
	ExecutionMonitorTotalToolLimit int  `env:"EXECUTION_MONITOR_TOTAL_TOOL_LIMIT" envDefault:"10"`

	// === Agent Tool Execution Limits ===
	MaxGeneralAgentToolCalls int `env:"MAX_GENERAL_AGENT_TOOL_CALLS" envDefault:"100"`
	MaxLimitedAgentToolCalls int `env:"MAX_LIMITED_AGENT_TOOL_CALLS" envDefault:"20"`

	// === Agent Planning Phase Configuration ===
	AgentPlanningStepEnabled bool `env:"AGENT_PLANNING_STEP_ENABLED" envDefault:"false"`
}

func NewConfig() (*Config, error) {
	// Attempt to load .env file (silently ignore if not present)
	_ = godotenv.Load()

	var config Config
	if err := env.ParseWithOptions(&config, env.Options{
		RequiredIfNoDef: false,
		FuncMap: map[reflect.Type]env.ParserFunc{
			reflect.TypeOf(&url.URL{}): func(s string) (any, error) {
				if s == "" {
					return nil, nil
				}
				return url.Parse(s)
			},
		},
	}); err != nil {
		return nil, err
	}

	ensureInstallationID(&config)
	ensureLicenseKey(&config)

	return &config, nil
}

func ensureInstallationID(config *Config) {
	// Simplified: no installation ID needed for core version
}

func ensureLicenseKey(config *Config) {
	// Simplified: no license key needed for core version
}

// GetSecretPatterns returns a list of patterns for all secrets in the config
func (c *Config) GetSecretPatterns() []patterns.Pattern {
	var result []patterns.Pattern

	secrets := []struct {
		value string
		name  string
	}{
		{c.DatabaseURL, "Database URL"},
		{c.CookieSigningSalt, "Cookie Salt"},
		{c.DeepSeekAPIKey, "DeepSeek Key"},
		{c.EmbeddingKey, "Embedding Key"},
		{c.ProxyURL, "Proxy URL"},
	}

	for _, s := range secrets {
		trimmed := strings.TrimSpace(s.value)
		if trimmed == "" {
			continue
		}

		// escape regex special characters
		escaped := regexp.QuoteMeta(trimmed)
		pattern := patterns.Pattern{
			Name:  s.name,
			Regex: "(?P<replace>" + escaped + ")",
		}
		result = append(result, pattern)
	}

	return result
}
