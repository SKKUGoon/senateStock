# 0.0.1
<p>
create reJSON dataObject functions
</p>

[Add]
- ./dataObject
  - dao.go
    - struct DataBaseConfig
    - func processKey
  - reJson.go
    - (use golang generics)
    - func JSONGet
    - func JSONSet
- ./README.md
- ./Dockerfile
- ./.gitignore
- ./main.go
  - struct Test - test sample json file

[Change]
- ./dataObject
  - dao.go
    - func Conn 
      - store address and password on a different file - token.json
      - if pong is negative - database not working.

[Fix]

[Remove]

