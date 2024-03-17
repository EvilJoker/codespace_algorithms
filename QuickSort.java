/**
 * 题目名称：快速排序（QuickSort）
 *
 * 题目理解：
 * 快速排序是一种高效的比较排序算法，通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，然后再分别对这两部分记录继续进行排序，以达到整个序列有序的目的。
 *
 * 解题思路：
 * 1. 选择一个基准元素（pivot），通常选择数组的第一个元素或者最后一个元素。
   2. 重新排列数组，所有小于基准值的元素移动到基准前面，所有大于基准值的元素移动到基准后面（这个过程称为分区操作）。
   3. 对基准值前后的两个子序列重复上述步骤，递归调用快速排序，直到整个序列有序。

 * 注意要点：
 * - 关键数值：基准元素的选择对性能有影响，但不影响算法的正确性。
   - 易错点：在实现分区操作时要确保基准元素最终会被正确地放置在其最终位置上，并且左右两边的子数组都是无交叉的。
   - 临界值：对于小规模数组，可以设置一个阈值（例如10），当待排序数组长度小于该阈值时，转为使用插入排序等简单排序算法以减少递归深度和函数调用开销。

 * 时间复杂度：
 * - 最好情况下（每次都能均匀划分），时间复杂度为 O(n log n)。
   - 最坏情况下（输入数组已经完全有序或逆序），时间复杂度为 O(n^2)，但这种情况在实际应用中很少出现。
   - 平均情况下，时间复杂度也为 O(n log n)。

 */
public class QuickSort {
    
    // 三个函数： 排序， 内部的递归， 分区函数
    public static void quickSort(int[] arr){
        quickSortInternel(arr, 0, arr.length -1);
    }

    public static void quickSortInternel(int[] arr, int left, int right){
        if (left >= right){
            return;
        }

        int partionIndex = partion(arr, left, right);
        quickSortInternel(arr, left, partionIndex - 1);
        quickSortInternel(arr, partionIndex + 1, right);

    }

    public static int partion(int[] arr, int left, int right){
        int slot = arr[right];
        int i = left -1;

        for(int j =left;j<right;j++){
            if(arr[j]<slot){
                i++;
                swap(arr, i, j);
            }
        }
        swap(arr, i+1, right);
        return i + 1;

    }
    private static void swap(int[] arr, int i, int j) {
        int temp = arr[i];
        arr[i] = arr[j];
        arr[j] = temp;
    }

    public static void main(String[] args) {
        int[] arr = {5, 3, 7, 6, 2, 8, 1};
        quickSort(arr);
        for (int num : arr) {
            System.out.print(num + " ");
        }
    }
}