# Action Logger

Action Logger is a simple project demonstrating a web server built using Go, which will integrate with Apache Kafka and Apache Cassandra to log actions.


## Installation

1. Make sure you have Go installed (version 1.16+ is recommended). If not, you can download it [here](https://golang.org/dl/).

2. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/action-logger.git

3. Navigate to the cloned directory:

   ```bash
   cd action-logger

4. Install the dependencies:
   
   ```bash
   go mod download

4. Running the server:
From the root directory of the project, run:

   ```bash
   go run ./cmd/action-logger