# Votio

Votio is a voting application designed to facilitate the creation and management of polls. It allows users to create polls, vote on them, and view results in real-time.

## Features

- **Create Polls**: Easily create polls with multiple options.
- **Vote**: Users can vote on their preferred options.
- **Real-Time Results**: View poll results as votes are cast.

## Project Structure
```
votio/
├── service/                     # Backend (Golang, Echo)
│   ├── cmd/                        # Main application entry point
│   ├── internal/                   # Internal packages and handlers
│   ├── Dockerfile                  # Dockerfile for backend
│   ├── go.mod                      # Go module definition
│   └── go.sum                      # Go dependencies
├── client/                         # Frontend (Nuxt, TypeScript)
│   ├── public/                     # Static assets (CSS, images)
│   ├── server/                     # Vue Server
│   ├── pages/                      # Nuxt pages
│   ├── Dockerfile                  # Dockerfile for frontend
│   ├── package.json                # Node.js dependencies for frontend
│   ├── tsconfig.json               # TypeScript configuration
│   └── nuxt.config.ts              # Nuxt configuration
├── .tool-versions                  # asdf-managed tool versions (Node.js, Go)
├── .air.toml                       # Air config for live reload during 
├── docker-compose.yml              # Docker Compose to manage multiple services
└── Makefile                        # Automation commands for building/running
```

## Installation

To install and run Votio locally, follow these steps:

1. **Clone the repository**:
  ```sh
  git clone https://github.com/ardipta/votio.git
  ```
2. **Navigate to the project directory**:
  ```sh
  cd votio
  ```
3. **Install dependencies**:
  ```sh
  go mod tidy
  ```
4. **Run the application**:
  ```sh
  go run main.go
  ```

## Usage

1. **Create an Account**: Sign up for a new account or log in if you already have one.
2. **Create a Poll**: Navigate to the poll creation page and enter your poll details.
3. **Vote**: Select your preferred option and submit your vote.
4. **View Results**: Check the results page to see the current standings.

## Contributing

We welcome contributions! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any questions or feedback, please contact us at [ashiqur49.rahaman@gmail.com](mailto:ashiqur49.rahaman@gmail.com).
