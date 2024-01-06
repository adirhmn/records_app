# How to Running Records APP

Here are the steps to clone a project from GitHub to your computer.

## Prerequisites

Before you begin, make sure you have Git installed on your computer. If not, you can download and install it from the [official Git website](https://git-scm.com/).

## Clone Project Github
1. **Copy Repository URL**

    On the repository page, look for the `Code` or `Clone` button located at the top right. Click on the button and copy the displayed URL (usually in HTTPS or SSH format).

2. **Open Terminal**

    Open the terminal or command prompt on your computer.

3. **Navigate to Destination Directory**

    Use the `cd` command to navigate to the directory where you want to store the project. For example:

    ```bash
    cd path/to/destination/directory
    ```

4. **Clone Repository**

    Type the following command to clone the repository:

    ```bash
    git clone [Repository URL]
    ```

    Replace `[Repository URL]` with the URL

    ```bash
    git clone https://github.com/adirhmn/records_app.git
    ```

5. **Done!**

    The project has now been successfully cloned to your computer. You can start working or exploring the code of the project.

## Run Application Using Docker Compose
6. **Navigate to Project Directory**

    After cloning the repository, navigate to the project directory:

    ```bash
    cd repository-directory
    ```

7. **Start Docker Compose**

    Run the following command to start the application using Docker Compose:

    ```bash
    docker-compose up --build
    ```

    This command will build the Docker images and start the containers in detached mode.
## Create a database with pgadmin
9. **Open pgAdmin**

    Launch pgAdmin on your computer with access   ```http://localhost:8888/```

10. **Log In**

    - Enter the username and password to log in to pgAdmin.
        - **email**: `admin@example.com`
        - **password**: `admin` 

11. **Add Server**

    - In the pgAdmin dashboard, right-click on the `Servers` option and select `Create` > `Server`.
    - In the `General` tab, enter a name for the server ```mezinkserver```.
    - In the `Connection` tab, enter the following details:
        - **Hostname/address**: `db`
        - **Port**: `5432` 
        - **Maintenance database**: `postgres`
        - **Username**: `postgres`
        - **Password**: `admin` (or the password specified in your Docker Compose configuration)
    - Click `Save` to add the server.

12. **Open Query Tool in pgAdmin**

    - In pgAdmin, right-click on the `mezinkserver` server that you added and select `Query Tool`.
    - A new query editor window will open.

13. **Create Database**

    In the query editor window, enter the following SQL command to create a new database named `myappdb`:

    ```sql
    CREATE DATABASE myappdb
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;
    ```

14. **Execute SQL Command**

    Click on the `Execute` button (usually represented by a green arrow or similar icon) to execute the SQL command and create the database.

15. **Verify Database Creation**

    You should see a confirmation message indicating that the database `myappdb` has been created successfully.

## Running Application
16. **Restart the application**
    Terminate the terminal `ctrl+c` and run the following command to start the application using Docker Compose:

    ```bash
    docker-compose up --build
    ```
### Test the Application Using Postman

17. **Open Postman**

    Launch Postman on your computer.

18. **Create a New Request**

    - Click on the `+` tab to create a new request.
    - Select `POST` as the HTTP method.

19. **Enter Request URL**

    Enter the URL endpoint of the application. 

    ```
    http://localhost:8080/api/records
    ```


20. **Set Request Body**

    - Select the `Body` tab.
    - Choose `raw` and `JSON (application/json)` as the format.
    - Enter the following request body:

    ```json
    {
      "startDate": "2024-01-06",
      "endDate": "2024-01-31",
      "minCount": 100,
      "maxCount": 200
    }
    ```
21. **Set Headers**

    - Select the `Headers` tab.
    - Click on `Add Header` and enter the following details:
        - **Key**: `x-api-key`
        - **Value**: `RAHASIA`
        
22. **Send Request**

    Click on the `Send` button to send the POST request to the application.
    
23. **Get Response**

    ```json
    {
    "code": 0,
    "status": "Success",
    "data": [
        {
            "id": 1,
            "createdAt": "2024-01-06T01:52:18.930962Z",
            "totalMarks": 200
        },
        {
            "id": 2,
            "createdAt": "2024-01-06T01:52:18.930962Z",
            "totalMarks": 200
        },
        {
            "id": 3,
            "createdAt": "2024-01-06T01:52:18.930962Z",
            "totalMarks": 195
        },
        {
            "id": 4,
            "createdAt": "2024-01-06T01:52:18.930962Z",
            "totalMarks": 190
        },
        {
            "id": 5,
            "createdAt": "2024-01-06T01:52:18.930962Z",
            "totalMarks": 190
        },
        {
            "id": 6,
            "createdAt": "2024-01-06T01:52:18.930962Z",
            "totalMarks": 190
        },
        {
            "id": 7,
            "createdAt": "2024-01-06T01:52:18.930962Z",
            "totalMarks": 190
        },
        {
            "id": 8,
            "createdAt": "2024-01-06T01:52:18.930962Z",
            "totalMarks": 190
        },
        {
            "id": 9,
            "createdAt": "2024-01-06T01:52:18.930962Z",
            "totalMarks": 190
        },
        {
            "id": 10,
            "createdAt": "2024-01-06T01:52:18.930962Z",
            "totalMarks": 195
        },
        {
            "id": 11,
            "createdAt": "2024-01-06T01:52:18.930962Z",
            "totalMarks": 200
        },
        {
            "id": 16,
            "createdAt": "2024-01-06T01:52:18.930962Z",
            "totalMarks": 200
        },
        {
            "id": 21,
            "createdAt": "2024-01-06T01:52:18.930962Z",
            "totalMarks": 200
        },
        {
            "id": 22,
            "createdAt": "2024-01-06T01:52:18.930962Z",
            "totalMarks": 200
        },
        {
            "id": 23,
            "createdAt": "2024-01-06T01:52:18.930962Z",
            "totalMarks": 200
        },
        {
            "id": 24,
            "createdAt": "2024-01-06T01:52:18.930962Z",
            "totalMarks": 200
        },
        {
            "id": 25,
            "createdAt": "2024-01-06T01:52:18.930962Z",
            "totalMarks": 200
        },
        {
            "id": 26,
            "createdAt": "2024-01-06T01:52:18.930962Z",
            "totalMarks": 200
        },
        {
            "id": 27,
            "createdAt": "2024-01-06T01:52:18.930962Z",
            "totalMarks": 200
        },
        {
            "id": 28,
            "createdAt": "2024-01-06T01:52:18.930962Z",
            "totalMarks": 200
        }
    ]
    ```
    ### Test the Application Using curl

24. **Open Terminal or Command Prompt**

    Open a terminal or command prompt on your computer.

25. **Run curl Command**

    Enter the following `curl` command to send a POST request to the application:

    ```bash
    curl --location 'http://localhost:8080/api/records' \
    --header 'x-api-key: RAHASIA' \
    --header 'Content-Type: application/json' \
    --data '{
      "startDate": "2024-01-06",
      "endDate": "2024-01-31",
      "minCount": 100,
      "maxCount": 200
    }'
    ```



