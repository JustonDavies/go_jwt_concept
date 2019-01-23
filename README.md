### Purpose:
  This project exists as a simple web application to demonstrate some core JWT proof of concepts.

  This project was also designed to have minimal software and hardware requirements, be platform agnostic, and self contained.

  It's not much more than a sandbox for working through concepts with junior resources.

### Dependencies:
  This project depends on Go, vgo and some small selection of supporting tools.

  - Go: It is recommended you install Go `1.11.0` via `goenv` Instructions for installing can be found [here](https://github.com/syndbg/goenv) or by running the following commands:
    ```
    git clone https://github.com/syndbg/goenv.git ~/.goenv

    echo `export GOENV_ROOT="$HOME/.goenv"
          export PATH="$GOENV_ROOT/bin:$PATH"
          export PATH="$GOENV_ROOT/shims:$PATH"
          eval "$(goenv init -)"
          export GOPATH="$HOME/workspace/go"` > ~/.bash_profile
    source ~/.bash_profile

    goenv install 1.10.3
    goenv global  1.10.3
    ```

  - vgo: Dependencies are managed via the Go 1.11 built-in dependency management tool vgo to download dependencies needed for building:
    ```
    go mod init
    go mod tidy
    go mod download
    ```

### Configuration:
  Minimal configuration is required for this project, just download the dependencies and build.

### Secrets / Infrastructure:
  This project assumes that the predicate supporting services have already been deployed.

  This project is merely a scratch project and does not employ good secret hygiene, you should not copy the paradigms here.

### Deploying:
  To compile (and run) this code and deploy you need only build the binary via the following commands:

  ```
  make
  bin/go_jwt_concept
  ```

### Output:
  After deploying a web application will be available for use at the stated URL provided by the log output.

#### Endpoints
This application has just one HTTPS endpoint

`GET /`
  - Parameters:
    - URL: This endpoint will not acknowledge URL encoded parameters
    - Body: This endpoint will not acknowledge and parameters included in the body
  - Exceptions:
    - Internal Error: If the endpoint hits a critical error while encoding the results, or reading user attributes it will return a status of 500
  - Return:
    - If no errors are encountered the endpoint will return a JSON-encoded `Hello, world` and a status of 200

`GET /restricted`
  - Parameters:
    - URL: This endpoint will not acknowledge URL encoded parameters
    - Body: This endpoint will not acknowledge and parameters included in the body
  - Exceptions:
    - Unauthorized: If the endpoint is called without a proper JWT token as an authentication header it will return a status of 401
    - Internal Error: If the endpoint hits a critical error while encoding the results, or reading user attributes it will return a status of 500
  - Return:
    - If no errors are encountered the endpoint will return a JSON-encoded `Hello, world` and a status of 200

`POST /auth/login`
  - Parameters:
    - URL: This endpoint will not acknowledge URL encoded parameters
    - Body: This endpoint expects a JSON-encoded struct with the `email` and `password` to be used for creating a new JWT token
  - Exceptions:
    - Unauthorized: If the user could not be found it will return a status of 401
    - Internal Error: If the endpoint hits a critical error while encoding the results, or reading user attributes it will return a status of 500
  - Return:
    - If no errors are encountered the endpoint will return a new JWT token with a status of 200

`POST /auth/renew`
  - Parameters:
    - URL: This endpoint will not acknowledge URL encoded parameters
    - Body: This endpoint will not acknowledge and parameters included in the body
  - Exceptions:
    - Unauthorized: f the endpoint is called without a proper JWT token as an authentication header it will return a status of 401
    - Internal Error: If the endpoint hits a critical error while encoding the results, or reading user attributes it will return a status of 500
  - Return:
    - If no errors are encountered the endpoint will return renewed JWT token with a status of 200