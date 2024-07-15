
<h1 style="color:dodgerblue; font-size:36px;">Api Documentation</h1>

<h2 style="color:dodgerblue;"> Overview</h2>

<h4 style="color:white;">This API provides endpoints to retrieve repository information and commits stored in the database. It also provides 
endpoints to set the 
credentials of the github repository to monitor and retrieve repositories by language, stars, and name.
Github API is used to retrieve the repository information and commits.</h3>

<h3 style="color:dodgerblue;"> Endpoints </h3>

<p style="color:green; font-style: normal;text-decoration: underline; font-weight: bold;"> GET localhost:8086/v1/set/repo/credential?owner={owner}
&repo={reponame}

Set the credentials of the github repository to monitor.

**Request Query:**

- `owner` (query parameter, required): A github username.
- `repo` (query parameter, required): A github repository name.
**Response:**
- `200 OK`: repo credentials set successfully.
<p style="color:green; font-style:normal; font-weight: bold">Sample Response </p>
  
```json
  {
    "message": "repo credentials set successfully",
    "status": 200
  }
  ```

<p style="color:green; font-style:normal; font-weight: bold; text-decoration: underline;"> GET localhost:8086/v1/repo/{name}

Retrieve a repo by name.

**Request Parameters:**

- `name` (path parameter, required): The Name of the repo.

**Response:**

- `200 OK`: repo retrieved successfully.
<p style="color:green; font-style:normal; font-weight: bold"> Sample Response </p>
  
```json
  {
     "message": "repo retrieved successfully",
    "data": {
        "id": 497859013,
        "name": "shop-arena",
        "created_at": "2022-05-30T08:42:30Z",
        "updated_at": "2022-05-30T08:54:23Z",
        "html_url": "https://github.com/Dilly3/shop-arena",
        "description": "API(Golang) multi vendor platform ",
        "language": "Go",
        "forks": 0,
        "stargazers_count": 0,
        "open_issues": 0
    },
    "status": 200
  } 
  ```
- `404 Not Found`: repo not found.
<p style="color:green; font-style:normal; font-weight: bold"> Sample Response </p>
  
```json
  {
    "message": "record not found",
    "status": 404
  }
  ```
<p style="color:green; font-style:normal; font-weight: bold;text-decoration: underline;">GET localhost:8086/v1/commits/{name}/{limit} </p>

Retrieve commits by repo name.

**Request Parameters:**

- `name` (path parameter, required): The name of the repo.
- `limit` (path parameter, required): The limit of the commits to retrieve.

**Response:**

- `200 OK`: repo retrieved successfully.
<p style="color:green; font-style:normal; font-weight: bold">Sample Response </p>
  
```json
  {
     "message": "commits retrieved successfully",
    "data": {
      "ID": "41791b4b3e8cb9678271b34309a024ba8870682c",
      "repo_name": "houdini",
      "Message": "Merge pull request #6 from Dilly3/storage-functions\n\nStorage functions",
      "AuthorName": "D'TechNiShan",
      "AuthorEmail": "",
      "Date": "2024-07-14T20:02:14Z",
      "URL": ""

    },
    "status": 200
  } 
  ```
<p style="color:green; font-style:normal; font-weight: bold;text-decoration: underline;">GET localhost:8086/v1/repos/{language}/{limit} </p>

Retrieve repo by language.

**Request Parameters:**

- `language` (path parameter, required): The programming language used for the repo.
- `limit` (path parameter, required): The limit of the repos to retrieve.

**Response:**

- `200 OK`: repos retrieved successfully.
<p style="color:green; font-style:normal; font-weight: bold"> Sample Response </p>
  
```json
  {
  "message": "repo retrieved successfully",
  "data": [
    {
      "id": 828183354,
      "name": "houdini",
      "created_at": "2024-07-13T11:14:02Z",
      "updated_at": "2024-07-14T12:13:04Z",
      "html_url": "https://github.com/Dilly3/houdini",
      "description": "Houdini is a repository for github apis",
      "language": "Go",
      "forks": 0,
      "stargazers_count": 0,
      "open_issues": 0
    }
  ],
  "status": 200
  } 
  ```
 <p style="color:green; font-style:normal; font-weight: bold; text-decoration: underline;"> GET localhost:8086/v1/repo-stars </p>

Retrieve repos by stars.

**Response:**
- `200 OK`: repos retrieved successfully.
<p style="color:green; font-style:normal; font-weight: bold"> Sample Response </p>
  
```json
{ 
  "message": "repos retrieved successfully",
"data": [
{
"id": 1,
"name": "grit",
"created_at": "",
"html_url": "https://github.com/mojombo/grit",
"description": "**Grit is no longer maintained. Check out libgit2/rugged.** Grit gives you object oriented read/write access to Git repositories via Ruby.",
"language": "",
"forks": 0,
"stargazers_count": 0,
"open_issues": 0
}
  ],
    "status": 200
  } 
  ```