# bigdata-server

## Prerequisite

To run the server, you need to install the following program:

- [The latest version of Go](https://go.dev)

Then, you need to download the dataset from [this link](https://drive.google.com/file/d/1v2NCEJlPb8HxebSqePDKwrQQY51-3Flc/view?usp=drive_link).

## Running

To begin the server, you simply enter

```bash
go run .
```

into the terminal, which is open inside the repo folder.

## Request

If you send `POST` requests with following body in `application/json` format:

```json
{
    "index": 0, // any integer number.
    "allergens": [
        "nuts", "dairy" // list of allergens to filter
    ]
}
```

the server will respond like below:

```json
[
    {
        "id": "22",
        "ingredients_str": ["Up or Sprite", "vegetable oil", "Kikkoman soy sauce", "garlic salt"],
        "allergens_str": ["soy"],
        "title": "Cuddy Farms Marinated Turkey",
        "allergens": "soy"
    }, 
    ..., // and 18 more items
    {
        "id": "474",
        "ingredients_str": ["yeast", "cup sugar", "cups warm water", "eggs", "oleo", "salt", "cups flour"],
        "allergens_str": ["gluten", "egg"],
        "title": "My Caramel Rolls",
        "allergens": "gluten,egg"
    }
]
```
