# Task Description: Build a Task Management System

## Overview
In this assignment, you will develop a simple task management system using Go and PostgreSQL. Your system will allow users to add, update, delete, and list tasks. Additionally, you will implement transactions to ensure data integrity during updates.

## Objectives
- Understand and implement basic CRUD operations in a PostgreSQL database.
- Use transactions to manage database operations safely.
- Gain experience with SQL database connectivity and operations in Go.

## Requirements
- **Go Programming**: You should be familiar with the basics of programming in Go.
- **PostgreSQL Database**: You will need PostgreSQL installed on your machine or have access to a PostgreSQL server.

## Database Setup
Create a PostgreSQL database and a table for storing tasks:
```sql
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE
);
```

## Tasks

### Create a Task
- **Functionality**: Implement a function to insert a new task into the database.
- **Uniqueness**: Ensure that task names are unique to avoid duplicates in the system.

### Read Tasks
- **Functionality**: Implement a function to fetch and display all tasks from the database.
- **Display**: Ensure all tasks are retrieved and properly formatted for easy readability.

### Update a Task
- **Functionality**: Implement a function to mark a task as completed.
- **Transactions**: Use a transaction to ensure the operation is atomic. This means if the operation fails at any step, it should not commit any changes to the database.

### Delete a Task
- **Functionality**: Implement a function to remove a task based on its ID.
- **Transactions**: Use a transaction for this operation as well, ensuring all steps must complete successfully or none at all. If one step fails, no part of the operation should affect the database.

## Use the Query Function with a Context Timeout

- **Setup**: In your main function, create a context with a timeout shorter than the query's duration to simulate a timeout scenario.
- **Execution**: Call the `QueryDatabase` function with this context.
- **Error Handling**: Handle and print the error if the query is cancelled or times out.

## Testing Your Implementation

- **Test Setup**: Write a test in a separate file to verify that your `QueryDatabase` function behaves as expected.
- **Test Scenarios**:
  - The query completes successfully before the context times out.
  - The query is cancelled due to a context timeout.
- **Testing Methodology**: Use the `testing` package to implement your test cases and verify the correct behavior of your function.

