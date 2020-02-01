#!/bin/bash

curl http://localhost:5000/text/Hi_There
curl http://localhost:5000/images
curl http://localhost:5000/images 2> /dev/null | grep \" | sed -e 's/"//g' -e 's/,//g' -e 's/ //g' > images.out
for img in `cat images.out`
do
    curl http://localhost:5000/image/$img -X POST
done

for i in 1 2 4 8 16 32 64 128 256
do
    printf -v j "%08d" $(echo "obase=2;$i-1" | bc)
    curl http://localhost:5000/image/$j$j$j$j$j$j$j$j -X POST
done

curl http://localhost:5000/image/1111111111111111111111111111111111111111111111111111111111111111 -X POST
curl http://localhost:5000/image/1101101101101101101101101101101101101101101101101101101101101101 -X POST
curl http://localhost:5000/image/1100110000110011110011000011001111001100001100111100110000110011 -X POST
for i in {1..5}
do
    curl http://localhost:5000/image/1100110011001100001100110011001111001100110011000011001100110011 -X POST
    curl http://localhost:5000/image/0011001100110011110011001100110000110011001100111100110011001100 -X POST
done

rm images.out
