# SQL-Performance
Do more different tests based on https://studygolang.com/articles/3022

## Result
Each insert and update has size 100.   
[Insert]   
Method-1 Exec within the loop total time =  0.3380945(s)   
Method-2 Prepare and Exec within the loop. total time =  0.3410885(s)   
Method-3 Prepare. Then Exec within the loop. total time =  0.2513359(s)   
Method-4 DB Transaction. Then Euec within the loop. total time: =  0.0299109(s)   
Method-5 DB Transaction within the Loop. total time =  0.3650823(s)   
Method-6 Insert with multiple VALUES sets. total time =  0.0059277(s)   
[Query]   
Method-1 db.Query. total time =  0.0020196(s)   
Method-2 db.Prepare -> Query. total time =  0.001995(s)   
Method-3 db Transaction. total time =  0.0009982(s)   
[Update]   
Method-1 Exec within the loop. total time =  0.0339128(s)   
Method-2 Prepare and Exec within the loop.. total time =  0.0339213(s)   
Method-3 Prepare. Then Exec within the loop. total time =  0.017937(s)   
Method-4 DB Transaction. Then Euec within the loop. total time =  0.0319183(s)   
Method-5 DB Transaction within the loop. total time =  0.0528547(s)   
Method-6 UPDATE with duplicate keys. total time =  0.010971(s)   
[Query]   
Method-1 db.Query. total time =  0.0019958(s)   
Method-2 db.Prepare -> Query. total time =  0.0019946(s)   
Method-3 db Transaction. total time =  0.0019943(s)   
[Delete]   
Method-1 Exec Delete. total time =  0.523601(s)   
Method-2 Prepare and Exec within the loop.. total time: 0.3380925(s)   
Method-3 Prepare. Then Exec within the loop. total time =  0.3939471(s)   
Method-4 db Transaction. Exec Delete within the loop. total time =  0.0448794(s)   
Method-5 DB Transaction within the loop. total time =  0.3361002(s)   
Method-6 Just Delete. total time =  0.0349062(s)   