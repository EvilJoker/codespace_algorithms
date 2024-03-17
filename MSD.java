public class MSD {

    private static int R = 256; // 256个 ascall
    private static final int M = 15;// 小数组的阀值，小数组场景切换算法
    private static String[] aux; // 辅助数组，辅助 分类
    private static int charAt(String s, int d){
        if (d < s.length()){
            return s.charAt(d);
        }else{
            return -1;
        }
    }

    public static void sort(String[] a) {
        int N = a.length;
        aux = new String[N];
        sort(a, 0, a.length, 0);
    }
    public static void sort(String[] a, int lo, int hi, int d) {

        // 以第 d 个字符为键将 a[lo] 到 a[hi] 排序
        // 小数组场景切换算法
        if(hi <= lo + M){
            Insertion.sort(a, lo, hi, d);
            return;
        }
        // 计算出现的频率
        int[] count = new int[R + 2];
        for (int i = lo; i <= hi; i++) {
            count[charAt(a[i], d) + 2]++; // 因为 charAt 可以返回 -1，所以加2

        }
        // 频率换算成索引
        for (int r = 0; r < R + 1; r++) {
            count[r + 1] += count[r];
        }
            
        // 数据分配，放到对应区间
        for (int i = lo; i <= hi; i++) {
            int c = charAt(a[i], d); // 字母编号，起始索引。 count[c + 1] 开始加 1
            aux[count[c + 1]++] = a[i];
        }
        for (int i = lo; i <= hi; i++) { // 回写
            a[i] = aux[i - lo];
        }
        // 递归 以每个字符为键，对子数组进行排序
        // 对 长度不够的字符，就不必要排序了，所以要缩小数组
        // 这里是，每个字母有自己的区间。lo + count[r] 到 lo + count[r + 1]
        // 比如z开头的所有字母，要按照 下一个字母 去 排，就是 d + 1
        // 巧妙地 去除了首字母。
        for(int r = 0 ; r < R; r++){
            sort(a, lo + count[r], lo + count[r + 1], d + 1);
        }


    }

}
