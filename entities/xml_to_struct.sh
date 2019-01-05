#!/bin/sh

cd `dirname $0`

if [ $# -ne 2 ]; then
  echo "Usage: sh $0 {target_name} {prefix}"
  exit 1
fi

name=$1
prefix=$2
echo "package entity" > $name.go
echo "" >> $name.go
echo "import \"encoding/xml\"" >> $name.go
echo "" >> $name.go
chidley -G -x -t -e "Mws${prefix}" $name.xml >> $name.go

# replace string string => String string
sed -i '' 's/string string/String string/' $name.go
