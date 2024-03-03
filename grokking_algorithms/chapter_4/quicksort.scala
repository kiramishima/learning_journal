package Chapter4

def quicksort(ls: List[Int]): List[Int] = {
  ls match {
    case Nil => Nil
    case pivot :: tail => {
      val (less, greater) = tail.partition(_ < pivot)
      quicksort(less) ::: pivot :: quicksort(greater)
    }
  }
}

object Chapter4 {
    def main(args: Array[String]): Unit =
        println(quicksort(List(10, 5, 2, 3)))
}