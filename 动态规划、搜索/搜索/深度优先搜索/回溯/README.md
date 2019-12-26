
## 简介
1. 个人认为回溯一般要解决的问题是: 需要找出结果集的所有结果，而不只是求结果集的长度，举个例子: 比如求递增子序列，在动态规划模块也有这道题，但是当时求的是递增子序列长度，而回溯模块这要求求出所有的递增子序列结果。


## 解决的问题是什么？
元素大于0 | 有重复元素 |  可重复选取|  限定选取个数 | 限制选取总和| 问题类型|对应文件|
-|-|-|-|-|-|-|
<font color="00bb00">**-**<font/>| <font color="bb0000">**No**<font/> |<font color="bb0000">**No**<font/> |<font color="00bb00">**Yes**<font/>|<font color="bb0000">**No**<font/>|**排列问题**|**求全排列_1.go**|
<font color="00bb00">**-**<font/>|  <font color="00bb00">**Yes**<font/> |<font color="bb0000">**No**<font/> |<font color="00bb00">**Yes**<font/>|<font color="bb0000">**No**<font/>|**排列问题**|**求全排列_2.go**|
<font color="00bb00">**-**<font/>|  <font color="bb0000">**No**<font/> |<font color="bb0000">**No**<font/> |<font color="00bb00">**Yes**<font/>|<font color="bb0000">**No**<font/>|**组合问题**|**求组合.go**|
<font color="00bb00">**-**<font/>|  <font color="bb0000">**No**<font/> |<font color="bb0000">**No**<font/> |<font color="bb0000">**No**<font/>|<font color="bb0000">**No**<font/>|**组合问题**|**求子集_无重复元素.go**|
<font color="00bb00">**-**<font/>|<font  color="00bb00">**Yes**<font/> |<font color="bb0000">**No**<font/> |<font color="bb0000">**No**<font/>|<font color="bb0000">**No**<font/>|**组合问题**|**求子集_有重复元素.go**|
<font color="00bb00">**Yes**<font/>|  <font color="bb0000">**No**<font/> |<font color="00bb00">**Yes**<font/> |<font color="bb0000">**No**|<font color="00bb00">**Yes**<font/>|**组合问题**|**求数组中等于目标值的组合_1.go**|
<font color="00bb00">**Yes**<font/>| <font color="00bb00">**Yes**<font/> |  <font color="bb0000">**No**<font/> |<font color="bb0000">**No**|<font color="00bb00">**Yes**<font/>|**组合问题**|**求数组中等于目标值的组合_2.go**|
<font color="00bb00">**Yes**<font/>| <font color="bb0000">**No**<font/> |  <font color="bb0000">**No**<font/> |<font color="00bb00">**Yes**<font/>|<font color="00bb00">**Yes**<font/>|**组合问题**|**求数组中等于目标值的组合_3.go**|
<font color="00bb00">**Yes**<font/>| <font color="bb0000">**No**<font/> | <font color="00bb00">**Yes**<font/> |<font color="bb0000">**No**<font/>|<font color="00bb00">**Yes**<font/>|**排列问题**|**求数组中等于目标值的组合_4.go**|
<font color="00bb00">**-**<font/>| <font color="bb0000">**No**<font/> |<font color="bb0000">**No**<font/> |<font color="bb0000">**No**<font/>|<font color="bb0000">**No**<font/>|**-**|**求递增子序列.go**|

