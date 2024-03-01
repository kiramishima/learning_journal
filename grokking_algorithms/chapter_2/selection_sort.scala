package Chapter2

def findSmallest(xs: List[Int]): List[Int] =
    (List(xs.head) /: xs.tail) {
        (ys, x) =>
            if(x < ys.head) (x :: ys)
            else            (ys.head :: x :: ys.tail)
    }

def selectionSort(xs: List[Int]): List[Int] =
    if (xs.isEmpty) then List()
    else
        val ys = findSmallest(xs)
        if (ys.tail.isEmpty) then
            ys
        else
            ys.head :: selectionSort(ys.tail)

object Chapter2 {
    def main(args: Array[String]): Unit =
        val nums = List(5, 3, 6, 2, 10)
        println(nums)
        println(selectionSort(nums))
}