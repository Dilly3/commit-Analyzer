## HOUDINI v1.0.0
Houdini is a simple REST API that retrieves information about repositories and commits from GitHub.
### Env Variables
<p> The project uses the following environment variables </p>

```Go,
 
export HOUDINI_PORT=:8086

export HOUDINI_CRON_INTERVAL=10

export HOUDINI_POSTGRES_HOST=db
export HOUDINI_POSTGRES_PASSWORD=docker
export HOUDINI_POSTGRES_USER=docker
export HOUDINI_POSTGRES_PORT=5432
export HOUDINI_POSTGRES_DB=houdini
export HOUDINI_GITHUB_TOKEN=
export HOUDINI_GITHUB_OWNER=
export HOUDINI_GITHUB_REPO=
export HOUDINI_POSTGRES_TIMEZONE=
export HOUDINI_GITHUB_BASE_URL=https://api.github.com/
```
<p> the env file in the root directory of the project houses the necessary environment variables for the project to run successfully. </p> 
<p> Add your github token to the HOUDINI_GITHUB_TOKEN variable in the .env file </p>
<p> The HOUDINI_CRON_INTERVAL variable is the delay time in minutes for the cron job to run. </p>


## Installation
- <p> To install the project, you need docker running on ur machine </p>
- <p> Clone the project from the repository </p>
- <p> Run the command `go mod tidy` </p>
- <p style="color: yellow; font-weight: bold;"> *** Set up ur environment variables (github owner, github repo , github token) </p>
- <p> Run the command `make up` to start the project </p>

## Test
<p> To run the test, run the command `make test` </p>
```Go,
 Features
- Retrieve repositories by language
- Retrieve commits by repository name
- Retrieve commits by repository name and limit
- Retrieve repositories by language and limit
- Retrieve commits by repository name and limit
- 
```
<p style="color: yellow; font-weight: bold;"> *** You can read more on the API documentation in the docs folder in the root directory </p>
### Author
<p> Name: Anikammadu Michael  </p>
<p> Email: michael.anikamadu@gmail.com </p>
<p> GitHub: dilly3</p>
