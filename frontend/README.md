# React Golang CRUD Frontend

This is a React 18 + Vite + TailwindCSS frontend project designed to interact with a Golang REST API for CRUD (Create, Read, Update, Delete) operations on products.

## Features

- **Product Listing:** View all products.
- **Add Product:** (Admin only) Add new products.
- **Edit Product:** (Admin only) Update existing products.
- **Delete Product:** (Admin only) Delete products.
- **Role-Based Behavior:** Simulates 'user' and 'admin' roles with different access levels.
- **TailwindCSS:** Modern and responsive UI styling.
- **React Router:** Seamless navigation between pages.
- **Alerts:** Simple success/error messages.

## API Endpoints

The frontend expects a Golang REST API running at `http://localhost:8080/products` with the following endpoints:

- `GET /products`: Get all products.
- `POST /products`: Add a new product.
- `PUT /products/:id`: Update an existing product.
- `DELETE /products/:id`: Delete a product.

## Setup Instructions

1.  **Navigate to the project directory:**

    ```bash
    cd react-golang-crud-frontend
    ```

2.  **Install dependencies:**

    ```bash
    npm install
    ```

3.  **Run the development server:**

    ```bash
    npm run dev
    ```

    The application will typically be available at `http://localhost:5173`.

## Role Switching

You can switch between 'User' and 'Admin' roles using the button in the navigation bar to test the role-based access control.

- **User Role:** Can only view the product list.
- **Admin Role:** Can view, add, edit, and delete products.
