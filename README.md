# LeetCode 题集

## 1. 两数之和

* 问题描述:

    给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

* 示例:
    给定 nums = [2, 7, 11, 15], target = 9
    因为 nums[0] + nums[1] = 2 + 7 = 9
    所以返回 [0, 1]

* [题解](./leetcode-go/sum_test.go)

## 2. 两数相加

* 问题描述:

    给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

* 示例:
    
    给定 nums = [2, 7, 11, 15], target = 9  
    因为 nums[0] + nums[1] = 2 + 7 = 9  
    所以返回 [0, 1]

* [题解](./leetcode-go/sum2_test.go)

## 3. 无重复字符最长子串

* 问题描述:

    给定一个字符串，找出其中最长没有重复子串的长度。

* 示例:
    
    给定字符串"abcabcbb"，最长不重复子串为"abc"，长度为3  
    给定字符串"bbbbb"，最长不重复子串为"b"，长度为1  
    给定字符串"pwwkew"，最长不重复子串为"wke"或"kew"，长度为3，注意子串必须要连续，"pwke"不符合要求  

* [题解](./leetcode-go/longest_string_test.go)

## 4. 寻找两个有序数组的中位数

* 问题描述:

    给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
    请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
    你可以假设 nums1 和 nums2 不会同时为空。

* 示例 1:
    
    nums1 = [1, 3]  
    nums2 = [2]  
    则中位数是 2.0  

* 示例 2:

    nums1 = [1, 2]  
    nums2 = [3, 4]  
    则中位数是 (2 + 3)/2 = 2.5  

* [题解](./leetcode-go/median_test.go)

## 4. 最长回文子串

* 问题描述:

    给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

* 示例 1:
    
    输入: "babad"  abba abcba
    输出: "bab"  
    注意: "aba" 也是一个有效答案。  

* 示例 2:

    输入: "cbbd"  
    输出: "bb"

* [题解](./leetcode-go/longest_palindrome_test.go)

## 6. Z 字形变换

* 问题描述:

    将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
    比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：  
    ```
    L   C   I   R  
    E T O E S I I G  
    E   D   H   N 
    ```
    ```
    0   4   8    12
    1 3 5 7 9 11 13 15
    2   6   10   14
    ```
    ```
    0     6       12
    1   5 7    11 13
    2 4   8 10    14
    3     9       15
    ```
    之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
    请你实现这个将字符串进行指定行数变换的函数：
    string convert(string s, int numRows);

* 示例 1:

    输入: s = "LEETCODEISHIRING", numRows = 3
    输出: "LCIRETOESIIGEDHN"

* 示例 2:

    输入: s = "LEETCODEISHIRING", numRows = 4
    输出: "LDREOEIIECIHNTSG"
    
* [题解](./leetcode-go/convert_test.go)
