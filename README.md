# Smerch
A high-performance Python package manager written in Go.

Inspired by [Poetry](https://python-poetry.org/) & [Pipenv](https://pipenv.pypa.io/en/latest/).

### Project status
At this stage this is just an experiment to see if I can write a package manager in Go that is considerably more performant than existing solutions.

### Why does this project exist?
Go is one of the best languages for networking, making it ideal for handling the large number of client and HTTP requests a package manager generates. It also works seamlessly in Docker containers and cloud environments, which are common deployment scenarios for package management tools.

Additionally, there are a growing number of Python developers who have picked up Go. Building this package manager in Go allows these developers to read and understand the source code, contribute easily, and benefit from Go’s performance and concurrency features.

### Build & Run
```bash
make run
```

### Contributions
Please open up an issue first with your new feature request before submitting any code.