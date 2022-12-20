# photoapi

Get Photo List

localhost:7070\photo\getList
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
            "Id": "63a18655a785546e1b66c6d8",
            "Title": "image edited",
            "Description": "axe 2 edit shde",
            "ImageName": "be77c3d5-9362-4d3a-b116-7498df1cfa16.jfif",
            "CreateDate": "2022-12-20T10:11:57.288Z",
            "CreatorUserId": "",
            "VisitCount": 0,
            "LikeCount": 0
        },
        {
            "Id": "63a18a78b9a93c1c33d15880",
            "Title": "image edited",
            "Description": "axe 3 edit shde",
            "ImageName": "82926914-5412-4341-ac30-b18f28aec500.jfif",
            "CreateDate": "2022-12-20T10:18:23.933Z",
            "CreatorUserId": "",
            "VisitCount": 0,
            "LikeCount": 0
        }
    ]
}

Post photo

localhost:7070\photo\Create

{
    "Title":"image",
    "Description":"axe 2",
    "ImageName":"file"
}

edit photo and desc

localhost:7070\photo/Edit/id:

Title:image edited
Description:axe 3 edit shde
ImageName:fil333

delete photo 

localhost:7070\photo/Delete/id:

{
    "ResCode": "200",
    "ResMessage": "Success",
    "Data": null
}

