# API Documentation

## Description

API for CRUD operations on users.

## Endpoints

### Create User

- **Endpoint:** `/createUser`
- **Method:** POST
- **Description:** Create a new user with the provided user information
- **Consumes:** application/json
- **Produces:** application/json
- **Tags:** Users
- **Summary:** Create a new user

#### Parameters

- **userRequest:** User information for registration
  - Type: Object
  - Required: true
  - Schema: [UserRequest](#userrequest)

#### Responses

- 200: OK
  - Description: User created successfully
  - Schema: [UserResponse](#userresponse)
- 400: Bad Request
  - Description: Invalid user information
  - Schema: [RestErr](#resterr)
- 500: Internal Server Error
  - Description: Error processing the request
  - Schema: [RestErr](#resterr)

### Delete User

- **Endpoint:** `/deleteUser/{userId}`
- **Method:** DELETE
- **Description:** Deletes a user based on the ID provided as a parameter.
- **Consumes:** application/json
- **Produces:** application/json
- **Tags:** Users
- **Summary:** Delete User

#### Parameters

- **userId:** ID of the user to be deleted
  - Type: String
  - In: Path
  - Required: true
- **Authorization:** Insert your access token
  - Type: String
  - In: Header
  - Default: Bearer <Add access token here>
  - Required: true

#### Responses

- 200: OK
  - Description: User deleted successfully
- 400: Bad Request
  - Description: Invalid user ID
  - Schema: [RestErr](#resterr)
- 500: Internal Server Error
  - Description: Error processing the request
  - Schema: [RestErr](#resterr)

### Get User by Email

- **Endpoint:** `/getUserByEmail/{userEmail}`
- **Method:** GET
- **Description:** Retrieves user details based on the email provided as a parameter.
- **Consumes:** application/json
- **Produces:** application/json
- **Tags:** Users
- **Summary:** Find User by Email

#### Parameters

- **userEmail:** Email of the user to be retrieved
  - Type: String
  - In: Path
  - Required: true
- **Authorization:** Insert your access token
  - Type: String
  - In: Header
  - Default: Bearer <Add access token here>
  - Required: true

#### Responses

- 200: User information retrieved successfully
  - Description: User details
  - Schema: [UserResponse](#userresponse)
- 400: Error: Invalid user ID
  - Description: Invalid user ID
  - Schema: [RestErr](#resterr)
- 404: User not found
  - Description: User not found
  - Schema: [RestErr](#resterr)

### Get User by ID

- **Endpoint:** `/getUserById/{userId}`
- **Method:** GET
- **Description:** Retrieves user details based on the user ID provided as a parameter.
- **Consumes:** application/json
- **Produces:** application/json
- **Tags:** Users
- **Summary:** Find User by ID

#### Parameters

- **userId:** ID of the user to be retrieved
  - Type: String
  - In: Path
  - Required: true
- **Authorization:** Insert your access token
  - Type: String
  - In: Header
  - Default: Bearer <Add access token here>
  - Required: true

#### Responses

- 200: User information retrieved successfully
  - Description: User details
  - Schema: [UserResponse](#userresponse)
- 400: Error: Invalid user ID
  - Description: Invalid user ID
  - Schema: [RestErr](#resterr)
- 404: User not found
  - Description: User not found
  - Schema: [RestErr](#resterr)

### User Login

- **Endpoint:** `/login`
- **Method:** POST
- **Description:** Allows a user to log in and receive an authentication token.
- **Consumes:** application/json
- **Produces:** application/json
- **Tags:** Authentication
- **Summary:** User Login

#### Parameters

- **userLogin:** User login credentials
  - Type: Object
  - Required: true
  - Schema: [UserLogin](#userlogin)

#### Responses

- 200: Login successful, authentication token provided
  - Description: Authentication token
  - Headers:
    - Authorization: String
  - Schema: [UserResponse](#userresponse)
- 403: Error: Invalid login credentials
  - Description: Invalid login credentials
  - Schema: [RestErr](#resterr)

### Update User

- **Endpoint:** `/updateUser/{userId}`
- **Method:** PUT
- **Description:** Updates user details based on the ID provided as a parameter.
- **Consumes:** application/json
- **Produces:** application/json
- **Tags:** Users
- **Summary:** Update User

#### Parameters

- **userId:** ID of the user to be updated
  - Type: String
  - In: Path
  - Required: true
- **userRequest:** User information for update
  - Type: Object
  - Required: true
  - Schema: [UserUpdateRequest](#userupdaterequest)
- **Authorization:** Insert your access token
  - Type: String
  - In: Header
  - Default: Bearer <Add access token here>
  - Required: true

#### Responses

- 200: OK
  - Description: User details updated successfully
- 400: Bad Request
  - Description: Invalid user information
  - Schema: [RestErr](#resterr)
- 500: Internal Server Error
  - Description: Error processing the request
  - Schema: [RestErr](#resterr)

## Definitions

### UserLogin

Structure containing the necessary fields for user login.

- **Description:** User's email and password
- **Type:** Object
- **Properties:**
  - email (string): User's email (required and must be a valid email address).
  - password (string): User's password (required, minimum of 6 characters, and must contain at least one of the characters: !@#$%*).

### UserRequest

Structure containing the required fields for creating a new user.

- **Description:** User's age, email, name, and password
- **Type:** Object
- **Properties:**
  - age (integer): User's age (required, must be between 1 and 140).
  - email (string): User's email (required and must be a valid email address).
  - name (string): User's name (required, minimum of 4 characters, maximum of 100 characters).
  - password (string): User's password (required, minimum of 6 characters, and must contain at least one of the characters: !@#$%*).

### UserUpdateRequest

Structure for updating user details.

- **Type:** Object
- **Properties:**
  - age (integer): User's age (required, must be between 1 and 140).
  - name (string): User's name (required, minimum of 4 characters, maximum of 100 characters).

### UserResponse

Structure containing user details.

- **Type:** Object
- **Properties:**
  - age (integer): User's age.
  - email (string): User's email.
  - id (string): User's ID.
  - name (string): User's name.

### RestErr

Structure for describing why the error occurred.

- **Type:** Object
- **Properties:**
  - causes (array): Error causes.
    - field (string): Field associated with the error cause.
    - message (string): Error message describing the cause.
  - code (integer): Error code.
  - error (string): Error description.
  - message (string): Error message.
