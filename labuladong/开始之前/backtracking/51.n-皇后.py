#
# @lc app=leetcode.cn id=51 lang=python3
#
# [51] N 皇后
#

# @lc code=start
class Solution:
    def new_board(self, n):
        board = []
        for _ in range(n):
            board.append(["."] * n)
        return board

    def flat_board(self, board):
        res = []
        for row in board:
            res.append("".join(row))
        return res

    def is_valid(self, board, row, idx):
        """
        判断在board的第row行的idx处放一个皇后，看看是否合法
        由于决策是从第一行开始判断的，所以只需要判断
        1. 该列（idx）列上方是否有皇后
        2. 该位置左上方是否有皇后
        3. 改位置右上方是否有皇后
        """

        # 1.
        for i in range(0, row):
            if board[i][idx] == "Q":
                return False
        # 2.
        i, j = row - 1, idx - 1
        while i >= 0 and 0 <= j < len(board):
            if board[i][j] == "Q":
                return False
            i -= 1
            j -= 1
        # 3.
        i, j = row - 1, idx + 1
        while i >= 0 and 0 <= j < len(board):
            if board[i][j] == "Q":
                return False
            i -= 1
            j += 1
        return True

    def solveNQueens(self, n: int):
        res = []

        def bt(board, row):
            if row == n:
                # 选择到最后一行啦
                res.append(self.flat_board(board))
                return
            # 开始做选择
            for idx in range(n):
                if not self.is_valid(board, row, idx):
                    # 排除掉不合法的选择
                    continue
                # 做选择 塞皇后
                board[row][idx] = "Q"
                # 进入下一次决策
                bt(board, row + 1)
                # 撤销选择
                board[row][idx] = "."

        bt(self.new_board(n), 0)
        return res

    def test(self, n=4):
        res = self.solveNQueens(n)
        for ans in res:
            for line in ans:
                print(line)
            print("-------------------")


# @lc code=end
