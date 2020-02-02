
### Server https://perfectapi.extensionsoft.biz

## Authentication api 

### GetAllItem

    Endpoint : "/product/v1/item"
    Content-Type : application/json 

request :
``` json 
{
	"limit": 5
}
```

response 
``` json 
{
    "response": "success",
    "data": [
        {
            "item_id": 1,
            "item_name": "item1",
            "image": "",
            "degree": 5,
            "price": 155,
            "qty": 5
        },
        {
            "item_id": 2,
            "item_name": "item2",
            "image": "",
            "degree": 2,
            "price": 155,
            "qty": 5
        },
        {
            "item_id": 3,
            "item_name": "item3",
            "image": "",
            "degree": 4,
            "price": 155,
            "qty": 5
        },
        {
            "item_id": 4,
            "item_name": "item4",
            "image": "",
            "degree": 4,
            "price": 200,
            "qty": 2
        }
    ]
}
``` 

## GetItemByID
    Endpoint : "/product/v1/item/id"
    Content-Type : application/json 
    
``` json 
{
	"id": 1
}
```

response 
``` json 
{
    "response": "success",
    "data": {
        "item_id": 1,
        "item_name": "item1",
        "image": "",
        "degree": 5,
        "price": 155,
        "qty": 5
    }
}
```

## GetCategory
    Endpoint : "/product/v1/category"
    Content-Type : application/json 

response 
``` json 
{
    "response": "success",
    "data": [
        {
            "id": 1,
            "category_name": "category1",
            "image": "url"
        },
        {
            "id": 2,
            "category_name": "category2",
            "image": "url"
        },
        {
            "id": 3,
            "category_name": "category3",
            "image": "url"
        }
    ]
}
```

## GetCategoryById
    Endpoint : "/product/v1/category/id"
    Content-Type : application/json 
    
``` json 
{
	"id": 1
}
```

response 
``` json 
{
    "response": "success",
    "data": [
        {
            "item_id": 1,
            "item_name": "item1",
            "image": "",
            "degree": 5,
            "price": 155,
            "qty": 5
        },
        {
            "item_id": 2,
            "item_name": "item2",
            "image": "",
            "degree": 2,
            "price": 155,
            "qty": 5
        },
        {
            "item_id": 3,
            "item_name": "item3",
            "image": "",
            "degree": 4,
            "price": 155,
            "qty": 5
        },
        {
            "item_id": 4,
            "item_name": "item4",
            "image": "",
            "degree": 4,
            "price": 200,
            "qty": 2
        }
    ]
}
```