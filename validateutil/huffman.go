package validateutil

/*
哈夫曼编码：
1. 首先需要统计字符出现的频率
2. 将其从小到大排序
3. 将最左边两个字符统计到一颗树上，自顶向下。权重为两者相加，然后重新排序（这里需要优先队列）
4. 将左边标为1，右边标为0。则每一个字符的哈夫曼编码就都出现了。
*/


