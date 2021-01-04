# 关键点

1. 对于 BST 的每一个节点node，左子树节点的值都比node的值要小，右子树节点的值都比node的值大。

2. 对于 BST 的每一个节点node，它的左侧子树和右侧子树都是 BST

3. BST，除了它的定义，还有一个重要的性质：BST 的中序遍历结果是有序的（升序）。

4. BST，除了它的定义，还有一个重要的性质：BST 的后序遍历结果是有序的（降序）。

``` java
void BST(TreeNode root, int target) {
    if (root.val == target)
        // 找到目标，做点什么
    if (root.val < target)
        BST(root.right, target);
    if (root.val > target)
        BST(root.left, target);
}
```
