'''
LeetCode: Judge Route Circle

description:
Initially, there is a Robot at position (0, 0). Given a sequence of its moves, judge if this robot makes a circle, which means it moves back to the original place.

The move sequence is represented by a string. And each move is represent by a character. The valid robot moves are R (Right), L (Left), U (Up) and D (down). The output should be true or false representing whether the robot makes a circle.

author: Ehco1996
time: 2017-08-14
'''


class Solution(object):
    def judgeCircle(self, moves):
        """
        :type moves: str
        :rtype: bool
        """
        rec = {
            'U': 0,
            'D': 0,
            'R': 0,
            'L': 0
        }

        for move in moves:
            rec[move] += 1

        if rec['U'] == rec['D'] and rec['R'] == rec['L']:
            return True
        else:
            return False


moves = 'UD'

print(Solution().judgeCircle(moves))
