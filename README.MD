
# SMS Engine Simulator (Golang)

## Overview
**SMS Engine Simulator** is a Go-based application designed to simulate the functionalities of an SMS engine. The project manages SMS routing, processing, and delivery through an internal simulation, making it ideal for testing and development purposes without the need for a live SMS gateway.

## Features
- **SMS Routing:** Simulates routing of messages based on predefined rules.
- **Database Management:** Utilizes SQLite for storing SMS and user data.
- **Concurrency:** Handles multiple SMS processing tasks concurrently.
- **Logging:** Provides detailed logs of SMS processing and delivery status.

## Installation
1. **Clone the Repository:**
   ```bash
   git clone https://github.com/Mohamed-Salem-Ali/SMS-Engine-Simulator-Golang.git
   ```
2. **Navigate to Project Directory:**
   ```bash
   cd SMS-Engine-Simulator-Golang
   ```
3. **Install Dependencies:**
   Ensure you have Go installed, then run:
   ```bash
   go mod tidy
   ```

## Usage
1. **Run the Simulator:**
   ```bash
   go run main.go
   ```
2. **Configuration:**
   - Adjust SMS routing rules and processing settings via configuration files in the `Config/` directory.
   - Modify database settings in `dbconfig.go`.

## Project Structure
- **SmsEngine1/:** Core logic for SMS processing.
- **DataBase/:** SQLite database schemas and related operations.
- **Config/:** Configuration files for the simulator.

## Contributing
Feel free to fork this repository, create a branch, and submit pull requests for any enhancements or bug fixes.

## License
This project is licensed under the MIT License. See the `LICENSE` file for more details.