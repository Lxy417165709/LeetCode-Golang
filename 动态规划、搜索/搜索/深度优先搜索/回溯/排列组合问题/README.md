
## 简介
1. 个人认为回溯一般要解决的问题是: 需要找出结果集的所有结果，而不只是求结果集的长度，举个例子: 比如求递增子序列，在动态规划模块也有这道题，但是当时求的是递增子序列长度，而回溯模块这要求求出所有的递增子序列结果。

## 解决的问题是什么？
元素大于0 | 有重复元素 |  可重复选取|  限定选取个数 | 限制选取总和| 问题类型|对应题目|
-|-|-|-|-|-|-|
<font color="00bb00">**-**<font/>| <font color="bb0000">**No**<font/> |<font color="bb0000">**No**<font/> |<font color="00bb00">**Yes**<font/>|<font color="bb0000">**No**<font/>|**排列问题**|**LeetCode_46_全排列**|
<font color="00bb00">**-**<font/>|  <font color="00bb00">**Yes**<font/> |<font color="bb0000">**No**<font/> |<font color="00bb00">**Yes**<font/>|<font color="bb0000">**No**<font/>|**排列问题**|**LeetCode_47_全排列 II**|
<font color="00bb00">**-**<font/>|  <font color="bb0000">**No**<font/> |<font color="bb0000">**No**<font/> |<font color="00bb00">**Yes**<font/>|<font color="bb0000">**No**<font/>|**组合问题**|**LeetCode_77_组合**|
<font color="00bb00">**-**<font/>|  <font color="bb0000">**No**<font/> |<font color="bb0000">**No**<font/> |<font color="bb0000">**No**<font/>|<font color="bb0000">**No**<font/>|**组合问题**|**LeetCode_78_子集**|
<font color="00bb00">**-**<font/>|<font  color="00bb00">**Yes**<font/> |<font color="bb0000">**No**<font/> |<font color="bb0000">**No**<font/>|<font color="bb0000">**No**<font/>|**组合问题**|**LeetCode_90_子集 II**|
<font color="00bb00">**Yes**<font/>|  <font color="bb0000">**No**<font/> |<font color="00bb00">**Yes**<font/> |<font color="bb0000">**No**|<font color="00bb00">**Yes**<font/>|**组合问题**|**LeetCode_39_组合总和**|
<font color="00bb00">**Yes**<font/>| <font color="00bb00">**Yes**<font/> |  <font color="bb0000">**No**<font/> |<font color="bb0000">**No**|<font color="00bb00">**Yes**<font/>|**组合问题**|**LeetCode_40_组合总和 II**|
<font color="00bb00">**Yes**<font/>| <font color="bb0000">**No**<font/> |  <font color="bb0000">**No**<font/> |<font color="00bb00">**Yes**<font/>|<font color="00bb00">**Yes**<font/>|**组合问题**|**LeetCode_216_组合总和 III**|
<font color="00bb00">**Yes**<font/>| <font color="bb0000">**No**<font/> | <font color="00bb00">**Yes**<font/> |<font color="bb0000">**No**<font/>|<font color="00bb00">**Yes**<font/>|**排列问题**|**LeetCode_377_组合总和 Ⅳ**|

