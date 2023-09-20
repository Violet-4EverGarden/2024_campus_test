# n个数字的全排列，使用回溯法解决；假设n张卡片的数字构成的数组为nums，那么具体思路如下：
# 1、首先对nums进行排序，避免使用0作为n位数的第一位；
# 2、维护一个used数组（bool型），表示当前递归过程中nums的哪些元素(index)已被使用；维护一个path数组（int型），存储当前搜索路径中已见过的元素；维护一个result数组，记录有多少组满足条件的path；
# 3、从根节点开始递归搜索：若len(path)==0且nums[i]==0，由于0不能作为n位数的第一位，因此直接continue（i++）；否则将nums[i]加入path，并将used[i]置为True表示索引为i的元素已使用；然后递归遍历下一层，当len(path)==n时结束递归，回溯到上一层；回溯后需要将nums[i]从path中弹出，同时将used[i]置为False。当完成nums的for循环遍历后，即可返回最终结果：len(result)。
# 这里用到了回溯法的三部曲：（1）确定递归函数参数；（2）确定递归终止条件；（3）确定单层搜索逻辑。
# 使用回溯法的关键在于要将递归搜索树的结构给画出来：for循环是横向遍历，递归则是纵向遍历。

class Solution:
    # 全排列
    def permute(self, nums):
        result = []
        nums.sort()  # 先排序
        self.backtracking(nums, [], [False]*len(nums), result)
        return len(result)

    # 回溯法
    def backtracking(self, nums, path, used, result):
        # 递归终止条件
        if len(path) == len(nums):
            result.append(path[:])
            return
        for i in range(len(nums)):
            # 0不能作为第一位
            if len(path)==0 and nums[i] == 0:
                continue
            # 如果当前数在上一层已使用，也跳过
            if used[i]:
                continue
            # 处理当前层节点
            used[i] = True
            path.append(nums[i])
            # 递归，搜索下一层
            self.backtracking(nums, path, used, result)
            # 回溯，撤销当前层的处理
            path.pop()
            used[i] = False


# 输入（用空格分隔）：1 2 5 8 0 6 4
# 输出：4320
if __name__ == "__main__":
    n = input()
    input_list = n.split()
    nums = [int(x) for x in input_list]
    # print(nums)
    solution = Solution()
    print(solution.permute(nums))