########map底层源码分析

1.hmap，mapextra，bmap结构
2.创建
    1.是否创建溢出桶，取决于元素多少
    2.正常桶和溢出桶是挨着的，数组
    3.

3.扩容

扩容条件：装载因子已经超过 6.5； 哈希使用了太多溢出桶；

等量扩容：新创建的桶数量和旧的桶一样，只创建，不实际转移数据
翻倍扩容：


4.缩容 delete关键字进行删除




5.渐进式搬迁元素
6.读写
方法一：
_ = hash[key]

    1.先找到是哪一个桶，然后计算hash高8位，挨个遍历正常桶和溢出桶
    2.内存空间布局：tophash+key+value,挨着


方法二：
for k, v := range hash {
// k, v
}













