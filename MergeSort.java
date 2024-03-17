/**
 * 题目名称：归并排序（MergeSort）
 *
 * 题目理解：
 * 归并排序是利用分治策略的一种排序算法。它将待排序的序列分成两半，对每一半分别进行排序，然后将结果合并成一个完整的有序序列。
 *
 * 解题思路：
 * 1. 将数组 arr 分成两个子数组 left 和 right，递归地对它们进行归并排序。
   2. 对于已经排序好的 left 和 right 子数组，创建一个新的临时数组 temp，从左到右比较两个子数组中的元素，并按照升序将较小的元素放入 temp 数组中，直到某个子数组遍历完，再将另一个子数组剩余部分直接复制到 temp 数组相应位置。
   3. 最后，将排好序的 temp 数组覆盖原数组 arr。

 * 注意要点：
 * - 关键数值：不存在特别关键数值，但需要注意的是分割和合并的过程。
   - 易错点：在合并过程中，需正确处理两个子数组的边界条件以及合并规则，确保最终得到的序列有序。
   - 临界值：当子数组长度为 1 时，认为其已经是有序的，此时不需要进一步拆分或合并。

 * 时间复杂度：
 * - 归并排序的时间复杂度无论最好、最坏还是平均情况下均为 O(n log n)，其中 n 是数组的长度。
   - 空间复杂度为 O(n)，主要用于合并过程中创建的临时数组。

 */
public class MergeSort {
    
    public static void mergeSort(int[] arr) {
        if (arr == null || arr.length <= 1) {
            return;
        }
        
        int[] temp = new int[arr.length];
        mergeSortInternal(arr, temp, 0, arr.length - 1);
    }

    private static void mergeSortInternal(int[] arr, int[] temp, int left, int right) {
        if (left >= right) {
            return;
        }
        // 特别注意，right + letf 会溢出。当 (0 + 3) /2 = 1
        int mid = left + (right - left) / 2; // 防止整数溢出，计算中间索引. mid 偶数时 偏移左 0123 ，取1
        mergeSortInternal(arr, temp, left, mid);
        mergeSortInternal(arr, temp, mid + 1, right);
        merge(arr, temp, left, mid, right);
    }

    private static void merge(int[] arr, int[] temp, int left, int mid, int right) {
        int i = left;
        int j = mid + 1;
        int k = left;

        while (i <= mid && j <= right) {
            if (arr[i] <= arr[j]) {
                temp[k++] = arr[i++];
            } else {
                temp[k++] = arr[j++];
            }
        }

        // 复制左边剩余部分
        while (i <= mid) {
            temp[k++] = arr[i++];
        }

        // 复制右边剩余部分
        while (j <= right) {
            temp[k++] = arr[j++];
        }

        // 将temp数组中的元素复制回原数组
        System.arraycopy(temp, left, arr, left, right - left + 1);
    }

    public static void main(String[] args) {
        int[] arr = {5, 2, 9, 1, 7, 6};
        mergeSort(arr);
        for (int num : arr) {
            System.out.print(num + " ");
        }
    }
}