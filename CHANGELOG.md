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

# 0.0.2
<p>
create re_json re_hash dao_struct dao_conn with golang generics. 
upgrade to implement go-redis v8
</p>

[Add]
- ./dataObject
  - re_json.go
    - method JsonGet
    - method JsonSet
  - re_hash.go
    - method HashGet
    - method HashSet
  - dao_struct.go
    - RdbJson interface - JsonGet, JsonSet
    - struct RedisConn
- ./main.go
  - func reJson - testing function
  
[Change]
- ./dataObject
  - dao_conn.go
    - func conn - delete ping pong message check

[Fix]

[Remove]
