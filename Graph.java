import java.util.ArrayList;
import java.util.List;

public class Graph {
    private final int v; // 顶点数目
    private int E; // 边数目
    private List<List<Integer>> adj;// 邻接表

    public Graph(int v) {
        String a;
        a.
        this.v = v;
        // 初始化邻接表
        adj = new ArrayList<>(v);
        for (int i = 0; i < v; i++) {
            adj.add(new ArrayList<>());
        }
    }
# 题目名称：Z字形变换

# 题目理解：
# 此题要求将输入字符串 s 按照指定的行数 numRows 进行 Z 字形排列，即将字符串排布成类似字母 Z 的形状，
# 然后从上到下、从左到右逐行读取字符组成一个新的字符串返回。

# 解题思路：
# 1. 初始化一个大小为 numRows 的列表，用于存储每一行的结果字符串。
# 2. 设置初始行索引 i 为 0，以及一个标志位 flag，用于控制行索引的增减。初始设置 flag 为 -1，使得第一次循环时 i 会从 0 增加到 1，进入第二行。
# 3. 遍历输入字符串 s 中的每一个字符 c，将其添加到当前行 res[i] 的末尾。
# 4. 当到达首行（i=0）或末行（i=numRows-1）时，翻转 flag 的符号，使行索引在奇数行间递增，在偶数行间递减，从而形成 Z 字形排列。
# 5. 根据翻转后的 flag 更新行索引 i。
# 6. 循环结束后，将列表 res 中的所有字符串拼接起来，形成最终的 Z 字形排列字符串。

# 注意要点：
# - 当 numRows 等于 1 或者等于字符串长度时，不需要做任何变换，直接返回原字符串。
# - 通过改变标志位 flag，可以简洁地实现行索引的交替增减。

# 时间复杂度：
# O(n)，其中 n 为字符串 s 的长度，因为需要遍历字符串 s 一次。


public class Solution {
    public String convert(String s, int numRows) {
        // 特殊情况处理，行数为1时无需变换
        if (numRows == 1) {
            return s;
        }

        // 字符串长度
        int len = s.length();
        // 核心思路是，每行都为一个字符串。遍历 s 的每个字符确认放入哪一行
        // 规律是，先是从 0放到 n,再从n 放到 0 字符串
        // 边界： 遇到 0 和 n-1 要变向
        StringBuilder[] rows = new StringBuilder[numRows];
        for(int i=0;i<numRows;i++){
            rows[i] = new StringBuilder();
        }
        int flag = -1;
        int rows_num =0;
        for (int i=0;i<len;i++){
            char c = s.charAt(i);
            rows[rows_num].append(c);
            if(rows_num==0 ||rows_num ==numRows-1){
                flag =-flag;
            }
            rows_num +=flag;
        }
        StringBuilder ret = new StringBuilder();
        for(int j =0; j<numRows;j++){
            ret.append(rows[j]);
        }
        return ret.toString();

    }
}
    public void addEdge(int v, int w) {
        adj.get(v).add(w);
        adj.get(w).add(v); // 有向图， 需要删除这个
    }

    public Iterable<Integer> adj(int v) {
        return adj.get(v);
    }
}
