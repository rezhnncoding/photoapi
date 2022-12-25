# photoapi

GetPhotoList

localhost:8080/photo/getList

{

    "ResCode": "200",
    "ResMessage": "Success",
    "Data": [
        {
            "Id": "63a15d4d774e4e4ba6a3e0ce",
            "Title": "photo1",
            "Description": "in axe 1ast",
            "ImageName": "",
            "CreateDate": "0001-01-01T00:00:00Z",
            "CreatorUserId": "",
            "VisitCount": 0,
            "LikeCount": 0
        },
        {
            "Id": "63a81610f3510b4354c498bf",
            "Title": "baraye axe 4 hast",
            "Description": "in axe 4 teste",
            "ImageName": "71002909-f25e-4704-9414-e4debf5cd0a7.png",
            "CreateDate": "2022-12-25T09:21:19.391Z",
            "CreatorUserId": "",
            "VisitCount": 0,
            "LikeCount": 0
        }
    ]
}


GetPhotoBId

localhost:8080/photo/63a81610f3510b4354c498bf

{

    "ResCode": "200",
    "ResMessage": "Success",
    "Data": {
        "Id": "63a81610f3510b4354c498bf",
        "Title": "baraye axe 4 hast",
        "Description": "in axe 4 teste",
        "ImageName": "71002909-f25e-4704-9414-e4debf5cd0a7.png",
        "CreateDate": "2022-12-25T09:21:19.391Z",
        "CreatorUserId": "",
        "VisitCount": 0,
        "LikeCount": 0
    }
}


UploadPhoto

localhost:8080/photo/Create

{

    "ResCode": "200",
    "ResMessage": "Success",
    "Data": {
        "NewUserId": "63a81610f3510b4354c498bf"
    }
}


DeletePhoto

localhost:8080/photo/Delete/63a18655a785546e1b66c6d8

{

    "ResCode": "200",
    "ResMessage": "Success",
    "Data": null
}
