## HOUDINI v1.0.0
Houdini is a simple REST API that retrieves information about repositories and commits from GitHub.
### Env Variables
<p> The project uses the following environment variables </p>

```Go,
 
export HOUDINI_PORT=:8086
# set cron interval value in minutes (60 = 1 hour)
export HOUDINI_CRON_INTERVAL=4
export HOUDINI_NETWORK_RETRY=3

export HOUDINI_POSTGRES_HOST=db
export HOUDINI_POSTGRES_PASSWORD=docker
export HOUDINI_POSTGRES_USER=docker
export HOUDINI_POSTGRES_PORT=5432
export HOUDINI_POSTGRES_DB=houdini
export HOUDINI_POSTGRES_TIMEZONE=Africa/Lagos

export HOUDINI_GITHUB_BASE_URL=https://api.github.com/
export HOUDINI_GITHUB_TOKEN={your_github_token}
export HOUDINI_GITHUB_OWNER=chromium
export HOUDINI_GITHUB_REPO=chromium
export HOUDINI_GITHUB_PER_PAGE=60
export HOUDINI_GITHUB_SINCE=2023-07-01T00:00:00Z

export HOUDINI_REDIS_PASSWORD=redis123
export HOUDINI_REDIS_User=user123
export HOUDINI_REDIS_ADDR=redis:6379
export HOUDINI_REDIS_PORT=6379
```
<p> the env file in the root directory of the project houses the necessary environment variables for the project to run successfully. </p> 
<p> *** N/B Add your github token to the HOUDINI_GITHUB_TOKEN variable in the .env file </p>
<p> The HOUDINI_CRON_INTERVAL variable is the delay time in minutes for the cron job to run. </p>


## Installation
- <p style="color: red; font-weight: bold;"> To install the project, you need docker running on ur machine </p>
- <p> Clone the project from the repository <a href="https://github.com/Dilly3/houdini">Github</a> </p>
- <p> Run the command `go mod tidy` </p>
- <p style="color: red; font-weight: bold;"> *** Set up ur environment variables (github token). Add the Github token before you run the app </p>
- <p> How to create gitHub token <a href="https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#creating-a-personal-access-token-classic">
- Create Token </a> </p>
- <p> After you have created your token, insert it in the env file</p>
- <p> Run the command `make up-build` to build the project initially </p>
- <p> Run the command `make up` to start the project subsequently </p>

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
