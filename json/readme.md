**Có hai loại dữ liệu bạn sẽ gặp khi làm việc với JSON:**

`Structured data (dữ liệu có cấu trúc)`

`Unstructured data (dữ liệu không có cấu trúc)`

-  Structured data : 

Đây là loại dữ liệu mà bạn biết trước cấu trúc. 

``` 
   {
     "species": "pigeon",
     "decription": "likes to perch on rocks"
   }
    
```
Để xử lý thằng này thì chỉ cần tạo 1 struct y như đoạn json 
```
    type Bird struct {
      Species string
      Description string
    }
```
```
    birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
    var bird Bird	
    json.Unmarshal([]byte(birdJson), &bird)
    fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)
```
https://play.golang.org/p/DtA6sEppLO

``
Theo quy ước, Go sử dụng tên của thuộc tính để biểu diễn thuộc tính của JSON cùng tên. Do đó thuộc tính Species trong struct Bird sẽ được ánh xạ tới species, Species hoặc sPeCiEs của chuỗi JSON.
``
Nếu là array thì làm như thế nào? 
```
    [
      {
        "species": "pigeon",
        "decription": "likes to perch on rocks"
      },
      {
        "species":"eagle",
        "description":"bird of prey"
      }
    ]
```
```
    birdJson := `[{"species":"pigeon","decription":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
    var birds []Bird
    json.Unmarshal([]byte(birdJson), &birds)
    fmt.Printf("Birds : %+v", birds)
```
- Embedded objects:
Thế giờ trong trường hợp bạn có một thuộc tính Dimensions chứa độ cao và độ dài 
```
    {
      "species": "pigeon",
      "decription": "likes to perch on rocks"
      "dimensions": {
        "height": 24,
        "width": 10
      }
    }
```

Giống như ví dụ trước, chúng ta cần tạo một struct ý như cấu trúc của chuỗi JSON. Tạo thêm struct Simensions và nhúng vào.

```
    type Dimensions struct {
      Height int
      Width int
    }
``` 

 
Rồi, giờ struct Bird sẽ như sau:
```
    type Bird struct {
        Species string
        Description string
        Dimensions Dimensions
    }
```

```
    birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`
    var birds Bird
    json.Unmarshal([]byte(birdJson), &birds)
    fmt.Printf(bird)
    // {pigeon likes to perch on rocks {24 10}}
```

https://play.golang.org/p/zOUMUNH4w9

- Custom attribute names:

``
Go dùng quy ước để ánh xạ dữ liệu từ chuỗi JSON với struct. Nhưng trong nhiều trường hợp cần đặt lại tên thuộc tính ở struct khác với chuỗi JSON thì làm như nào?
``
example: 
```
    {
      "birdType": "pigeon",
      "what it does": "likes to perch on rocks"
    }
```
Như chuỗi JSON này, mình thích dùng birdType cho Species hơn.

Và không thể tạo được thuộc tính của struct để parse cái thể loại này “what it does”.

```
type Bird struct {
  Species string `json:"birdType"`
  Description string `json:"what it does"`
}
```

https://play.golang.org/p/-_0XddCakR

- Unstructured data:

Nếu bạn có chuỗi JSON mà api trả về không có cấu trúc xác định hoặc tên thuộc tính mà bạn không chắc chắn, bạn không thể sử dụng struct để parse liệu của mình, vì có biết dữ liệu nó biến đổi khôn lường. Thay vào đó bạn có thể sử dụng map.

```
{
  "birds": {
    "pigeon":"likes to perch on rocks",
    "eagle":"bird of prey"
  },
  "animals": "none"
}
```
``
Không thể tạo một struct để biểu diễn dữ liệu như bên trên có tất cả các trường hợp biến đổi, lúc dài lúc ngắn, lúc thêm trường này lúc thêm trường kia.
``

Để đối phó với loại này ta tạo một map với key là string và value để ở kiểu interface:

https://play.golang.org/p/xbVxASrffo


Mỗi string tương ứng với thuộc tính của JSON, và nó được ánh xa đến kiểu interface{}tương ứng với giá trị, nó có thể là bất cứ kiểu dữ kiệu nào.