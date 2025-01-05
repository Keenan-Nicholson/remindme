# RemindMe! - A Discord Reminder Bot

Welcome to **RemindMe!**, a customizable reminder bot for your Discord server. This bot allows users to set reminders that trigger notifications at specified times using cron job scheduling.

---

## Installation

To install and run **RemindMe!**, follow these steps:

### Prerequisites

- [Go](https://golang.org/doc/install)
- A [Discord Bot Token](https://discord.com/developers/applications)

### Clone the Repository
```bash
$ git clone https://github.com/keenan-nicholson/remindme.git
$ cd remindme
```

### Install Dependencies
```bash
$ go install
```

---

## Configuration

1. Create a `.env` file in the root directory:
   ```
   DISCORD_BOT_TOKEN=your_bot_token_here
   ```
---

## Usage

1. **Add the Bot to Your Server**
   - Invite the bot using your custom bot link.

2. **Commands**
   - `remindme!<time>: <message>`: Sets a reminder for the specified time.
   - Example: `remindme!5 days: Ask team members if they are on schedule`

---

## Development

### Run the Bot Locally
```bash
$ go run main.go
```

---

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your improvements.

---

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

---
