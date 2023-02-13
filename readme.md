# Catalog API

## About this Project

This API catalog was built using Golang with the gin framework and also gorm for ORM on Golang.

If you want to run this API make sure you install the Go programming language

After that you create a database first so that GORM can carry out its work (make sure the database name is the same as the one in the database directory)

You can run this program by running

```sh
go run main.go
```

## How to use ?

### Register

On Postman, you can register users by sending the following POST Request

```json
// # POST::localhost:3000/user/register
{
  "email": "your-email@email.com",
  "full_name": "your full name",
  "password": "your password"
}
```

If successful, the JSON will respond as follows

```json
{
  "id": 1,
  "email": "your-email@email.com",
  "full_name": "your full name"
}
```

### Login

In the login function, you need to send a Request like the following

```json
// # POST::localhost:3000/user/login
{
  "email": "your-email@email.com",
  "password": "your password"
}
```

After that the API will respond by sending an Authentication token as follows

```json
{
  "token": "[AuthToken]"
}
```

### Create the Product

To create a product, you need to send a Request like the following (Keep in mind, you need to embed the Auth Token in the Authorization Bearer Token to make it work)

```json
// # POST::localhost:3000/products/
{
  "product_name": "your product name",
  "description": "your product description",
  "image": "product.img"
}
```

If successful, the API will provide a response

```json
{
  "id": 1,
  "product_name": "your product name",
  "description": "your product description",
  "image": "product.jpg",
  "created_at": "[datetime]"
}
```

### Get Products

You can display all data products in the database by doing a GET Request at the following URL

```sh
localhost:3000/products/
```

The API will respond to JSON as follows

```json
{
  "code": 200,
  "result": [
    {
      "id": 1,
      "product_name": "your product name",
      "description": "your product description",
      "image": "product.jpg",
      "created_at": "[datetime]",
      "updated_at": "[datetime]"
    },
    {
      "id": 2,
      "product_name": "your product name",
      "description": "your product description",
      "image": "product.jpg",
      "created_at": "[datetime]",
      "updated_at": "[datetime]"
    }
  ]
}
```

### Get a Product

You can display one of the products you want by requesting the following URL

```
localhost:3000/products/[id]
```

Then what will be the API response is

```json
{
  "code": 200,
  "result": {
    "id": 1,
    "product_name": "your product name",
    "description": "your product description",
    "image": "product.jpg",
    "created_at": "[datetime]",
    "updated_at": "[datetime]"
  }
}
```

### Edit a Product

If you want to update a product, you can make a request like the following

```json
// # PUT::localhost:3000/products/[id]
{
  "product_name": "your product name",
  "description": "your product description",
  "image": "product.jpg"
}
```

If successful, the API will respond

```json
{
  "id": 1,
  "product_name": "Asus Zenfone 7",
  "description": "your product description",
  "image": "asus-zenfone-7.jpg",
  "updated_at": "2023-02-13T10:40:54.5392119+07:00"
}
```

### Delete a Product

If you want to delete a product, you can send a request like this using the DELETE method

```
localhost:3000/products/[id]
```

And if successful then the API will respond

```
{
    "message": "Product Successfully Deleted"
}
```
