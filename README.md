# FonteynHolidayParks Web Server

This is a simple web server written in Go for FonteynHolidayParks. It provides functionality for handling reservation forms and interacting with an external API. The server serves static files, handles different routes, and sends reservation confirmation emails.

## Getting Started

### Prerequisites
- [Go](https://golang.org/doc/install) installed

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/tacoroumen/Dynamic_website_golang
   ```

2. Change to the project directory:
   ```bash
   cd Dynamic_website_golang
   ```

3. Build and run the application:
   ```bash
   go run main.go
   ```

## Usage

### Routes
- `/`: Home page
- `/Aanvraag/`: Reservation form page
- `/Aanvraag/reservering/`: Handles reservation form submissions
- `/Overons/`: About us page

### Configuration
- The server listens on port 80 by default. You can change this by modifying the `addr` constant in `main.go`.

### Dependencies
- [github.com/go-sql-driver/mysql](https://pkg.go.dev/github.com/go-sql-driver/mysql): MySQL driver for Go

## Configuration Files

### `config.json`
This file contains configuration settings for connecting to the MySQL database.
```json
{ 
    "server": "Server ip/dns name",
    "usersql": "your_username_here",
    "password": "your_password_here",
    "database": "database_name"
}
```

### `mail.json`
This file includes email configuration details for sending reservation confirmation emails.
```json
{
    "mailadress": "your_email@gmail.com",
    "wachtwoord": "your_email_password_here",
    "smtp_server": "smtp.gmail.com",
    "smtp_poort": 587
}
```

## Contributing

Feel free to contribute to this project by opening issues or submitting pull requests.

## License

This project is licensed under the [MIT License](LICENSE).