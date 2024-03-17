/**
 * 题目名称：选择排序 (Selection Sort)
 *
 * 题目理解：
 * 选择排序是一种简单直观的排序算法，它的工作原理是每一次从未排序的序列中找到最小（或最大）的元素，
 * 然后将其放到已排序序列的末尾。重复这个过程，直到所有元素都排序完毕。

 * 解题思路：
 * 1. 遍历整个数组，寻找当前未排序部分的最小（或最大）值。
   2. 找到后，将该最小（或最大）值与未排序部分的第一个元素交换位置，使得第一个元素成为当前未排序部分中最小（或最大）的元素。
   3. 接着对剩余未排序的部分（即原数组第二个到最后一个元素）重复步骤 1 和 2。
   4. 持续此过程直至遍历完整个数组，完成排序。

 * 注意要点：
 * - 选择排序的时间复杂度在最好、最坏和平均情况下都是 O(n^2)，因此对于大规模数据效率较低。
 * - 选择排序是不稳定的排序算法，因为相等的元素可能会改变原有的相对顺序。
 * - 由于每次只需找出剩余部分中的最小值并交换，不需要进行复杂的比较和移动操作，所以在数据规模较小或者特定场景下具有一定优势。

 * 时间复杂度：
 * - 平均时间复杂度：O(n^2)
 * - 最好情况时间复杂度：O(n^2) （当输入数组已经完全有序时）
 * - 最坏情况时间复杂度：O(n^2)
 * - 空间复杂度：O(1) （原地排序）

 */
// 代码实现
public class SelectionSort {

    /**
     * 选择排序方法
     * @param arr 待排序的整数数组
     */
    public static void selectionSort(int[] arr) {
        // 遍历整个数组
        for (int i = 0; i < arr.length - 1; i++) {
            // 找到剩余未排序部分中的最小值索引
            int minIndex = i;
            for (int j = i + 1; j < arr.length; j++) {
                if (arr[j] < arr[minIndex]) {
                    minIndex = j;
                }
            }

            // 将找到的最小值与未排序部分的第一个元素交换
            if (minIndex != i) {
                int temp = arr[i];
                arr[i] = arr[minIndex];
                arr[minIndex] = temp;
            }
        }
    }

    public static void main(String[] args) {
        int[] arrayToSort = {9, 2, 5, 1, 7};

        // 调用选择排序方法对数组进行排序
        selectionSort(arrayToSort);

        // 输出排序后的数组
        System.out.println("Sorted array: ");
        for (int num : arrayToSort) {
            System.out.print(num + " ");
        }
    }
}