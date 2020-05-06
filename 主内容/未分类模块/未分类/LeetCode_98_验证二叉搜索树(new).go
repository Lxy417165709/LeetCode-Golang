package main


var lastNum = -10000000000
func isValidBST(root *TreeNode) bool {
    lastNum = -10000000000
    return solve(root)
}

func solve(root *TreeNode) bool{
    if root==nil{
        return true
    }
    if solve(root.Left)==false{
        return false
    }
    if root.Val > lastNum {
        lastNum = root.Val
    }else{
        return false
    }
    if solve(root.Right)==false{
        return false
    }
    return true
}