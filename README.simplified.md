# PentAGI - Simplified DeepSeek Edition

**Penetration testing Artificial General Intelligence**

A streamlined version of PentAGI configured to run with just DeepSeek API, perfect for getting started quickly.

## Quick Start

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)
- DeepSeek API key (get one at [https://platform.deepseek.com/](https://platform.deepseek.com/))

### Setup

1. **Clone the repository**
```bash
git clone https://github.com/vxcontrol/pentagi.git
cd pentagi
```

2. **Configure environment**
```bash
cp .env.example .env
```

Edit `.env` and set your DeepSeek API key:
```bash
DEEPSEEK_API_KEY=your-api-key-here
```

**IMPORTANT**: Also change these security settings:
```bash
COOKIE_SIGNING_SALT=your-random-salt-here
PENTAGI_POSTGRES_PASSWORD=your-secure-password-here
```

3. **Start the platform**
```bash
docker compose up -d
```

4. **Access the UI**
Open https://localhost:8443 in your browser

On first access, you'll need to create an admin account.

## What's Included

This simplified version includes:
- ✅ **DeepSeek** as the LLM provider
- ✅ **PostgreSQL with pgvector** for data storage
- ✅ **DuckDuckGo** search engine (no API key needed)
- ✅ **Web scraper** for gathering information
- ✅ **Docker isolation** for safe penetration testing
- ✅ **Multi-agent system** (Researcher, Developer, Executor)

## What's Removed

To keep things simple, these features are not included:
- ❌ Multiple LLM providers (OpenAI, Anthropic, etc.)
- ❌ OAuth authentication (Google, GitHub)
- ❌ Advanced search engines (Tavily, Perplexity, etc.)
- ❌ Observability stack (Grafana, Langfuse)
- ❌ Knowledge graph (Graphiti/Neo4j)
- ❌ PentAGI Cloud integration

## Usage

1. **Create a new penetration test**
   - Click "New Flow" in the web UI
   - Enter your target and objectives
   - Let the AI agents work autonomously

2. **Monitor progress**
   - View real-time agent logs
   - See tool executions
   - Review findings

3. **Export results**
   - Download detailed vulnerability reports
   - Review exploitation guides

## Configuration

All configuration is in the `.env` file. Key settings:

```bash
# Required: DeepSeek API
DEEPSEEK_API_KEY=your-key-here

# Security (CHANGE THESE!)
COOKIE_SIGNING_SALT=change-this-salt-for-security
PENTAGI_POSTGRES_PASSWORD=change-this-postgres-password

# Optional: Search
DUCKDUCKGO_ENABLED=true

# Optional: Agent limits
MAX_GENERAL_AGENT_TOOL_CALLS=100
AGENT_PLANNING_STEP_ENABLED=false
```

## Troubleshooting

### DeepSeek API Issues
```bash
# Check if API key is set correctly
docker compose logs pentagi | grep -i deepseek
```

### Database Connection Issues
```bash
# Check database is running
docker compose ps pgvector

# Check logs
docker compose logs pgvector
```

### Port Already in Use
If port 8443 is already in use:
```bash
# Edit .env and change SERVER_PORT
SERVER_PORT=9443
```

## Stopping the Platform

```bash
# Stop all services
docker compose down

# Stop and remove all data
docker compose down -v
```

## Upgrading to Full Version

To use additional features like multiple LLM providers, observability, or knowledge graphs, see the [main README](https://github.com/vxcontrol/pentagi#readme) for the complete setup guide.

## Support

- Documentation: [CLAUDE.md](CLAUDE.md)
- Issues: [GitHub Issues](https://github.com/vxcontrol/pentagi/issues)
- Discord: [Join our community](https://discord.gg/2xrMh7qX6m)
- Telegram: [Join our channel](https://t.me/+Ka9i6CNwe71hMWQy)

## License

See [LICENSE](LICENSE)
