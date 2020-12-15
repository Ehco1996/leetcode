#
# @lc app=leetcode.cn id=752 lang=python3
#
# [752] 打开转盘锁
#

# @lc code=start
class Solution:
    def up(self, s, idx):
        """将指定位置的数向上波动一次"""
        t = int(s[idx])
        if t == 9:
            t = 0
        else:
            t += 1
        return s[:idx] + str(t) + s[idx + 1 :]

    def down(self, s, idx):
        """将指定位置的数向下波动一次"""
        t = int(s[idx])
        if t == 0:
            t = 9
        else:
            t -= 1
        return s[:idx] + str(t) + s[idx + 1 :]

    def openLock(self, deadends, target: str):
        """
        仔细想想，这就可以抽象成一幅图，每个节点有 8 个相邻的节点，又让你求最短距离
        """
        q = ["0000"]
        step = 0
        deadends = set(deadends)
        visited = set()

        while len(q) > 0:
            size = len(q)
            for i in range(0, size):
                cur = q.pop(0)
                if cur in deadends:
                    continue
                if cur == target:
                    return step

                # 开始拨动锁盘
                for i in range(0, len(cur)):
                    up = self.up(cur, i)
                    if up not in visited:
                        q.append(up)
                        visited.add(up)

                    down = self.down(cur, i)
                    if down not in visited:
                        q.append(down)
                        visited.add(down)
            step += 1

        return -1

    def test(self, deadends=["0201", "0101", "0102", "1212", "2002"], target="0202"):
        return self.openLock(deadends=deadends, target=target)


# @lc code=end

