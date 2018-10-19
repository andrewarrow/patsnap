# patsnap

patsnap 'msg,time,my_field1,my_field2' '@my_field1:10 @msg:*something*'

will make a url like:

https://app.datadoghq.com/logs?cols=%5B%22log_msg%22%2C%22log_time%22%2C%22log_my_field1%22%2C%22log_my_field2%22%5D&event&from_ts=1539980475000&index=main*live=true&query=%40my_field1%3A10+%40msg%3A%2Asomething%2A&stream_sort=desc&to_ts=1539981375000


