#
# @lc app=leetcode.cn id=105 lang=python3
#
# [105] 从前序与中序遍历序列构造二叉树
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution:
    def buildTree(self, preorder, inorder):
        """
        https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247487270&idx=1&sn=2f7ad74aabc88b53d94012ceccbe51be&chksm=9bd7f12eaca078384733168971147866c140496cb257946f8170f05e46d16099f3eef98d39d9&scene=21#wechat_redirect

        用preorder确定左树 用inorder确定右数
        """
        return self.helper(preorder, 0, len(preorder) - 1, inorder, 0, len(inorder) - 1)

    def helper(self, preorder, p_start, p_end, inorder, in_start, in_end):
        if p_start > p_end:
            return None

        root_val = preorder[p_start]
        # 在inorder中找到root位置
        idx = inorder.index(root_val)

        # 左子树的大小
        left_size = idx - in_start

        root = TreeNode(root_val)
        root.left = self.helper(
            preorder, p_start + 1, p_start + left_size, inorder, in_start, idx - 1
        )
        root.right = self.helper(
            preorder, p_start + left_size + 1, p_end, inorder, idx + 1, in_end
        )
        return root

    def test(self, preorder=[3, 9, 20, 15, 7], inorder=[9, 3, 15, 20, 7]):
        return self.buildTree(preorder, inorder)

    def _buildTree(self, preorder, inorder):
        """
        先把序列拆分成root和左右两颗子树，然后反过来递归
	 			in order
        +------------+------+-------------+
        | left child | root | right child |
        +------------+------+-------------+

                    pre order
        +------------+-------------+------+
        | root| left child | right child  |
        +------------+-------------+------+
        """
        if len(inorder) == 0 and len(preorder) == 0:
            return None

        root_val = preorder[0]
        root = TreeNode(root_val)
        pos = inorder.index(root_val)
        if len(inorder) == 1 and len(preorder) == 1:
            return root
        root.Left = self.buildTree(preorder[1 : pos + 1], inorder[:pos])
        root.Right = self.buildTree(preorder[pos + 1 :], inorder[pos + 1 :])
        return root


# @lc code=end

