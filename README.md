# RemindMe! - A Discord Reminder Bot

Welcome to **RemindMe!**, a customizable reminder bot for your Discord server. This bot allows users to set reminders that trigger notifications at specified times using cron job scheduling.

---

## Installation

To install and run **RemindMe!**, follow these steps:

### Prerequisites

- [Go](https://golang.org/doc/install)
- [Discord Bot Token](https://discord.com/developers/applications)
- [Docker](https://www.docker.com/get-started)

### Clone the Repository
```bash
$ git clone https://github.com/keenan-nicholson/remindme.git
$ cd remindme
```

### Install Dependencies (for local development)
```bash
$ go install
```

---

## Configuration

1. Create a `.env` file in the root directory. See `.env.dist` for a template.
2. If running locally (*without Docker*) set desired paths for the database in `/database/db.go` and logger in `/utils/logger.go`.

---

## Running the Bot with Docker

To run the bot using Docker Compose for an easy, reproducible setup:

1. Ensure your `.env` file is created with the required environment variables.
2. Use Docker Compose to build and run the bot:

   ```bash
   $ docker-compose up -d
   ```

   This command builds the Docker image and starts the bot service. Logs will be displayed in your terminal. The bot will automatically restart if it crashes.

---

## Usage

**Commands**
1. `/settimer <duration> <unit> <user> <reminder message>`
   - **Description**: Sets a timer that will trigger a reminder after a specified duration. The duration can be set in days, minutes, hours, or seconds.
   - **Example output**: `Hey Keeborg, this is your reminder: take a study break!`

2. `/setdate <year> <month> <day> <hour> <minute> <user> <reminder message>`
   - **Description**: Sets a reminder for a specific date and time. You provide the full date (year, month, day) along with the time (hour, minute). Currently only supports UTC.
   - **Example output**: `Hey Keeborg, this is your reminder: you have a dentist appointment today!`

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

