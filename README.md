# Linux Evangelist

A Discord bot written in Golang (made in 3 hours) to evangelize what Linux really is.

It receives real messages from [Stallman's words](https://stallman-copypasta.github.io) and prettifies them for your eyes.

Test it now!

## Requisites
Golang 1.23.3 (I downloaded it today)

### Third party Libraries
- bwmarrin/discordgo
- joho/godotenv

## How to run it
First, resolve all dependencies with:
```
go mod tidy
```
Then configure your .env following .env.example format:
```yaml
DISCORD_BOT_TOKEN=YOUR_DISCORD_BOT_TOKEN
```
And lastly, run it
```
go run .
```

If you use [Nix](https://nixos.org/), you may alternatively simply run the following:

```bash
nix run github:csl-club/linux-eva
```
