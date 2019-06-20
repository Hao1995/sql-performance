# SQL-Performance
Do more different tests based on https://studygolang.com/articles/3022

## Result
[Insert]   
Method-1 Exec within the loop total time =     0.2463522   
Method-2 Prepare and Exec within the loop. total    time =  0.2912122   
Method-3 Prepare. Then Exec within the loop.    total time =  0.2323949   
Method-4 DB Transaction. Then Euec within the    loop. total time: =  0.0338908   
Method-5 DB Transaction within the Loop. total    time =  0.2513278   
[Query]   
Method-1 db.Query. total time =  0.0029884   
Method-2 db.Prepare -> Query. total time =     0.0029903   
Method-3 db Transaction. total time =  0.0029941   
[Update]   
Method-1 Exec within the loop. total time =     0.0488677   
Method-2 Prepare and Exec within the loop.. total    time =  0.0428858   
Method-3 Prepare. Then Exec within the loop.    total time =  0.019947   
Method-4 DB Transaction. Then Euec within the    loop. total time =  0.0309157   
Method-5 DB Transaction within the loop. total    time =  0.0508648   
[Query]   
Method-1 db.Query. total time =  0.0019958   
Method-2 db.Prepare -> Query. total time =     0.001993   
Method-3 db Transaction. total time =  0.0019947   
[Delete]   
Method-1 Exec Delete. total time =  0.3540523   
Method-2 Prepare and Exec within the loop.. total    time: 0.4198821   
Method-3 Prepare. Then Exec within the loop.    total time =  0.3859615   
Method-4 db Transaction. Exec Delete within the    loop. total time =  0.0388956   
Method-5 DB Transaction within the loop. total    time =  0.3789859   