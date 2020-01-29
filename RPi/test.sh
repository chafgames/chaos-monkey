#!/bin/bash

curl http://localhost:5000/images
curl http://localhost:5000/images 2> /dev/null | grep \" | sed -e 's/"//g' -e 's/,//g' -e 's/ //g' > images.out
for img in `cat images.out`
do
    curl http://localhost:5000/image/$img -X POST
done

curl http://localhost:5000/image/0000000000000000000000000000000000000000000000000000000000000000 -X POST
curl http://localhost:5000/image/1111111100000000000000000000000000000000000000000000000000000000 -X POST
curl http://localhost:5000/image/1111111111111111111111111111111111111111111111111111111111111111 -X POST

rm images.out
