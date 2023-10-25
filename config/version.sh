#!/bin/sh

# 从配置文件中提取版本号
version=$(grep "version =" ./config/config.ini | awk -F '=' '{print $2}' | tr -d ' ')

# 将版本号拆分为主版本、次版本和修订版本
IFS='.' read -r part1 part2 part3 <<EOF
$version
EOF

# 增加修订版本
part3=$((part3 + 1))

# 处理版本号溢出
if [ $part3 -eq 10 ]; then
  part3=0
  part2=$((part2 + 1))
fi

if [ $part2 -eq 10 ]; then
  part2=0
  part1=$((part1 + 1))
fi

# 重新构建版本号
new_version="$part1.$part2.$part3"

# 将新版本号写入配置文件
sed -i "s/version = $version/version = $new_version/g" ./config/config.ini

# 输出版本号
echo "Original Version: $version"
echo "New Version: $new_version"
