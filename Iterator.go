package main

import (
   "fmt"
)

var pointer int

type QueryResultsImpl struct{
   array [5] int
}

type QueryResults interface{
   fetchData()
}

type QueryResultsIterator interface {
   isDone() bool
   current() int
   first() int
   next()  int
}

func (qri QueryResultsImpl) fetchData() [5]int {
   qri.array[0]=5
   qri.array[1]=15
   qri.array[2]=25
   qri.array[3]=35
   qri.array[4]=45
   return qri.array
}

type QueryResultsIteratorImpl struct{
   max int
   ar[5] int
}

func (qrii QueryResultsIteratorImpl) next() int{
   pointer=pointer+1
   return qrii.ar[pointer]
}

func (qrii QueryResultsIteratorImpl) first() int{
   return qrii.ar[0]
}

func (qrii QueryResultsIteratorImpl) isDone() bool{
   return pointer==(qrii.max)-1
}

func (qrii QueryResultsIteratorImpl) current() int{
   return qrii.ar[pointer]
}

func main(){
   var newarray [5] int
   a:=QueryResultsImpl{newarray}

   newarray=a.fetchData()
   k:=QueryResultsIteratorImpl{len(newarray),newarray}
   fmt.Println(k.first())
   for{
      if k.isDone()==true{
         break
      }
      fmt.Println(k.next())
   }
}
