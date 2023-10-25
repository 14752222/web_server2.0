#!/bin/bash

# 从配置文件中提取版本号
version=$(grep "version =" ./config/config.ini | awk -F '=' '{print $2}' | tr -d ' ')

# 将版本号拆分为主版本、次版本和修订版本
IFS='.' read -ra parts <<< "$version"

# 增加修订版本
parts[2]=$((${parts[2]+1}))

# 处理版本号溢出
if [ "${parts[2]}" -eq 10 ]; then
  parts[2]=0
  parts[1]=$((${parts[1]+1}))
fi

if [ "${parts[1]}" -eq 10 ]; then
  parts[1]=0
  parts[0]=$((${parts[0]+1}))
fi

# 重新构建版本号
new_version="${parts[0]}.${parts[1]}.${parts[2]}"

# 将新版本号写入配置文件
sed -i "s/version = $version/version = $new_version/g" ./config.ini

# 输出版本号
echo "Original Version: $version"
echo "New Version: $new_version"



