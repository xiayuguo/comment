
# comment

## Deploy
```shell
wget -O- https://raw.githubusercontent.com/xiayuguo/comment/master/quickstart.py | python -
```

## Features
- comment
- reply to comment
- 👍
- 👎

## Api Design
- get all comments `Get /comments`
- add/reply comment `Post /comment`
- edit comment `Put /comment/:id`
- delete comment `Delete /comment/:id`
- like `Post /comment/:id/like`
- dislike `Post /comment/:id/dislike`
- cancel like `Delete /comment/:id/like`
- cancel dislike `Delete /comment/:id/dislike`
- get count of like `Get /comment/:id/likes`
- get count of dislike `Get /comment/:id/dislikes`

[If you want more details about api, please click it.](https://github.com/hugoxia/comment/wiki/Api)

## Database Design
* user

|字段	|类型	|描述	|备注|
|:-|:-|:-|:-|
|\_id	|ObjectId	|用户 id||	
|avatar	|string	|头像	|图片链接|
|username	|string	|用户名|	

* comment

|字段	|类型	|描述	|备注|
|:-|:-|:-|:-|
|\_id	|ObjectId	|评论 id	||
|reply_id	|string	|被评论 id	|表 comment 的 \_id|
|user_id	|string	|用户 id	|表 user 的 \_id|
|content	|string	|评论内容	||
|create_time	|int	|创建时间	|时间戳|
|update_time	|int	|更新时间	|时间戳|

* like

|字段	|类型	|描述	|备注|
|:-|:-|:-|:-|
|\_id|ObjectId|点赞 id||
|comment_id	|string	|评论 id	|表 comment 的 \_id|
|user_id	|string	|用户 id	|表 user 的 \_id|
|is_like	|bool	|是否点赞|	|
