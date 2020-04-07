#!/bin/sh

version="0.1.0"
main_file="../cmd"
target_dir="../release/"
target_file="ytdlgo"

rm -rf $target_dir

if [[ ! -x "$target_dir"  ]]; then
        mkdir  -p "$target_dir"
fi

if [[ -z ${target_file}  ]];then
        target_file=${main_file%".go"*}
fi

echo "build start"

for goos in "linux" "darwin" "freebsd" "windows"
    do
    # For AMD64
    out_f_name=$target_dir$target_file-$goos-amd64
    if [ "$goos" == "windows" ]; then
        out_f_name=$out_f_name.exe
    fi
    GOOS=$goos GOARCH=amd64 go build -o $out_f_name $main_file
    tar JcvfP $target_dir/$target_file-$goos-amd64.tar.xz $out_f_name
    rm -f $out_f_name
    # For 386
    
    out_f_name=$target_dir$target_file-$goos-386
    if [ "$goos" == "windows" ]; then
        out_f_name=$out_f_name.exe
    fi
    GOOS=$goos GOARCH=386 go build -o $out_f_name $main_file
    tar JcvfP $target_dir/$target_file-$goos-386.tar.xz $out_f_name
    rm -f $out_f_name
done

cd $target_dir
for file in ./*
do
    md5 -r $file >> sha1sum.txt
done
