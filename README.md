# Clear Badge Cache Web Interface 🧹✨

[![Go Version](https://img.shields.io/github/go-mod/go-version/judahpaul16/clear-badge-cache)](https://go.dev/dl/)
[![Go Report Card](https://goreportcard.com/badge/github.com/judahpaul16/www-clear-badge-cache)](https://goreportcard.com/report/github.com/judahpaul16/www-clear-badge-cache)
[![GitHub](https://img.shields.io/github/license/judahpaul16/clear-badge-cache)](LICENSE)
[![Website](https://img.shields.io/badge/website-https://clear--badge--cache.com-blue)](https://clear-badge-cache.com/)

[Clear-Badge-Cache.com](https://clear-badge-cache.com/) is a web-based tool designed to make it easier to clear the cache for GitHub badge images by providing a simple and intuitive user interface. This tool is based on the Clear Badge Cache CLI tool but adds a user-friendly web interface for the same functionality.

## Features 🌟

- **Simple User Interface**: Just enter the URL of the badge you need to clear, and the cache will be cleared without needing to handle CLI operations.
- **Built with Go, Templ, and HTMX**: Leveraging modern technologies for a lightweight and responsive experience.

## Quick Start 🚀

### Prerequisites 📋

- Go (1.22.2 or later)

### Installation 🛠

1. Clone the repository:

   ```bash
   git clone https://github.com/judahpaul16/www-clear-badge-cache.git
   ```
   
2. Navigate to the project directory:

   ```bash
   cd www-clear-badge-cache
   ```

3. Run the server:

   ```bash
   go run main.go
   ```

   This will start the web server on `localhost:8080`.

### Usage 🖥️

- Open a web browser and go to `http://localhost:8080`.
- You will see a form to input the URL of the badge you wish to clear from the cache.
- Enter the URL and submit the form to clear the cache.

### Building 🔨

To build an executable, run:

```bash
go build -o www-clear-badge-cache
```

### Running 🏃

After building, you can run the application directly:

```bash
./www-clear-badge-cache
```

## Contributing 🤝

Contributions are welcome! Feel free to open pull requests or issues to improve the functionality or documentation.

## License 📝

Distributed under the GNU GPL-3 License. See `LICENSE` for more information.
