# Polling App

Polling App is a web application that allows users to create and vote on polls. This project is built using Go.

## Features

- Create polls with multiple options
- Vote on polls
- View poll results
- Prevent multiple votes from the same user

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/polling-app.git
    cd polling-app
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Set up the database:
    - Ensure you have a running instance of your database.
    - Update the database configuration in `database/config.go`.

4. Run the application:
    ```sh
    go run main.go
    ```

## Usage

- Access the home page at `http://localhost:8080/`.
- Create a new poll by filling out the form on the home page.
- Vote on a poll by selecting an option and submitting your vote.
- View the results of a poll after voting.

## Project Structure

- `main.go`: Entry point of the application.
- `routes/`: Contains route handlers for different endpoints.
- `templates/`: Contains HTML templates for rendering pages.
- `models/`: Contains database models.
- `database/`: Contains database connection and configuration.

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.