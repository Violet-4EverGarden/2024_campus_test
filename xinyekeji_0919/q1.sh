awk '/tcp/ {print $6}' | sort | uniq -c | sort -nr | awk '{print $2" "$1}'

# 统计所有TCP连接状态，并按照连接状态的不同数量进行排序和输出
# awk '/tcp/ {print $6}': 执行 awk 命令，筛选出包含 "tcp" 字符串的行，并打印出每行中第 6 列的内容，也就是 TCP 连接状态。
# sort: 对输出的 TCP 连接状态进行排序。
# uniq -c: 统计并输出每个 TCP 连接状态的数量。
# sort -nr: 对 TCP 连接状态的数量进行倒序排序。
# awk '{print $2" "$1}': 打印出每个 TCP 连接状态的名称和数量，格式为：状态名称 数量。