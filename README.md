#Build
```
docker build . -t omdbapi
```

#Run
set the API_KEY env variable as part of your deployment orchestration system and have it run
```
docker run -it --env API_KEY=$API_KEY omdbapi ./main -title Twister
```
or  pass it in directly
```
docker run -it --env API_KEY=<your api key here> omdbapi ./main -title Twister
```





#Design
It may be a bit over engineered, but there are multiple uses implied. Also, the *PrettyPrint* can be useful in debugging.

The output format is intended to be normalized over all score types (assuming they're percentages) making them decimal floats, as opposed to 0-100, for percentages gives a more usable output.