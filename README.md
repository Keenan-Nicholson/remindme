# RemindMe! - A Discord Reminder Bot

Welcome to **RemindMe!**, a customizable reminder bot for your Discord server. This bot allows users to set reminders that trigger notifications at specified times using cron job scheduling.

---

## Installation

To install and run **RemindMe!**, follow these steps:

### Prerequisites

- [Go](https://golang.org/doc/install)
- [Discord Bot Token](https://discord.com/developers/applications)

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

1. Create a `.env` file in the root directory. See `.env.dist` for a template.
---

## Usage

**Commands**
1. `/settimer <duration> <days,mins,hours,seconds> <user> <reminder message>`
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
