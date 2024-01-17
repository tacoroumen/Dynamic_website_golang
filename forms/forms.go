package forms

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
)

type EmailConfig struct {
	Mailadress  string `json:"mailadress"`
	Wachtwoord  string `json:"wachtwoord"`
	Smtp_server string `json:"smtp_server"`
	Smtp_poort  int    `json:"smtp_poort"`
}

type config struct {
	Server   string `json:"server"`
	UserSQL  string `json:"userSql"`
	Password string `json:"password"`
	Database string `json:"database"`
}

var loadedTemplate *template.Template

func init() {
	templatesPath := filepath.Join("templates", "reservering.html")
	loadedTemplate = template.Must(template.ParseFiles(templatesPath))
}

func LoadConfig(filename string) (*config, error) {
	config := &config{}

	file, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("fout bij het lezen van het configuratiebestand: %v", err)
		return nil, err
	}

	if err := json.Unmarshal(file, config); err != nil {
		log.Printf("fout bij het decoderen van JSON: %v", err)
		return nil, err
	}

	return config, nil
}

func azureSql() (*sql.DB, error) {
	conf, err := LoadConfig("config.json")
	if err != nil {
		log.Fatal("fout bij het laden van configuratie:", err)
	}

	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?tls=true", conf.UserSQL, conf.Password, conf.Server, conf.Database)

	db, err := sql.Open("mysql", connectionString)
	return db, err
}

func HandleForms(w http.ResponseWriter, r *http.Request) {
	var email string
	var password string
	err := loadedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("error executing Forms:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	// Make a connection to the database using azureSql function
	db, err := azureSql()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if r.FormValue("create_account") == "on" {
		url := fmt.Sprintf("https://gatekeeper.fonteynholidayparks.com/user/add?firstname=%s&lastname=%s&email=%s&phone=%s&password=%s&postalcode=%s&housenumber=%s&street=%s&town=%s&country=%s&birthdate=%s&licenseplate=%s", r.FormValue("first_name"), r.FormValue("last_name"), r.FormValue("email"), r.FormValue("phone"), r.FormValue("password"), r.FormValue("postalcode"), r.FormValue("house_number"), r.FormValue("street"), r.FormValue("city"), r.FormValue("country"), r.FormValue("birthdate"), r.FormValue("licenseplate"))
		response, err := http.Post(url, "application/json", nil)
		email = r.FormValue("email")
		password = r.FormValue("password")
		if err != nil {
			fmt.Println("Error making POST request:", err)
			return
		}
		defer response.Body.Close()

		// Check the response status code
		if response.StatusCode != http.StatusOK {
			fmt.Printf("Error: Unexpected status code %d\n", response.StatusCode)
			return
		}
	}

	if r.FormValue("already_have_account") == "on" {
		url := fmt.Sprintf("https://gatekeeper.fonteynholidayparks.com/login?email=%s&password=%s", r.FormValue("email_existing"), r.FormValue("password_exisisting"))
		response, err := http.Get(url)
		email = r.FormValue("email_existing")
		password = r.FormValue("password_exisisting")
		if err != nil {
			fmt.Println("Error making GET request:", err)
			return
		}
		defer response.Body.Close()

		// Check the response status code
		if response.StatusCode != http.StatusOK {
			fmt.Printf("Error: Unexpected status code %d\n", response.StatusCode)
			return
		}
	}

	// get the firstname of the user
	var firstname string
	url := fmt.Sprintf("https://gatekeeper.fonteynholidayparks.com/user/get?email=%s&password=%s", email, password)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		fmt.Printf("Error: Unexpected status code %d\n", response.StatusCode)
		return
	}

	var userData map[string]interface{}
	responseapi := json.NewDecoder(response.Body)
	err = responseapi.Decode(&userData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Extract firstname from the userData map
	if val, ok := userData["firstname"].(string); ok {
		firstname = val
	} else {
		fmt.Println("Firstname not found in the JSON response")
	}

	// Create reservation
	url = fmt.Sprintf("https://gatekeeper.fonteynholidayparks.com/reservering?housenumber=%s&checkin=%s&checkout=%s&email=%s&password=%s", r.FormValue("house_number_park"), r.FormValue("checkin"), r.FormValue("checkout"), email, password)
	response, err = http.Post(url, "application/json", nil)
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer response.Body.Close()

	// Open and read the email configuration from a JSON file
	file, err := os.Open("mail.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := EmailConfig{}
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
		return
	}

	auth := smtp.PlainAuth("", config.Mailadress, config.Wachtwoord, config.Smtp_server)

	to := []string{email}
	subject := "Subject: " + "Reservation FonteynHolidayParks!" + "\r\n"
	body := "Hello, " + firstname + "\r\n" +
		"Your reservation for hous: " + r.FormValue("house_number_park") + "\r\n" +
		"from: " + r.FormValue("checkin") + " to " + r.FormValue("checkout") + " was succesful" + "\r\n" +
		"We are happy to recieve you " + r.FormValue("checkin") + "\r\n" + "\r\n" +
		"For more information please contact us by responding to this email" + "\r\n" + "\r\n" +
		"Kind regards," + "\r\n" +
		"FonteynHolidayParks"

	msg := []byte(subject +
		"\r\n" +
		body)

	err = smtp.SendMail(fmt.Sprintf("%s:%d", config.Smtp_server, config.Smtp_poort), auth, config.Mailadress, to, msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email sent successfully!")
}
