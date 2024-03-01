package Chapeter3

def fact(x: Int): Int =
    if x == 1 then
        return 1
    else
        return x * fact(x-1)

object Chapter3 {
    def main(args: Array[String]): Unit =
        println(fact(5)) // 120
}