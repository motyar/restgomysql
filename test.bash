#!/bin/bash
curl -XDELETE "http://localhost:1234/api/1"
curl -XDELETE "http://localhost:1234/api/2"
curl -XDELETE "http://localhost:1234/api/3"
curl -XDELETE "http://localhost:1234/api/4"
mt=`curl "http://localhost:1234/api/" 2>/dev/null`
echo "mt;$mt;"
if [ $mt != '[]' ]; then
  echo "Warning: There should not be any entries"
fi
curl -X POST -d 'name=new' http://localhost:1234/api/
one=`curl "http://localhost:1234/api/" 2>/dev/null`
if [ $one == '[]' ]; then
  echo "Warning: There should be an entry."
fi
